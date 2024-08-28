package orders

import (
	"fmt"
	"gorm.io/gorm"
	"storage/services/products"
	"time"
)

func RepoCreateNewOrder(db *gorm.DB, newOrder NewOrder) (int64, error) {
	var orderID int64

	err := db.Transaction(func(tx *gorm.DB) error {
		var totalAmount float64
		var orderDetails []OrderDetail

		for _, product := range newOrder.Products {
			var dbProduct products.Product
			// Check if product exists and has enough quantity
			if err := tx.First(&dbProduct, product.ID).Error; err != nil {
				return fmt.Errorf("product with ID %d not found", product.ID)
			}
			if dbProduct.QuantityInStock < product.Quantity {
				return fmt.Errorf("not enough stock for product ID %d", product.ID)
			}

			// Prepare order details to be inserted later
			orderDetail := OrderDetail{
				ProductID: product.ID,
				Quantity:  product.Quantity,
				Price:     dbProduct.Price,
			}
			totalAmount += dbProduct.Price * float64(product.Quantity)
			orderDetails = append(orderDetails, orderDetail)
		}

		// Create the order if all checks pass
		order := Order{
			CustomerID:  1,
			OrderDate:   time.Now(),
			OrderStatus: "Pending",
			TotalAmount: totalAmount,
		}

		if err := tx.Create(&order).Error; err != nil {
			return err
		}

		// Capture the generated OrderID
		orderID = order.OrderID

		// Insert all order details with the order ID
		for i := range orderDetails {
			orderDetails[i].OrderID = orderID
		}
		if err := tx.Create(&orderDetails).Error; err != nil {
			return err
		}

		// Deduct the quantity from the stock after the order is successful
		for _, product := range newOrder.Products {
			if err := tx.Model(&products.Product{}).Where("product_id = ?", product.ID).
				Update("quantity_in_stock", gorm.Expr("quantity_in_stock - ?", product.Quantity)).Error; err != nil {
				return err
			}
		}

		return nil
	})

	// Return the OrderID and any error encountered
	return orderID, err
}

//func RepoCreateNewOrder(db *gorm.DB, newOrder NewOrder) (int64, error) {
//
//	order := Order{
//		CustomerID:  1,
//		OrderDate:   time.Now(),
//		OrderStatus: "Pending",
//		TotalAmount: 0,
//	}
//
//	if err := db.Create(&order).Error; err != nil {
//		return 0, err
//	}
//
//	var totalAmount float64
//
//	for _, product := range newOrder.Products {
//		var dbProduct products.Product
//		if err := db.First(&dbProduct, product.ID).Error; err != nil {
//			return 0, errors.New(fmt.Sprint("product not found: ", product.ID))
//		}
//
//		orderDetail := OrderDetail{
//			OrderID:   order.OrderID,
//			ProductID: product.ID,
//			Quantity:  product.Quantity,
//			Price:     dbProduct.Price,
//		}
//
//		totalAmount += dbProduct.Price * float64(product.Quantity)
//
//		newQuantity := dbProduct.QuantityInStock - product.Quantity
//
//		if newQuantity < 0 {
//			db.Delete(&order)
//		} else {
//			dbProduct.QuantityInStock = newQuantity
//		}
//
//		if err := db.Save(&dbProduct).Error; err != nil {
//			return 0, err
//		}
//
//		if err := db.Create(&orderDetail).Error; err != nil {
//			return 0, err
//		}
//	}
//
//	order.TotalAmount = int64(totalAmount)
//	if err := db.Save(&order).Error; err != nil {
//		return 0, err
//	}
//
//	return order.OrderID, nil
//}

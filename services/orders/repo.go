package orders

import (
	"fmt"
	"storage/services/products"
	"time"

	"gorm.io/gorm"
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

func RepoGetOrderById(db *gorm.DB, orderID int64) (OrderView, error) {
	var orderView OrderView

	// Query the order information
	if err := db.Model(&Order{}).
		Where("order_id = ?", orderID).
		First(&orderView).Error; err != nil {
		return OrderView{}, err
	}

	// Query the order details and join with the products
	var products []products.Product
	err := db.Table("storageuser.order_details as od").
		Joins("inner join storageuser.products as p on p.product_id = od.product_id").
		Select("p.*, od.quantity").
		Where("od.order_id = ?", orderID).
		Scan(&products).Error

	if err != nil {
		// handle error
	}

	if err != nil {
		return OrderView{}, err
	}

	// Assign products to the order view
	orderView.Products = products

	return orderView, nil
}

func RepoGetAllOrders(db *gorm.DB) ([]Order, error) {
	var orders []Order
	if err := db.Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

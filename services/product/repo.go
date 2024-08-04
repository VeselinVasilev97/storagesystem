package product

import (
	// "database/sql"

	"gorm.io/gorm"
)

func RepoGetAllProducts(db *gorm.DB) ([]Product, error) {
	var products []Product
	if err := db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func RepoGetProductById(db *gorm.DB, productId int64) (Product, error) {
	var product Product
	result := db.Where("product_id = ?", productId).First(&product)
	if result.RowsAffected == 0 {
		return Product{}, nil
	}
	return product, result.Error
}

// func RepoGetProductByIdDetailed(db *gorm.DB, productId int64) (Product, error) {
// 	var product Product
// 	var category category.Category
// 	var supplier Supplier

// 	query := `
// 		SELECT p.product_id, p.product_name, p.product_description, p.category_id, p.supplier_id, p.quantity_in_stock, p.price,
// 			   c.category_id, c.category_name, c.description,
// 			   s.supplier_id, s.supplier_name, s.contact_name, s.contact_email, s.contact_phone, s.supplier_address
// 		FROM storageuser.products p
// 		JOIN storageuser.categories c ON p.category_id = c.category_id
// 		JOIN storageuser.suppliers s ON p.supplier_id = s.supplier_id
// 		WHERE p.product_id = ?
// 	`
// 	row := db.Raw(query, productId).Row()
// 	err := row.Scan(
// 		&product.ProductID, &product.ProductName, &product.ProductDescription,
// 		&product.CategoryID, &product.SupplierID, &product.QuantityInStock, &product.Price,
// 		&category.CategoryID, &category.CategoryName, &category.Description,
// 		&supplier.SupplierID, &supplier.SupplierName, &supplier.ContactName, &supplier.ContactEmail,
// 		&supplier.ContactPhone, &supplier.SupplierAddress,
// 	)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return Product{}, nil
// 		}
// 		return Product{}, err
// 	}

// 	product.Category = category
// 	product.Supplier = supplier

// 	return product, nil
// }

package product

import "storage/services/categories"

type Product struct {
	ProductID          int                 `gorm:"column:product_id;primaryKey" json:"product_id"`
	ProductName        string              `gorm:"column:product_name;size:50;not null" json:"product_name"`
	ProductDescription *string             `gorm:"column:product_description;size:200" json:"product_description,omitempty"`
	CategoryID         int                 `gorm:"column:category_id;not null" json:"category_id"`
	SupplierID         int                 `gorm:"column:supplier_id;not null" json:"supplier_id"`
	QuantityInStock    int                 `gorm:"column:quantity_in_stock;not null" json:"quantity_in_stock"`
	Price              float64             `gorm:"column:price;type:numeric(10,2);not null" json:"price"`
	Category           categories.Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	// Supplier           Supplier `gorm:"foreignKey:SupplierID" json:"supplier,omitempty"`
}

func (Product) TableName() string {
	return "storageuser.products"
}

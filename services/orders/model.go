package orders

import (
	"storage/services/products"
	"time"
)

type Order struct {
	OrderID     int64     `gorm:"column:order_id;primaryKey" json:"order_id"`
	CustomerID  int64     `gorm:"column:customer_id;not null" json:"customer_id"`
	OrderDate   time.Time `gorm:"column:order_date;not null" json:"order_date"`
	OrderStatus string    `gorm:"column:order_status;not null;size:10" json:"order_status"`
	TotalAmount float64   `gorm:"column:total_amount;not null;type:numeric(10,2)" json:"total_amount"`
}

type OrderDetail struct {
	OrderDetailID int64   `gorm:"column:order_detail_id;primaryKey" json:"order_detail_id"`
	OrderID       int64   `gorm:"column:order_id;not null" json:"order_id"`
	ProductID     int64   `gorm:"column:product_id;not null" json:"product_id"`
	Quantity      int64   `gorm:"column:quantity;not null" json:"quantity"`
	Price         float64 `gorm:"column:price;not null;type:numeric(10,2)" json:"price"`
}

type OrderView struct {
	OrderID     int64              `gorm:"column:order_id;primaryKey" json:"order_id"`
	CustomerID  int64              `gorm:"column:customer_id;not null" json:"customer_id"`
	OrderDate   time.Time          `gorm:"column:order_date;not null" json:"order_date"`
	OrderStatus string             `gorm:"column:order_status;not null;size:10" json:"order_status"`
	TotalAmount float64            `gorm:"column:total_amount;not null;type:numeric(10,2)" json:"total_amount"`
	Products    []products.Product `gorm:"foreignKey:ProductID" json:"products"`
}

type NewOrder struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID       int64 `json:"product_id"`
	Quantity int64 `json:"quantity"`
}

func (Order) TableName() string {
	return "storageuser.orders"
}

func (OrderDetail) TableName() string {
	return "storageuser.order_details"
}

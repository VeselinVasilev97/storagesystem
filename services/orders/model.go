package orders

import "time"

type Order struct {
	OrderID     int64     `gorm:"column:order_id;primaryKey" json:"order_id"`
	CustomerID  int64     `gorm:"column:customer_id;not null" json:"customer_id"`
	OrderDate   time.Time `gorm:"column:order_date;not null" json:"order_date"`
	OrderStatus string    `gorm:"column:order_status;not null;size:10" json:"order_status"`
	TotalAmount int64     `gorm:"column:total_amount;not null;type:numeric(10,2)" json:"total_amount"`
}

type NewOrder struct {
	Products []Product `json:"products"`
}

type Product struct {
	ID       int64 `json:"order_id"`
	Quantity int64 `json:"quantity"`
}

func (Order) TableName() string {
	return "storageuser.orders"
}

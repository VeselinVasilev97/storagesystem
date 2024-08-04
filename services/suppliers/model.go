package suppliers

type Supplier struct {
	SupplierID      int    `gorm:"column:supplier_id;primaryKey" json:"supplier_id"`
	SupplierName    string `gorm:"column:supplier_name;size:50;not null" json:"supplier_name"`
	ContactName     string `gorm:"column:contact_name;size:100;not null" json:"contact_name"`
	ContactEmail    string `gorm:"column:contact_email;size:100;not null" json:"contact_email"`
	ContactPhone    string `gorm:"column:contact_phone;size:15;not null" json:"contact_phone"`
	SupplierAddress string `gorm:"column:supplier_address;size:200;not null" json:"supplier_address"`
}

func (Supplier) TableName() string {
	return "storageuser.suppliers"
}

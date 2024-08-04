package categories

type Category struct {
	CategoryID          int    `gorm:"column:category_id;primaryKey" json:"category_id"`
	CategoryName        string `gorm:"column:category_name;size:50;not null" json:"category_name"`
	CategoryDescription string `gorm:"column:category_description;size:500" json:"category_description,omitempty"`
}

func (Category) TableName() string {
	return "storageuser.categories"
}

package categories

import (
	// "database/sql"

	"gorm.io/gorm"
)

func RepoGetAllCategories(db *gorm.DB) ([]Category, error) {
	var categories []Category
	if err := db.Find(&categories).Error; err != nil {
		return nil, err
	}
	return categories, nil
}

func RepoGetCategoryById(db *gorm.DB, categoryId int64) (Category, error) {
	var category Category

	result := db.Where("category_id = ?", categoryId).First(&category)
	if result.Error != nil {
		return Category{}, result.Error
	}
	return category, nil
}

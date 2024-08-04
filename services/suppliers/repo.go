package suppliers

import "gorm.io/gorm"

func RepoGetAllSuppliers(db *gorm.DB) ([]Supplier, error) {
	var suppliers []Supplier
	if err := db.Find(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

func RepoGetSupplierById(db *gorm.DB, supplierId int64) (Supplier, error) {
	var supplier Supplier
	if err := db.Where("supplier_id = ?", supplierId).First(&supplier).Error; err != nil {
		return Supplier{}, err
	}
	return supplier, nil
}

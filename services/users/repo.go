package users

import "gorm.io/gorm"

func RepoGetAllUsers(db *gorm.DB) ([]User, error) {
	var suppliers []User
	if err := db.Find(&suppliers).Error; err != nil {
		return nil, err
	}
	return suppliers, nil
}

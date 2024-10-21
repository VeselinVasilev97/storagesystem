package user

import "gorm.io/gorm"

func RepoGetAllUsers(db *gorm.DB) ([]User, error) {
	var users []User
	if err := db.Preload("Roles").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

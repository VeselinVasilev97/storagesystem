package users

import "time"

type User struct {
	UserID    int64     `gorm:"column:id;primaryKey" json:"id"`
	Username  string    `gorm:"column:username;size:50;not null" json:"username"`
	Email     string    `gorm:"column:email;size:100;not null" json:"email"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	LastLogin time.Time `gorm:"column:last_login" json:"last_login"`
	IsActive  bool      `gorm:"column:is_active" json:"is_active"`
}

func (User) TableName() string {
	return "storageuser.users"
}

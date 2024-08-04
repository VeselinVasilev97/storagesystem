package configuration

import "gorm.io/gorm"

type Config struct {
	Port     string `json:"port"`
	Db       *gorm.DB
	Database DB `json:"database"`
}

type DB struct {
	User         string `json:"user"`
	Password     string `json:"password"`
	DatabaseName string `json:"database_name"`
	Port         string `json:"port"`
	Host         string `json:"host"`
}

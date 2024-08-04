package configuration

import (
	"encoding/json"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
	"os"
)

var Configuration *Config

func LoadConfig() *Config {
	data, err := os.ReadFile("./configuration/config.json")
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	err = json.Unmarshal(data, &Configuration)
	if err != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	Configuration.Db = SetUpDatabase()

	return Configuration
}

func SetUpDatabase() *gorm.DB {
	dsn := "user=" + Configuration.Database.User +
		" password=" + Configuration.Database.Password +
		" dbname=" + Configuration.Database.DatabaseName +
		" port=" + Configuration.Database.Port +
		" host=" + Configuration.Database.Host +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

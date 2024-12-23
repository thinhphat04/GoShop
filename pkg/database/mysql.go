package database

import (
	"API/config"
	"API/internal/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := config.GetMySQLDSN()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	// Lưu kết nối vào biến toàn cục
	DB = db

	// Auto-migrate nếu cần
	err = DB.AutoMigrate(
		&models.User{},
		&models.Order{},
		&models.Product{},
		&models.Cart{},
		&models.Coupon{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Connected to MySQL and migrated database!")
}

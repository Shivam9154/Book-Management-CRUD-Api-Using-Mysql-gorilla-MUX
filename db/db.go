package db

import (
	"book-management/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Using localhost since MySQL runs in Docker locally
	dsn := "root:Shivam@123@tcp(localhost:3306)/bookdb?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	// Auto migrate the Book model
	DB.AutoMigrate(&models.Book{})
}

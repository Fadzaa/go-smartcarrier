package infrastructure

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-gin-api/domain/auth"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func ConnectToDatabase() *gorm.DB {
	err := godotenv.Load()
	dsn := os.Getenv("DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&auth.User{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected")

	return db
}

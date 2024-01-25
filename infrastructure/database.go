package infrastructure

import (
	"fmt"
	"go-gin-api/domains/job"
	"go-gin-api/domains/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectToDatabase() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/api_golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&user.User{}, &job.Job{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database Connected")

	return db
}

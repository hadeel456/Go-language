package config

import (
	"fmt"
	model "go-language/model"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	godotenv.Load()
	dbhost := os.Getenv("MySQL_HOST")
	dbuser := os.Getenv("MySQL_USER")
	dbpassword := os.Getenv("MySQL_PASSWORD")
	dbname := os.Getenv("MySQL_DBNAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbname)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
	DB = db
	fmt.Println("db connectd successfully")
	AutoMigrate(db)
}
func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&model.Author{},
		&model.Book{},
	)

}

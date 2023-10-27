package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"backend/models"
)

var DB *gorm.DB

func getMySQLDialector(host, user, password, dbname string, port uint64 ) gorm.Dialector{

	// EXAMPLE MySQL :"user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbname)
	return mysql.Open(dsn)
}

func autoMigrate(){
	DB.AutoMigrate(&models.Category{})
	DB.AutoMigrate(&models.Location{})
	DB.AutoMigrate(&models.Post{})
	DB.AutoMigrate(&models.Site{})
	DB.AutoMigrate(&models.Subcategory{})
	DB.AutoMigrate(&models.User{})

}


func ConnectDB() {
	var err error

    p := os.Getenv("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
    if err != nil {
        log.Println("Error")
    }
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dialector := getMySQLDialector(host, user, password, dbname, port)

    DB, err = gorm.Open(dialector)

    if err != nil {
        panic("failed to connect database")
    }

    fmt.Println("Connection Opened to Database")
}

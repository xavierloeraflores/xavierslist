package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
	dbname := os.Getenv("DB_Name")

    dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	
    DB, err = gorm.Open(mysql.Open(dsn))

    if err != nil {
        panic("failed to connect database")
    }

    fmt.Println("Connection Opened to Database")
}

package database

import (
	"log"
	"os"

	"github.com/codewithmujab/gocrud/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

// connect DB
func ConnectDb() {
	dbUser := "root"                  //user database
	dbPassword := "root"              // password database
	dbName := "gocrud"                // nama database
	dbHost := "@tcp(127.0.0.1:8889)/" //nama host

	dsn := dbUser + ":" + dbPassword + dbHost + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database. \n", err)
		os.Exit(2)
	}

	log.Println("connected to database")

	//jalankan migrate dari model
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.Book{})
	DBConn = db
}

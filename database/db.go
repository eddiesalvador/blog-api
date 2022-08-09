package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() (*gorm.DB) {
	var err error
	dsn := DB_USERNAME +":"+ DB_PASSWORD +"@tcp"+ "(" + DB_HOST + ":" + DB_PORT +")/" + DB_NAME + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	

	if err != nil {
		fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}

	return db
}
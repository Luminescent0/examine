package dao

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var dB *gorm.DB

func InitDB() {
	sqlDB, err := sql.Open("mysql", "root:xianye@tcp(localhost)/db1")
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	dB = gormDB

}

package controllers

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/db_uts?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// func connectGorm() (*gorm.DB, error) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/db_uts?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return db, nil
// }

package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DataBase *gorm.DB

func InitDataBase() {
	dsn := "user=arifu password=arifu dbname=test port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	return err
	// }

	DataBase = db
}

func GetDb() **gorm.DB {
	return &DataBase
}

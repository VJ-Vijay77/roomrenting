package database

import (
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//connecting database and returning database engine
func GetDb() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Unable to get data from env!!")
	}
	// dsn := os.Getenv("database_address")
	dsn := "host=localhost user=vijay password=zmxmcmvbn dbname=r4room port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database!")
	}
	return db
}

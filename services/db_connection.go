package services

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func ConnectDB() (*gorm.DB,error){

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		return nil,err
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")	

	strConnection := fmt.Sprintf("host=%s user=%s password=%s dbname=sellpointgo port=%s", host,user,password,port)

	db,err := gorm.Open(postgres.New(postgres.Config{
		DSN:strConnection,
		PreferSimpleProtocol:true,
	}),&gorm.Config{})
		
	return db,err
}
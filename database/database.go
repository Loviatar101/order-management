package database

import (
	"log"
	"os"

	"github.com/loviatar101/order-consumer/database/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


type Dbinstance struct{
	Db *gorm.DB
	
}

var Database Dbinstance

func ConnectDb(){
	db,err:=gorm.Open(sqlite.Open("order.db"),&gorm.Config{})
	if err!=nil{
		log.Fatal("failed to connect to databse\n",err.Error())
		os.Exit(2)
	}
	log.Println("connected to database succefully")
	db.Logger=logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.User{},&models.Product{},models.Order{})


	Database=Dbinstance{Db: db}

}
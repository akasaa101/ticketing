package database

import (
	"fmt"
	"github.com/akasaa101/ticketing/internal/config"
	"github.com/akasaa101/ticketing/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
)

type DbInstance struct {
	Db *gorm.DB
}

var DB DbInstance

func Connect() {
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		fmt.Println("Error parsing str to int")
	}
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai", config.Config("DB_HOST"), config.Config("DB_USERNAME"), config.Config("DB_PASSWORD"), config.Config("DB_DATABASE"), port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}
	log.Println("Connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	err = db.AutoMigrate(&models.Ticket{})
	if err != nil {
		return
	}
	DB = DbInstance{
		Db: db,
	}
}

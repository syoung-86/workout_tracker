package database

import (
	//"database/sql"
	//"fmt"
	"log"

	//"log"

	//"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"workout_tracker/internal/models"
)

type Dbisntance struct {
	Db *gorm.DB
}

var DB Dbisntance

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "workout_tracker"
)

func Connect() {
	var err error
	dsn := "host=localhost user=postgres password='' dbname=workout_tracker port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database \n", err)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("running migrations")
	db.AutoMigrate(&models.User{})

	DB = Dbisntance{
		Db: db,
	}
}
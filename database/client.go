package database

import (
	"crud-golang/entities"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// define the global variable that will be used to access the database
var Instance *gorm.DB
var err error

// connect to the database via gorm using the connection string
// from the config file
func Connect(connectionString string) {
	Instance, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to database")
	}

	log.Println("Database connection established")
}

// ensure that a table named product ios created on the database
func Migrate() {
	Instance.AutoMigrate(&entities.Product{})
	log.Println("Database migration completed")
}

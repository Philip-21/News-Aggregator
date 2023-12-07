package database

import (
	"fmt"
	"log"
	"os"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Unable to open db, ", err)
		return nil, err
	}
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Println("Error in Migration", err)
	}
	//verifies if a connection to the database is still alive, establishing a connection if necessary.
	sqlDB, err := db.DB()
	if err != nil {
		log.Println("error in getting sql")
	}

	err = sqlDB.Ping()
	if err != nil {
		log.Println("Error in connection", err)
		return nil, err
	}
	return db, nil

}

var counts int64

// connects to the database properly
func ConnectToDB() *gorm.DB {
	// dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s",
	// 	os.Getenv("POSTGRES_USER"),
	// 	os.Getenv("POSTGRES_PASSWORD"),
	// 	os.Getenv("POSTGRES_DB"),
	// 	os.Getenv("POSTGRES_HOST"),
	// 	os.Getenv("POSTGRES_PORT"),
	// )
	dsn := os.Getenv("DSN")
	fmt.Println("DSN:", dsn)
	//an infinite for loop to connect to the database
	for {
		connection, err := OpenDB(dsn)
		if err != nil {
			log.Println("Postgres not yet connected......")
			counts++ //add 1 to counts
		} else {
			log.Println("Connected to Postgres")
			return connection
		}
		//tryin to connect to database for 20 seconds
		if counts > 10 {
			log.Println("connection rate exceeded", err)
			return nil
		}
		log.Println("Backing of for two seconds.......")
		time.Sleep(2 * time.Second) //waiting for 2sec each time
		continue
	}
}

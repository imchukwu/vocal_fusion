package config

import (
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// ConnectDB initializes and returns a GORM DB connection
func ConnectDB() *gorm.DB {
	once.Do(func() {
		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			log.Fatal("❌ DATABASE_URL not set in environment")
		}

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("❌ Failed to connect to database: %v", err)
		}

		db = conn
		log.Println("✅ Database connection established")
	})

	return db
}

// GetDB returns the global DB instance
func GetDB() *gorm.DB {
	if db == nil {
		return ConnectDB()
	}
	return db
}

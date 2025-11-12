// db/db.go
package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes and returns a PostgreSQL database connection.
func InitDB() (*gorm.DB, error) {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("--------------DATABASE_URL is not set")
	}

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dbURL,
		PreferSimpleProtocol: true, // ✅ disables prepared statement caching to avoid "already exists" error
	}), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("--------------failed to connect database: %v", err)
	}

	// Optional: Ping to verify connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("--------------failed to get sql.DB: %v", err)
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("--------------database ping failed: %v", err)
	}

	return db, nil
}

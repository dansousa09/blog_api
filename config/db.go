package config

import (
	"os"

	"github.com/dansousa09/blog_api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeDB() (*gorm.DB, error) {
	logger := GetLogger("db")
	dbPath := "./db/blog.db"

	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found. Creating...")
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}
		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("Error opening database: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Post{})
	if err != nil {
		logger.Errorf("Error automigrating database: %v", err)
		return nil, err
	}

	return db, nil
}

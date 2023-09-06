package config

import (
	"fmt"

	"gorm.io/gorm"
)

var logger *Logger
var db *gorm.DB

func Init() error {
	var err error

	db, err = InitializeDB()

	if err != nil {
		return fmt.Errorf("error initializing db: %v", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger, err := NewLogger(p)
	if err != nil {
		fmt.Errorf("error on initialize Logger: %s", err.Error())
	}
	return logger
}

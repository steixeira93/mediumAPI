package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.db
	logger *Logger
)

func Init() error {
	var err error
	db, err = InitializeSqlite()

	if err != nil {
		return fmt.Errorf("failed to initialize database: %v", err)
	}
	return nil
}

func GetSQliteDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

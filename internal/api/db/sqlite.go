package db

import (
	"log/slog"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectSqlite(fname string, logger *slog.Logger) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(fname), &gorm.Config{})
	if err != nil {
		logger.Error("err", err)
	}

	return db
}

package database

import (
	"github.com/dickanirwansyah/go-app/go-tech/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectionDB() {
	database, err := gorm.Open(mysql.Open("root:rootroot@/dbgolang"), &gorm.Config{})
	if err != nil {
		panic("error tidak dapat terkoneksi dengan database")
	}

	database.AutoMigrate(&models.User{})
}

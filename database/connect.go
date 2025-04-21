package database

import (
	"log"

	"github.com/doneowth/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:my-secret-pw@/mydb"), &gorm.Config{})

	if err != nil {
		panic("hi db issue")
	}
	
	DB = db
	db.AutoMigrate(models.User{})
	log.Println("db connected")
}

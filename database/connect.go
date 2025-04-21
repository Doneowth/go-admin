package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func Connect() {
	db, err := gorm.Open(mysql.Open("root:my-secret-pw@/mydb"), &gorm.Config{})

	if err != nil {
		panic("hi db issue")
	}

	log.Print(db)
}

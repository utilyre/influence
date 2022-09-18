package database

import (
	"fmt"
	"os"
	"server/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Connect() error {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASS"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_PORT"),
		os.Getenv("DATABASE_NAME"),
	)

	var (
		db  *gorm.DB
		err error
	)
	for i := 0; i < 10; i++ {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		}

		fmt.Println("Successfully connected to database")
		break
	}

	if err != nil {
		return err
	}

	db.AutoMigrate(
		&models.User{},
		&models.Blog{},
	)

	Instance = db
	return nil
}

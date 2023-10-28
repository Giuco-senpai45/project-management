package db

import (
	"fmt"
	"go-project/models"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

func NewAdapter(config DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", config.Username, config.Password, config.Host, config.Port, config.DBName)
	var db *gorm.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		} else {
			fmt.Printf("Failed to connect to database (%d), retrying...\n", i)
			time.Sleep(5 * time.Second)
		}
	}
	if err != nil {
		fmt.Printf("Failed to connect to database after 5 retries\n")
		return nil, err
	}

	err = testDB(db)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Project{})
	db.AutoMigrate(&models.Task{})
	db.AutoMigrate(&models.SubTask{})

	return db, nil
}

func testDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	fmt.Println("*** Pinged database successfully! ***")

	log.Println("MUIE LA CAI")

	return nil
}

func CloseDbConnection(db *gorm.DB) {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}

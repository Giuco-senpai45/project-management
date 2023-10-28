package models

import (
	"time"

	"gorm.io/gorm"
)

var db *gorm.DB

const dbTimeout = time.Second * 3

func New(dbPool *gorm.DB) Models {
	db = dbPool
	return Models{}
}

type Models struct {
	User    User
	Project Project
	Task    Task
	SubTask SubTask
}

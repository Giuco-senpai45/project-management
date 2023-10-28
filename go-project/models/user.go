package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string     `json:"Name"`
	Email    string     `json:"Email"`
	Password string     `json:"Password"`
	Projects []*Project `gorm:"many2many:user_projects;"`
}

func (u *User) CreateUser(user *User) (*User, error) {
	result := db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *User) GetUserByID(id int64) (*User, error) {
	var user User
	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *User) GetAllUsers() ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (u *User) UpdateUser(user *User) (*User, error) {
	result := db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (u *User) LoginUser(email string) (*User, error) {
	var user User
	result := db.First(&user, "Email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (u *User) GetAllUserProjects(currentUser *User) ([]Project, error) {
	var projects []Project
	err := db.Model(currentUser).Association("Projects").Find(&projects)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

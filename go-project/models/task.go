package models

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	Urgency     int64     `json:"urgency"`
	DueDate     time.Time `json:"dueDate"`
	ProjectID   uint64    `json:"projectID"`
}

func (t *Task) CreateTask(task *Task, project *Project) (*Task, error) {
	result := db.Create(&task)
	db.Model(project).Association("Tasks").Append(task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (t *Task) GetTaskByID(id int64) (*Task, error) {
	var task Task
	result := db.First(&task, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &task, nil
}

func (t *Task) GetAllTasks() ([]Task, error) {
	var tasks []Task
	result := db.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}

func (t *Task) UpdateTask(task *Task) (*Task, error) {
	result := db.Save(&task)
	if result.Error != nil {
		return nil, result.Error
	}
	return task, nil
}

func (t *Task) DeleteTask(task *Task) error {
	result := db.Delete(task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package models

import "gorm.io/gorm"

type SubTask struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

func (st *SubTask) CreateSubTask(subtask *SubTask) (*SubTask, error) {
	result := db.Create(&subtask)
	if result.Error != nil {
		return nil, result.Error
	}
	return subtask, nil
}

func (st *SubTask) GetSubTaskByID(id int64) (*SubTask, error) {
	var subtask SubTask
	result := db.First(&subtask, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &subtask, nil
}

func (st *SubTask) GetAllSubTasks() ([]SubTask, error) {
	var subtasks []SubTask
	result := db.Find(&subtasks)
	if result.Error != nil {
		return nil, result.Error
	}
	return subtasks, nil
}

func (st *SubTask) UpdateSubTask(subtask *SubTask) (*SubTask, error) {
	result := db.Save(&subtask)
	if result.Error != nil {
		return nil, result.Error
	}
	return subtask, nil
}

func (st *SubTask) DeleteSubTask(subtask *SubTask) error {
	result := db.Delete(subtask)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Users       []*User `gorm:"many2many:user_projects;"`
	Tasks       []Task  `gorm:"foreignKey:ProjectID"`
}

func (p *Project) CreateProject(project *Project, user *User) (*Project, error) {
	result := db.Create(&project)
	db.Model(project).Association("Users").Append(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (p *Project) GetProjectByID(id int64) (*Project, error) {
	var project Project
	result := db.First(&project, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &project, nil
}

func (p *Project) GetAllProjects() ([]Project, error) {
	var projects []Project
	result := db.Find(&projects)
	if result.Error != nil {
		return nil, result.Error
	}
	return projects, nil
}

func (p *Project) UpdateProject(project *Project) (*Project, error) {
	result := db.Save(&project)
	if result.Error != nil {
		return nil, result.Error
	}
	return project, nil
}

func (p *Project) DeleteProject(project *Project) error {
	result := db.Delete(project)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (p *Project) GetAllProjectTasks(project *Project) ([]Task, error) {
	var tasks []Task
	err := db.Model(project).Association("Tasks").Find(&tasks)
	if err != nil {
		return nil, err
	}
	// project.Tasks = tasks
	return tasks, nil
}

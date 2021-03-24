package models

import (
	"time"

	"bitbucket.com/iamjollof/server/config"
	_ "github.com/go-playground/validator/v10"
)

//Project - project data struct
type Project struct {
	ID            int32      `json:"id,omitempty" sql:"primary_key"`
	Title         string     `json:"title" validate:"required"`
	Content       string     `json:"content" validate:"required"`
	Slug          string     `json:"slug"`
	Summary       string     `json:"summary" validate:"required"`
	FeaturedImage string     `json:"featured_image"`
	Company       string     `json:"company"`
	Location      string     `json:"location"`
	Website       string     `json:"website"`
	Repo          string     `json:"repo"`
	Active        string     `gorm:"default:false" json:"active"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

//TableName - projects db table name override
func (Project) TableName() string {
	return "projects"
}

//GetAllProjects - fetch all projects at once
func GetAllProjects(offset int32, limit int32) ([]Project, int32, error) {
	var count int32
	var project []Project
	if err := config.DB.Model(&Project{}).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(&project).Error; err != nil {
		return project, count, err
	}
	return project, count, nil
}

//CreateProject - create a project
func CreateProject(project *Project) (created bool, err error) {
	if errs := config.DB.Create(&project).Error; errs != nil {
		return false, errs
	}
	return false, nil
}

//GetProject - fetch one project
func GetProject(id int32) (Project, error) {
	var project Project
	if err := config.DB.Where("id = ?", id).First(&project).Error; err != nil {
		return project, err
	}
	return project, nil
}

//UpdateProject - update a project
func UpdateProject(project Project, id int32) (err error) {
	if err := config.DB.Model(&Project{}).Where("id = ?", id).Updates(project).Error; err != nil {
		return err
	}
	return nil
}

//DeleteProject - delete a project
func DeleteProject(id int32) (err error) {
	if err := config.DB.Where("id = ?", id).Unscoped().Delete(Project{}).Error; err != nil {
		return err
	}
	return nil
}

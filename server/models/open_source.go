package models

import (
	"time"

	"bitbucket.com/iamjollof/server/config"
	_ "github.com/go-playground/validator/v10"
)

//OpenSource - open source data struct
type OpenSource struct {
	ID            int32      `json:"id,omitempty" sql:"primary_key"`
	Title         string     `json:"title" validate:"required"`
	Content       string     `json:"content" validate:"required"`
	Slug          string     `json:"slug"`
	Summary       string     `json:"summary" validate:"required"`
	FeaturedImage string     `json:"featured_image"`
	Website       string     `json:"website"`
	Repo          string     `json:"repo"`
	Active        string     `gorm:"default:false" json:"active"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

//TableName - openSources db table name override
func (OpenSource) TableName() string {
	return "open_sources"
}

//GetAllOpenSources - fetch all openSources at once
func GetAllOpenSources(offset int32, limit int32) ([]OpenSource, int32, error) {
	var count int32
	var openSource []OpenSource
	if err := config.DB.Model(&OpenSource{}).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(&openSource).Error; err != nil {
		return openSource, count, err
	}
	return openSource, count, nil
}

//CreateOpenSource - create a openSource
func CreateOpenSource(openSource *OpenSource) (created bool, err error) {
	if errs := config.DB.Create(&openSource).Error; errs != nil {
		return false, errs
	}
	return false, nil
}

//GetOpenSource - fetch one openSource
func GetOpenSource(id int32) (OpenSource, error) {
	var openSource OpenSource
	if err := config.DB.Where("id = ?", id).First(&openSource).Error; err != nil {
		return openSource, err
	}
	return openSource, nil
}

//UpdateOpenSource - update a openSource
func UpdateOpenSource(openSource OpenSource, id int32) (err error) {
	if err := config.DB.Model(&OpenSource{}).Where("id = ?", id).Updates(openSource).Error; err != nil {
		return err
	}
	return nil
}

//DeleteOpenSource - delete a openSource
func DeleteOpenSource(id int32) (err error) {
	if err := config.DB.Where("id = ?", id).Unscoped().Delete(OpenSource{}).Error; err != nil {
		return err
	}
	return nil
}

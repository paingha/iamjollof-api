package models

import (
	"time"

	"bitbucket.com/iamjollof/server/config"
	_ "github.com/go-playground/validator/v10"
)

//Quote - quote data struct
type Quote struct {
	ID        int32      `json:"id,omitempty" sql:"primary_key"`
	Author    string     `json:"author" validate:"required"`
	Content   string     `json:"content" validate:"required"`
	Active    string     `gorm:"default:false" json:"active"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

//TableName - quotes db table name override
func (Quote) TableName() string {
	return "quotes"
}

//GetAllQuotes - fetch all quotes at once
func GetAllQuotes(offset int32, limit int32) ([]Quote, int32, error) {
	var count int32
	var quote []Quote
	if err := config.DB.Model(&Quote{}).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(&quote).Error; err != nil {
		return quote, count, err
	}
	return quote, count, nil
}

//CreateQuote - create a quote
func CreateQuote(quote *Quote) (created bool, err error) {
	if errs := config.DB.Create(&quote).Error; errs != nil {
		return false, errs
	}
	return false, nil
}

//GetQuote - fetch one quote
func GetQuote(id int32) (Quote, error) {
	var quote Quote
	if err := config.DB.Where("id = ?", id).First(&quote).Error; err != nil {
		return quote, err
	}
	return quote, nil
}

//UpdateQuote - update a quote
func UpdateQuote(quote Quote, id int32) (err error) {
	if err := config.DB.Model(&Quote{}).Where("id = ?", id).Updates(quote).Error; err != nil {
		return err
	}
	return nil
}

//DeleteQuote - delete a quote
func DeleteQuote(id int32) (err error) {
	if err := config.DB.Where("id = ?", id).Unscoped().Delete(Quote{}).Error; err != nil {
		return err
	}
	return nil
}

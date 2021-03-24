package models

import (
	"time"

	"bitbucket.com/iamjollof/server/config"
	_ "github.com/go-playground/validator/v10"
)

//Blog - blog data struct
type Blog struct {
	ID            int32      `json:"id,omitempty" sql:"primary_key"`
	Title         string     `json:"title" validate:"required"`
	Content       string     `json:"content" validate:"required"`
	Slug          string     `json:"slug"`
	Summary       string     `json:"summary" validate:"required"`
	FeaturedImage string     `json:"featured_image"`
	Active        string     `gorm:"default:false" json:"active"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

//TableName - blogs db table name override
func (Blog) TableName() string {
	return "blogs"
}

//GetAllBlogs - fetch all blogs at once
func GetAllBlogs(offset int32, limit int32) ([]Blog, int32, error) {
	var count int32
	var blog []Blog
	if err := config.DB.Model(&Blog{}).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(&blog).Error; err != nil {
		return blog, count, err
	}
	return blog, count, nil
}

//CreateBlog - create a blog
func CreateBlog(blog *Blog) (created bool, err error) {
	if errs := config.DB.Create(&blog).Error; errs != nil {
		return false, errs
	}
	return false, nil
}

//GetBlog - fetch one blog
func GetBlog(id int32) (Blog, error) {
	var blog Blog
	if err := config.DB.Where("id = ?", id).First(&blog).Error; err != nil {
		return blog, err
	}
	return blog, nil
}

//UpdateBlog - update a blog
func UpdateBlog(blog Blog, id int32) (err error) {
	if err := config.DB.Model(&Blog{}).Where("id = ?", id).Updates(blog).Error; err != nil {
		return err
	}
	return nil
}

//DeleteBlog - delete a blog
func DeleteBlog(id int32) (err error) {
	if err := config.DB.Where("id = ?", id).Unscoped().Delete(Blog{}).Error; err != nil {
		return err
	}
	return nil
}

package models

import (
	"time"

	"bitbucket.com/iamjollof/server/config"
	_ "github.com/go-playground/validator/v10"
)

//Recipe - recipe data struct
type Recipe struct {
	ID            int32      `json:"id,omitempty" sql:"primary_key"`
	Title         string     `json:"title" validate:"required"`
	Content       string     `json:"content" validate:"required"`
	Slug          string     `json:"slug"`
	FeaturedImage string     `json:"featured_image"`
	Summary       string     `json:"summary" validate:"required"`
	Active        string     `gorm:"default:false" json:"active"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	DeletedAt     *time.Time `json:"deleted_at"`
}

//TableName - recipes db table name override
func (Recipe) TableName() string {
	return "recipes"
}

//GetAllRecipes - fetch all recipes at once
func GetAllRecipes(offset int32, limit int32) ([]Recipe, int32, error) {
	var count int32
	var recipe []Recipe
	if err := config.DB.Model(&Recipe{}).Count(&count).Order("created_at desc").Offset(offset).Limit(limit).Find(&recipe).Error; err != nil {
		return recipe, count, err
	}
	return recipe, count, nil
}

//CreateRecipe - create a recipe
func CreateRecipe(recipe *Recipe) (created bool, err error) {
	if errs := config.DB.Create(&recipe).Error; errs != nil {
		return false, errs
	}
	return false, nil
}

//GetRecipe - fetch one recipe
func GetRecipe(id int32) (Recipe, error) {
	var recipe Recipe
	if err := config.DB.Where("id = ?", id).First(&recipe).Error; err != nil {
		return recipe, err
	}
	return recipe, nil
}

//UpdateRecipe - update a recipe
func UpdateRecipe(recipe Recipe, id int32) (err error) {
	if err := config.DB.Model(&Recipe{}).Where("id = ?", id).Updates(recipe).Error; err != nil {
		return err
	}
	return nil
}

//DeleteRecipe - delete a recipe
func DeleteRecipe(id int32) (err error) {
	if err := config.DB.Where("id = ?", id).Unscoped().Delete(Recipe{}).Error; err != nil {
		return err
	}
	return nil
}

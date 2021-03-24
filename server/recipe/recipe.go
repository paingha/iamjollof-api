// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package recipe

import (
	"context"

	"bitbucket.com/iamjollof/server/models"
	"bitbucket.com/iamjollof/server/plugins"
	"bitbucket.com/iamjollof/server/protos/recipe"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/copier"
)

//Server - Struct holding Recipes
type Server struct {
	recipe.RecipeServer
}

//GetAllRecipes - Lists all the recipes
//Todo: Handle errors
//sudo service postgresql start
//export PATH="$PATH:$(go env GOPATH)/bin"
func (*Server) GetAllRecipes(ctx context.Context, r *recipe.GetAllRecipesRequest) (*recipe.GetAllRecipesResponse, error) {
	recipes, count, err := models.GetAllRecipes(r.GetPageNumber(), r.GetResultPerPage())
	if err != nil {
		plugins.LogError("API Service", "Error at: GetAllRecipes", err)
		return &recipe.GetAllRecipesResponse{}, err
	}
	var arrayOfRecipes []*recipe.GetRecipeResponse
	copier.Copy(&arrayOfRecipes, recipes)
	for i := 0; i < len(arrayOfRecipes); i++ {
		createdAt, err := ptypes.TimestampProto(recipes[i].CreatedAt)
		if err != nil {
			plugins.LogError("API Service", "Error at: GetAllRecipes - time.Time to Proto TimeStamp conversion err", err)
			return &recipe.GetAllRecipesResponse{}, err
		}
		arrayOfRecipes[i].CreatedAt = createdAt
	}
	return &recipe.GetAllRecipesResponse{
		PageNumber:    r.GetPageNumber(),
		ResultPerPage: r.GetResultPerPage(),
		TotalCount:    count,
		Recipes:       arrayOfRecipes,
	}, nil
}

//CreateRecipe - Create new Recipe
func (*Server) CreateRecipe(ctx context.Context, r *recipe.CreateRecipeRequest) (*recipe.CreateRecipeResponse, error) {
	var newRecipe models.Recipe
	copier.Copy(&newRecipe, r)
	createdRecipe, err := models.CreateRecipe(&newRecipe)
	if err != nil {
		plugins.LogError("API Service", "Error at: CreateRecipe", err)
		return &recipe.CreateRecipeResponse{}, err
	}
	var test recipe.CreateRecipeResponse
	copier.Copy(&test, createdRecipe)
	return &test, nil
}

//GetRecipe - Get Recipe by ID
func (*Server) GetRecipe(ctx context.Context, r *recipe.GetRecipeRequest) (*recipe.GetRecipeResponse, error) {
	singleRecipe, err := models.GetRecipe(r.ID)
	if err != nil {
		plugins.LogError("API Service", "Error at: GetRecipe", err)
		return &recipe.GetRecipeResponse{}, err
	}
	var test recipe.GetRecipeResponse
	copier.Copy(&test, singleRecipe)
	createdAt, err := ptypes.TimestampProto(singleRecipe.CreatedAt)
	if err != nil {
		plugins.LogError("API Service", "Error at: GetRecipe - time.Time to Proto TimeStamp conversion err", err)
		return &recipe.GetRecipeResponse{}, err
	}
	test.CreatedAt = createdAt
	return &test, nil
}

//UpdateRecipe - Update Recipe by ID
func (*Server) UpdateRecipe(ctx context.Context, r *recipe.UpdateRecipeRequest) (*recipe.UpdateRecipeResponse, error) {
	var newRecipe models.Recipe
	copier.Copy(&newRecipe, r)
	if err := models.UpdateRecipe(newRecipe, r.ID); err != nil {
		plugins.LogError("API Service", "Error at: UpdateRecipe", err)
		return &recipe.UpdateRecipeResponse{}, err
	}
	return &recipe.UpdateRecipeResponse{
		Message: "Recipe Updated successfully",
	}, nil
}

//DeleteRecipe - Delete Recipe by ID
func (*Server) DeleteRecipe(ctx context.Context, r *recipe.DeleteRecipeRequest) (*recipe.DeleteRecipeResponse, error) {
	if err := models.DeleteRecipe(r.ID); err != nil {
		plugins.LogError("API Service", "Error at: DeleteRecipe", err)
		return &recipe.DeleteRecipeResponse{}, err
	}
	return &recipe.DeleteRecipeResponse{}, nil
}

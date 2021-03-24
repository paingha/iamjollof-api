// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package user

import (
	"context"
	"errors"
	"log"

	"bitbucket.com/iamjollof/server/models"
	"bitbucket.com/iamjollof/server/protos/user"
	"bitbucket.com/iamjollof/server/security"
	"github.com/jinzhu/copier"
)

//Server - Struct holding Users
type Server struct {
	user.UserServer
}

//RegisterUser - Create new User Account
func (*Server) RegisterUser(ctx context.Context, r *user.RegisterUserRequest) (*user.RegisterUserResponse, error) {
	var newUser models.User
	copier.Copy(&newUser, r)
	createdUser, err := models.CreateUser(&newUser)
	if err != nil {
		log.Printf("An error occured: %v", err)
		return &user.RegisterUserResponse{}, err
	}
	var test user.RegisterUserResponse
	copier.Copy(&test, createdUser)
	return &test, nil
}

//LoginUser - log user into account
func (*Server) LoginUser(ctx context.Context, r *user.LoginUserRequest) (*user.LoginUserResponse, error) {
	var newUser models.User
	copier.Copy(&newUser, r)
	userData, token, err := models.LoginUser(&newUser)
	if err != nil {
		log.Printf("An error occured: %v", err)
		return &user.LoginUserResponse{}, err
	}
	return &user.LoginUserResponse{
		Token:     token,
		ID:        userData.ID,
		FirstName: userData.FirstName,
		LastName:  userData.LastName,
		Email:     userData.Email,
	}, nil
}

//ForgotUser - User Account Password Reset
func (*Server) ForgotUser(ctx context.Context, r *user.ForgotUserRequest) (*user.EmptyUserResponse, error) {
	log.Printf("Receive message body from client: %s", r.Email)
	var newUser models.User
	copier.Copy(&newUser, r)
	_, err := models.ForgotUser(&newUser)
	if err != nil {
		log.Printf("An error occured: %v", err)
		return &user.EmptyUserResponse{}, err
	}
	return &user.EmptyUserResponse{}, nil
}

//GetAllUsers - Lists all the users
//Todo: Handle errors
func (*Server) GetAllUsers(ctx context.Context, r *user.GetAllUsersRequest) (*user.GetAllUsersResponse, error) {
	users, count, err := models.GetAllUsers(r.GetPageNumber(), r.GetResultPerPage())
	if err != nil {
		log.Printf("An error occured: %v", err)
		return &user.GetAllUsersResponse{}, err
	}
	var data []*user.GetUserResponse
	copier.Copy(&data, users)
	return &user.GetAllUsersResponse{
		PageNumber:    r.GetPageNumber(),
		ResultPerPage: r.GetResultPerPage(),
		TotalCount:    count,
		Users:         data,
	}, nil
}

//GetUser - Get User by ID
func (*Server) GetUser(ctx context.Context, r *user.GetUserRequest) (*user.GetUserResponse, error) {
	singleUser, err := models.GetUser(r.ID)
	if err != nil {
		log.Printf("An error occured: %v", err)
		return &user.GetUserResponse{}, err
	}
	var test user.GetUserResponse
	copier.Copy(&test, singleUser)
	return &test, nil
}

//UpdateUser - Update User by ID
func (*Server) UpdateUser(ctx context.Context, r *user.UpdateUserRequest) (*user.UpdateUserResponse, error) {
	log.Printf("Receive message body from client: %s", r.Email)
	var test models.User
	copier.Copy(&test, r)
	if err := models.UpdateUser(&test, r.ID); err != nil {
		log.Printf("An error occured: %v", err)
		return &user.UpdateUserResponse{
			Message: "An error occured",
		}, err
	}
	return &user.UpdateUserResponse{
		Message: "User successfully updated",
	}, nil
}

//VerifyTokenUser - Verifies if Token is expired or invalid
func (*Server) VerifyTokenUser(ctx context.Context, r *user.VerifyTokenUserRequest) (*user.VerifyTokenUserResponse, error) {
	log.Printf("Receive message body from client: %s", r.Token)
	if _, verify := security.VerifyJWT(r.Token); !verify {
		return &user.VerifyTokenUserResponse{
			Message: "Token Expired or invalid",
			Status:  false,
		}, errors.New("Token Expired or invalid")
	}
	return &user.VerifyTokenUserResponse{
		Message: "Token is valid",
		Status:  true,
	}, nil
}

//DeleteUser - Delete User by ID
func (*Server) DeleteUser(ctx context.Context, r *user.DeleteUserRequest) (*user.EmptyUserResponse, error) {
	if err := models.DeleteUser(r.ID); err != nil {
		log.Printf("An error occured: %v", err)
		return &user.EmptyUserResponse{}, err
	}
	return &user.EmptyUserResponse{}, nil
}

//VerifyEmailUser - Verifies User Account by Email
func (*Server) VerifyEmailUser(ctx context.Context, r *user.VerifyEmailUserRequest) (*user.VerifyEmailUserResponse, error) {
	log.Printf("Receive message body from client: %s", r.Code)
	if err := models.VerifyEmailUser(r.Code); err != nil {
		log.Printf("An error occured: %v", err)
		return &user.VerifyEmailUserResponse{}, err
	}
	return &user.VerifyEmailUserResponse{
		Message: "Email Verified Successfully",
	}, nil
}

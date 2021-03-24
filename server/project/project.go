// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package project

import (
	"context"

	"bitbucket.com/iamjollof/server/models"
	"bitbucket.com/iamjollof/server/plugins"
	"bitbucket.com/iamjollof/server/protos/project"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/copier"
)

//Server - Struct holding Projects
type Server struct {
	project.ProjectServer
}

//GetAllProjects - Lists all the projects
//Todo: Handle errors
//sudo service postgresql start
//export PATH="$PATH:$(go env GOPATH)/bin"
func (*Server) GetAllProjects(ctx context.Context, r *project.GetAllProjectsRequest) (*project.GetAllProjectsResponse, error) {
	projects, count, err := models.GetAllProjects(r.GetPageNumber(), r.GetResultPerPage())
	if err != nil {
		plugins.LogError("API Service", "Error at: GetAllProjects", err)
		return &project.GetAllProjectsResponse{}, err
	}
	var arrayOfProjects []*project.GetProjectResponse
	copier.Copy(&arrayOfProjects, projects)
	for i := 0; i < len(arrayOfProjects); i++ {
		createdAt, err := ptypes.TimestampProto(projects[i].CreatedAt)
		if err != nil {
			plugins.LogError("API Service", "Error at: GetAllProjects - time.Time to Proto TimeStamp conversion err", err)
			return &project.GetAllProjectsResponse{}, err
		}
		arrayOfProjects[i].CreatedAt = createdAt
	}
	return &project.GetAllProjectsResponse{
		PageNumber:    r.GetPageNumber(),
		ResultPerPage: r.GetResultPerPage(),
		TotalCount:    count,
		Projects:      arrayOfProjects,
	}, nil
}

//CreateProject - Create new Project
func (*Server) CreateProject(ctx context.Context, r *project.CreateProjectRequest) (*project.CreateProjectResponse, error) {
	var newProject models.Project
	copier.Copy(&newProject, r)
	createdProject, err := models.CreateProject(&newProject)
	if err != nil {
		plugins.LogError("API Service", "Error at: CreateProject", err)
		return &project.CreateProjectResponse{}, err
	}
	var test project.CreateProjectResponse
	copier.Copy(&test, createdProject)
	return &test, nil
}

//GetProject - Get Project by ID
func (*Server) GetProject(ctx context.Context, r *project.GetProjectRequest) (*project.GetProjectResponse, error) {
	singleProject, err := models.GetProject(r.ID)
	if err != nil {
		plugins.LogError("API Service", "Error at: GetProject", err)
		return &project.GetProjectResponse{}, err
	}
	var test project.GetProjectResponse
	copier.Copy(&test, singleProject)
	createdAt, err := ptypes.TimestampProto(singleProject.CreatedAt)
	if err != nil {
		plugins.LogError("API Service", "Error at: GetProject - time.Time to Proto TimeStamp conversion err", err)
		return &project.GetProjectResponse{}, err
	}
	test.CreatedAt = createdAt
	return &test, nil
}

//UpdateProject - Update Project by ID
func (*Server) UpdateProject(ctx context.Context, r *project.UpdateProjectRequest) (*project.UpdateProjectResponse, error) {
	var newProject models.Project
	copier.Copy(&newProject, r)
	if err := models.UpdateProject(newProject, r.ID); err != nil {
		plugins.LogError("API Service", "Error at: UpdateProject", err)
		return &project.UpdateProjectResponse{}, err
	}
	return &project.UpdateProjectResponse{
		Message: "Project Updated successfully",
	}, nil
}

//DeleteProject - Delete Project by ID
func (*Server) DeleteProject(ctx context.Context, r *project.DeleteProjectRequest) (*project.DeleteProjectResponse, error) {
	if err := models.DeleteProject(r.ID); err != nil {
		plugins.LogError("API Service", "Error at: DeleteProject", err)
		return &project.DeleteProjectResponse{}, err
	}
	return &project.DeleteProjectResponse{}, nil
}

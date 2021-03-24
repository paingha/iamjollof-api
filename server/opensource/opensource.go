// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package opensource

import (
	"context"

	"bitbucket.com/iamjollof/server/models"
	"bitbucket.com/iamjollof/server/plugins"
	"bitbucket.com/iamjollof/server/protos/opensource"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/copier"
)

//Server - Struct holding OpenSources
type Server struct {
	opensource.OpensourceServer
}

//GetAllOpenSources - Lists all the opensources
//Todo: Handle errors
//sudo service postgresql start
//export PATH="$PATH:$(go env GOPATH)/bin"
func (*Server) GetAllOpenSources(ctx context.Context, r *opensource.GetAllOpensourcesRequest) (*opensource.GetAllOpensourcesResponse, error) {
	opensources, count, err := models.GetAllOpenSources(r.GetPageNumber(), r.GetResultPerPage())
	if err != nil {
		plugins.LogError("API Service", "Error at: GetAllOpenSources", err)
		return &opensource.GetAllOpensourcesResponse{}, err
	}
	var arrayOfOpenSources []*opensource.GetOpensourceResponse
	copier.Copy(&arrayOfOpenSources, opensources)
	for i := 0; i < len(arrayOfOpenSources); i++ {
		createdAt, err := ptypes.TimestampProto(opensources[i].CreatedAt)
		if err != nil {
			plugins.LogError("API Service", "Error at: GetAllOpenSources - time.Time to Proto TimeStamp conversion err", err)
			return &opensource.GetAllOpensourcesResponse{}, err
		}
		arrayOfOpenSources[i].CreatedAt = createdAt
	}
	return &opensource.GetAllOpensourcesResponse{
		PageNumber:    r.GetPageNumber(),
		ResultPerPage: r.GetResultPerPage(),
		TotalCount:    count,
		Opensources:   arrayOfOpenSources,
	}, nil
}

//CreateOpenSource - Create new OpenSource
func (*Server) CreateOpenSource(ctx context.Context, r *opensource.CreateOpensourceRequest) (*opensource.CreateOpensourceResponse, error) {
	var newOpenSource models.OpenSource
	copier.Copy(&newOpenSource, r)
	createdOpenSource, err := models.CreateOpenSource(&newOpenSource)
	if err != nil {
		plugins.LogError("API Service", "Error at: CreateOpenSource", err)
		return &opensource.CreateOpensourceResponse{}, err
	}
	var test opensource.CreateOpensourceResponse
	copier.Copy(&test, createdOpenSource)
	return &test, nil
}

//GetOpenSource - Get OpenSource by ID
func (*Server) GetOpenSource(ctx context.Context, r *opensource.GetOpensourceRequest) (*opensource.GetOpensourceResponse, error) {
	singleOpenSource, err := models.GetOpenSource(r.ID)
	if err != nil {
		plugins.LogError("API Service", "Error at: GetOpenSource", err)
		return &opensource.GetOpensourceResponse{}, err
	}
	var test opensource.GetOpensourceResponse
	copier.Copy(&test, singleOpenSource)
	createdAt, err := ptypes.TimestampProto(singleOpenSource.CreatedAt)
	if err != nil {
		plugins.LogError("API Service", "Error at: GetOpenSource - time.Time to Proto TimeStamp conversion err", err)
		return &opensource.GetOpensourceResponse{}, err
	}
	test.CreatedAt = createdAt
	return &test, nil
}

//UpdateOpenSource - Update OpenSource by ID
func (*Server) UpdateOpenSource(ctx context.Context, r *opensource.UpdateOpensourceRequest) (*opensource.UpdateOpensourceResponse, error) {
	var newOpenSource models.OpenSource
	copier.Copy(&newOpenSource, r)
	if err := models.UpdateOpenSource(newOpenSource, r.ID); err != nil {
		plugins.LogError("API Service", "Error at: UpdateOpenSource", err)
		return &opensource.UpdateOpensourceResponse{}, err
	}
	return &opensource.UpdateOpensourceResponse{
		Message: "OpenSource Updated successfully",
	}, nil
}

//DeleteOpenSource - Delete OpenSource by ID
func (*Server) DeleteOpenSource(ctx context.Context, r *opensource.DeleteOpensourceRequest) (*opensource.DeleteOpensourceResponse, error) {
	if err := models.DeleteOpenSource(r.ID); err != nil {
		plugins.LogError("API Service", "Error at: DeleteOpenSource", err)
		return &opensource.DeleteOpensourceResponse{}, err
	}
	return &opensource.DeleteOpensourceResponse{}, nil
}

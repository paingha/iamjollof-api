// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blog

import (
	"context"

	"bitbucket.com/iamjollof/server/models"
	"bitbucket.com/iamjollof/server/plugins"
	"bitbucket.com/iamjollof/server/protos/blog"
	"github.com/golang/protobuf/ptypes"
	"github.com/jinzhu/copier"
)

//Server - Struct holding Blogs
type Server struct {
	blog.BlogServer
}

//GetAllBlogs - Lists all the blogs
//Todo: Handle errors
//sudo service postgresql start
//export PATH="$PATH:$(go env GOPATH)/bin"
func (*Server) GetAllBlogs(ctx context.Context, r *blog.GetAllBlogsRequest) (*blog.GetAllBlogsResponse, error) {
	blogs, count, err := models.GetAllBlogs(r.GetPageNumber(), r.GetResultPerPage())
	if err != nil {
		plugins.LogError("API Service", "Error at: GetAllBlogs", err)
		return &blog.GetAllBlogsResponse{}, err
	}
	var arrayOfBlogs []*blog.GetBlogResponse
	copier.Copy(&arrayOfBlogs, blogs)
	for i := 0; i < len(arrayOfBlogs); i++ {
		createdAt, err := ptypes.TimestampProto(blogs[i].CreatedAt)
		if err != nil {
			plugins.LogError("API Service", "Error at: GetAllBlogs - time.Time to Proto TimeStamp conversion err", err)
			return &blog.GetAllBlogsResponse{}, err
		}
		arrayOfBlogs[i].CreatedAt = createdAt
	}
	return &blog.GetAllBlogsResponse{
		PageNumber:    r.GetPageNumber(),
		ResultPerPage: r.GetResultPerPage(),
		TotalCount:    count,
		Blogs:         arrayOfBlogs,
	}, nil
}

//CreateBlog - Create new Blog
func (*Server) CreateBlog(ctx context.Context, r *blog.CreateBlogRequest) (*blog.CreateBlogResponse, error) {
	var newBlog models.Blog
	copier.Copy(&newBlog, r)
	createdBlog, err := models.CreateBlog(&newBlog)
	if err != nil {
		plugins.LogError("API Service", "Error at: CreateBlog", err)
		return &blog.CreateBlogResponse{}, err
	}
	var test blog.CreateBlogResponse
	copier.Copy(&test, createdBlog)
	return &test, nil
}

//GetBlog - Get Blog by ID
func (*Server) GetBlog(ctx context.Context, r *blog.GetBlogRequest) (*blog.GetBlogResponse, error) {
	singleBlog, err := models.GetBlog(r.ID)
	if err != nil {
		plugins.LogError("API Service", "Error at: GetBlog", err)
		return &blog.GetBlogResponse{}, err
	}
	var test blog.GetBlogResponse
	copier.Copy(&test, singleBlog)
	createdAt, err := ptypes.TimestampProto(singleBlog.CreatedAt)
	if err != nil {
		plugins.LogError("API Service", "Error at: GetBlog - time.Time to Proto TimeStamp conversion err", err)
		return &blog.GetBlogResponse{}, err
	}
	test.CreatedAt = createdAt
	return &test, nil
}

//UpdateBlog - Update Blog by ID
func (*Server) UpdateBlog(ctx context.Context, r *blog.UpdateBlogRequest) (*blog.UpdateBlogResponse, error) {
	var newBlog models.Blog
	copier.Copy(&newBlog, r)
	if err := models.UpdateBlog(newBlog, r.ID); err != nil {
		plugins.LogError("API Service", "Error at: UpdateBlog", err)
		return &blog.UpdateBlogResponse{}, err
	}
	return &blog.UpdateBlogResponse{
		Message: "Blog Updated successfully",
	}, nil
}

//DeleteBlog - Delete Blog by ID
func (*Server) DeleteBlog(ctx context.Context, r *blog.DeleteBlogRequest) (*blog.DeleteBlogResponse, error) {
	if err := models.DeleteBlog(r.ID); err != nil {
		plugins.LogError("API Service", "Error at: DeleteBlog", err)
		return &blog.DeleteBlogResponse{}, err
	}
	return &blog.DeleteBlogResponse{}, nil
}

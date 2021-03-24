// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"net"

	blogServer "bitbucket.com/iamjollof/server/blog"
	"bitbucket.com/iamjollof/server/config"
	contactServer "bitbucket.com/iamjollof/server/contact"
	"bitbucket.com/iamjollof/server/interceptors"
	"bitbucket.com/iamjollof/server/models"
	opensourceServer "bitbucket.com/iamjollof/server/opensource"
	"bitbucket.com/iamjollof/server/plugins"
	projectServer "bitbucket.com/iamjollof/server/project"
	"bitbucket.com/iamjollof/server/protos/blog"
	"bitbucket.com/iamjollof/server/protos/contact"
	"bitbucket.com/iamjollof/server/protos/opensource"
	"bitbucket.com/iamjollof/server/protos/project"
	"bitbucket.com/iamjollof/server/protos/quote"
	"bitbucket.com/iamjollof/server/protos/recipe"
	"bitbucket.com/iamjollof/server/protos/user"
	quoteServer "bitbucket.com/iamjollof/server/quote"
	recipeServer "bitbucket.com/iamjollof/server/recipe"
	userServer "bitbucket.com/iamjollof/server/user"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/credentials"
)

func main() {
	var errs error
	config.LoadEnvFile()
	systemCfg := &config.SystemConfig{}
	if err := config.InitConfig(systemCfg); err != nil {
		plugins.LogFatal("gRPC Server", "Environment Variables error", err)
	}
	// Connect to Database
	config.DB, errs = gorm.Open("postgres", config.GetConnectionContext())
	if errs != nil {
		plugins.LogFatal("gRPC Server", "Database connection error", errs)
	}
	defer config.DB.Close()
	config.DB.LogMode(true)

	//Run Database Migration here
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Recipe{})
	config.DB.AutoMigrate(&models.Quote{})
	config.DB.AutoMigrate(&models.OpenSource{})
	config.DB.AutoMigrate(&models.Project{})
	config.DB.AutoMigrate(&models.Blog{})
	/*creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		plugins.LogFatal("gRPC Server", "Failed to generate ssl credentials", err)
	}*/
	app, err := net.Listen("tcp", ":9000")
	if err != nil {
		plugins.LogFatal("gRPC Server", "An error occured ", err)
	}
	plugins.LogInfo("gRPC Server", "Running gRPC Server...")
	grpcServer := grpc.NewServer(
		/*grpc.Creds(creds)*/
		grpc.UnaryInterceptor(interceptors.AuthUnary()),
		grpc.StreamInterceptor(interceptors.AuthStream()),
	)
	contact.RegisterContactServer(grpcServer, &contactServer.Server{})
	recipe.RegisterRecipeServer(grpcServer, &recipeServer.Server{})
	quote.RegisterQuoteServer(grpcServer, &quoteServer.Server{})
	blog.RegisterBlogServer(grpcServer, &blogServer.Server{})
	opensource.RegisterOpensourceServer(grpcServer, &opensourceServer.Server{})
	project.RegisterProjectServer(grpcServer, &projectServer.Server{})
	user.RegisterUserServer(grpcServer, &userServer.Server{})
	if err := grpcServer.Serve(app); err != nil {
		plugins.LogFatal("gRPC Server", "failed to serve: ", err)
	}
}

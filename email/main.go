// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"net"

	"bitbucket.com/iamjollof/email/config"
	"bitbucket.com/iamjollof/email/models"
	"bitbucket.com/iamjollof/email/plugins"
	"bitbucket.com/iamjollof/email/protos/email"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"
)

//SystemCfg - System configuration with environment variables
var SystemCfg = &config.SystemConfig{}

//Server - Struct holding Sermons
type Server struct {
	email.EmailServer
}

//SendEmail - Create new Sermon
func (*Server) SendEmail(ctx context.Context, r *email.SendEmailRequest) (*email.SendEmailResponse, error) {
	var newEmail models.EmailParam
	copier.Copy(&newEmail, r)
	if err := SendEmail(newEmail, SystemCfg); err != nil {
		plugins.LogError("MailService", "an error occured", err)
		return &email.SendEmailResponse{
			Message: "An error occured while sending Email",
		}, err
	}
	return &email.SendEmailResponse{
		Message: "Email sent successfully",
	}, nil
}

//TODO: Implement SendManyEmails - gRPC stream to send emails to numerous emails

func main() {
	files := map[string]string{
		"TemplateVerifyEmail": "./templates/verify.html",
		"TemplateResetEmail":  "./templates/password-reset.html",
		"TemplateContactUs":   "./templates/contact-us.html",
	}
	getFilesContents(files)
	config.LoadEnvFile()
	SystemCfg = &config.SystemConfig{}
	if err := config.InitConfig(SystemCfg); err != nil {
		plugins.LogFatal("gRPC Email Service", "Email Service Environment Variables error", err)
	}
	app, err := net.Listen("tcp", ":9001")
	if err != nil {
		plugins.LogFatal("gRPC Email Service", "An error occured ", err)
	}
	plugins.LogInfo("gRPC Email Service", "Running gRPC Email Service...")
	grpcServer := grpc.NewServer()
	email.RegisterEmailServer(grpcServer, &Server{})
	if err := grpcServer.Serve(app); err != nil {
		plugins.LogFatal("gRPC Email Service", "Email Service failed to serve: ", err)
	}

}

package models

import (
	"context"

	"bitbucket.com/iamjollof/server/plugins"
	"bitbucket.com/iamjollof/server/protos/email"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc"

	//Needed for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//EmailParam - email service sending structure
type EmailParam struct {
	Template  string            `json:"template"`
	To        string            `json:"to"`
	Subject   string            `json:"subject"`
	BodyParam map[string]string `json:"body_param"`
}

//SendMail - gRPC client that sends email message to email service
func SendMail(emailData EmailParam) error {
	conn, err := grpc.Dial("pace-email-1:9001", grpc.WithInsecure())
	if err != nil {
		plugins.LogError("gRPC Server internal Client", "did not connect", err)
		return err
	}
	defer conn.Close()
	c := email.NewEmailClient(conn)
	var data email.SendEmailRequest
	copier.Copy(&data, emailData)
	if _, err := c.SendEmail(context.Background(), &data); err != nil {
		plugins.LogWarning("gRPC Server internal Client", "Error when calling Send Email", err)
		return err
	}
	return nil
}

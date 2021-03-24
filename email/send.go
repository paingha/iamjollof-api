package main

import (
	"bytes"
	"fmt"
	"html/template"

	"bitbucket.com/iamjollof/email/config"
	"bitbucket.com/iamjollof/email/models"
	"bitbucket.com/iamjollof/email/plugins"

	sendgrid "github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//SendEmail - send email through Sendgrid
func SendEmail(emailParam models.EmailParam, cfg *config.SystemConfig) error {
	var bodyTemplate string
	plugins.LogInfo("gRPC Email Service: Sending mail to ", emailParam.To)
	if t, ok := EmailTemplates[emailParam.Template]; ok {
		bodyTemplate = t
	} else {
		bodyTemplate = EmailTemplates["TemplateVerifyEmail"]
	}
	parsedTemplate, err := template.New("template").Parse(bodyTemplate)
	if err != nil {
		plugins.LogError("gRPC Email Service", "error parsing email template", err)
		return err
	}
	var buf bytes.Buffer
	err = parsedTemplate.Execute(&buf, emailParam.BodyParam)

	from := mail.NewEmail(cfg.SenderName, cfg.SenderEmail)
	subject := emailParam.Subject
	to := mail.NewEmail(fmt.Sprintf("%s %s ", emailParam.BodyParam["first_name"], emailParam.BodyParam["last_name"]), emailParam.To)

	htmlContent := buf.String()
	message := mail.NewSingleEmail(from, subject, to, "text", htmlContent)

	client := sendgrid.NewSendClient(cfg.SendgridAPIKey)
	response, err := client.Send(message)
	fmt.Println(response.StatusCode)
	if err != nil {
		plugins.LogError("gRPC Email Service", "error sending email", err)
		return err
	}
	plugins.LogInfo("gRPC Email Service: Email sent successfully", "200")
	return nil
}

// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package contact

import (
	"context"
	"log"

	"bitbucket.com/iamjollof/server/models"
	"bitbucket.com/iamjollof/server/protos/contact"
)

//Server - Struct holding Contacts
type Server struct {
	contact.ContactServer
}

//CreateContact - Create new Contact
func (*Server) CreateContact(ctx context.Context, r *contact.CreateContactRequest) (*contact.CreateContactResponse, error) {
	log.Printf("Receive message body from client: %s", r.Name)
	emailBody := map[string]string{
		"first_name": "Go Pace NG",
		"last_name":  "Contact Us Form",
		"name":       r.Name,
		"email":      r.Email,
		"phone":      r.Phone,
		"reason":     r.ReasonForEmail,
		"message":    r.Content,
	}
	emailInfo := models.EmailParam{
		To:        "inquiries@gopace.xyz",
		Subject:   "Go Pace Contact Us Email from " + r.Name + " Issue " + r.ReasonForEmail,
		BodyParam: emailBody,
		Template:  "TemplateContactUs",
	}
	if err := models.SendMail(emailInfo); err != nil {
		return &contact.CreateContactResponse{}, err
	}
	return &contact.CreateContactResponse{
		Message: "Message sent successfully",
	}, nil
}

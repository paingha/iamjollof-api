// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"log"

	"github.com/joho/godotenv"

	env "github.com/Netflix/go-env"
)

//SystemConfig represents system service configuration
type SystemConfig struct {
	SenderName     string `env:"ENV_EMAIL_SENDER_NAME"`
	SenderEmail    string `env:"ENV_SENDER_EMAIL"`
	SendgridAPIKey string `env:"ENV_SENDGRID_API_KEY"`
	SentryUrl      string `env:"ENV_SENTRY_URL"`
}

//InitConfig - initial the configuration struct with environment variables
func InitConfig(cfg interface{}) error {
	_, err := env.UnmarshalFromEnviron(cfg)
	if err != nil {
		return err
	}
	return nil
}

//LoadEnvFile - loads env file
func LoadEnvFile() {
	if err := godotenv.Load("example.env"); err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
}

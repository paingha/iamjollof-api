// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plugins

import (
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"go.uber.org/zap"
)

var logger *zap.Logger
var err error

func init() {
	logger, err = zap.NewProduction()
	if err != nil {
		log.Fatalf("error initializing zap logger: %v", err)
	}
	errs := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("ENV_SENTRY_URL"),
	})
	if errs != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}

//LogInfo - logs information message to stdout
func LogInfo(name, message string) {
	defer logger.Sync()
	defer sentry.Flush(2 * time.Second)
	logger.Info(message,
		zap.Reflect("service-name:", name),
	)
	sentry.CaptureMessage("service-name: " + name + ", message: " + message)
}

//LogWarning - logs warning information message to stdout
func LogWarning(name, message string, err error) {
	defer logger.Sync()
	defer sentry.Flush(2 * time.Second)
	logger.Warn(message,
		zap.String("service-name:", name),
		zap.String("verbose:", err.Error()),
		zap.Reflect("error:", err),
	)
	sentry.CaptureMessage("service-name: " + name + ", type: log-warning" + ", message: " + message + ", verbose:" + err.Error())
}

//LogError - logs error message to stdout
func LogError(name, message string, err error) {
	defer logger.Sync()
	defer sentry.Flush(2 * time.Second)
	logger.Error(message,
		zap.String("service-name:", name),
		zap.String("verbose:", err.Error()),
		zap.Reflect("error:", err),
	)
	sentry.CaptureMessage("service-name: " + name + ", type: log-error" + ", message: " + message + ", verbose:" + err.Error())
}

//LogPanic - logs error message to stdout and panics
func LogPanic(name, message string, err error) {
	defer logger.Sync()
	defer sentry.Flush(2 * time.Second)
	logger.Panic(message,
		zap.String("service-name:", name),
		zap.String("verbose:", err.Error()),
		zap.Reflect("error:", err),
	)
	sentry.CaptureMessage("service-name: " + name + ", type: log-panic" + ", message: " + message + ", verbose:" + err.Error())
}

//LogFatal - logs error message to stdout and panics and calls os.Exit(1)
func LogFatal(name, message string, err error) {
	defer logger.Sync()
	defer sentry.Flush(2 * time.Second)
	logger.Fatal(message,
		zap.String("service-name:", name),
		zap.String("verbose:", err.Error()),
		zap.Reflect("error:", err),
	)
	sentry.CaptureMessage("service-name: " + name + ", type: log-fatal" + ", message: " + message + ", verbose:" + err.Error())
}

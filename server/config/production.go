// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"

	"bitbucket.com/iamjollof/server/plugins"
	"bitbucket.com/iamjollof/server/utils"
)

//BuildProdDBConfig - Builds DB Config for production environment
func BuildProdDBConfig() *DBConfig {
	var cfg SystemConfig
	err := InitConfig(&cfg)
	if err != nil {
		plugins.LogFatal("API", "Wrong Prod System config", err)
	}
	port, _ := utils.ConvertStringToInt(cfg.DevDBPort)
	dbConfig := DBConfig{
		Host:     cfg.ProdDBHost,
		Port:     port,
		User:     cfg.ProdDBUser,
		DBName:   cfg.ProdDBDatabase,
		Password: cfg.ProdDBPass,
		SSL:      cfg.ProdDBSSL,
	}
	return &dbConfig
}

//ProdDbURL - returns connection string for production database
func ProdDbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
		dbConfig.SSL,
	)
}

// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils_test

import (
	"testing"

	"bitbucket.com/iamjollof/server/utils"
	"github.com/stretchr/testify/assert"
)

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

func TestValidateStruct(T *testing.T) {
	address := &Address{
		Street: "Eavesdown Docks",
		City:   "hello",
		Planet: "Persphone",
		Phone:  "none",
	}
	err := utils.ValidateStruct(address)
	assert.Nil(T, err)
}

func TestValidateVariable(T *testing.T) {
	email := "test@test.com"
	err := utils.ValidateVariable(email, "required,email")
	assert.Nil(T, err)
}

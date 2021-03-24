// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils_test

import (
	"testing"

	"bitbucket.com/iamjollof/server/utils"
	"github.com/stretchr/testify/assert"
)

func TestConvertstringtofloat(T *testing.T) {
	resp, err := utils.ConvertStringToFloat("3.1415")
	assert.Nil(T, err)
	assert.Equal(T, resp, 3.1415, "The returned float should be 3.1415")
}

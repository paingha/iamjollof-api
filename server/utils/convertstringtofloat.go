// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package utils

import (
	"strconv"
)

//ConvertStringToFloat - converts a string to an float
func ConvertStringToFloat(character string) (data float64, err error) {
	i, errs := strconv.ParseFloat(character, 64)
	return i, errs
}

// Copyright 2021 Paingha Joe Alagoa. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package models

//EmailParam - email service sending structure
type EmailParam struct {
	Template  string            `json:"template"`
	To        string            `json:"to"`
	Subject   string            `json:"subject"`
	BodyParam map[string]string `json:"body_param"`
}

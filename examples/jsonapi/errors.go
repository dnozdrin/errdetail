// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// license that can be found in the LICENSE file.

package jsonapi

import (
	"github.com/dnozdrin/errdetail"
)

// This package provides examples of formatting detailed errors
// according to the JSON:API Specification.
// For details see https://jsonapi.org/format/#error-objects.

type ErrorResponse struct {
	Errors []Error `json:"errors,omitempty"`
}

type Error struct {
	Code   string                 `json:"code,omitempty"`
	Title  string                 `json:"title,omitempty"`
	Status string                 `json:"status,omitempty"`
	Detail string                 `json:"detail,omitempty"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
}

func (err *Error) addMetaItem(name string, value interface{}) {
	if err.Meta == nil {
		err.Meta = make(map[string]interface{})
	}

	err.Meta[name] = value
}

func NewErrorResponse(err error) ErrorResponse {
	return ErrorResponse{
		Errors: toErrors(err),
	}
}

func toErrors(err error) []Error {
	extracted := errdetail.ExtractDetails(err)
	if len(extracted) == 0 {
		return nil
	}

	details := make([]Error, len(extracted))
	for i := range extracted {
		details[i] = Error{
			Code:   extracted[i].Code(),
			Title:  extracted[i].Description(),
			Status: toStatusCode[extracted[i].Code()],
			Detail: extracted[i].Reason(),
		}

		if field := extracted[i].Field(); field != "" {
			details[i].addMetaItem("field", field)
		}

		if domain := extracted[i].Domain(); domain != "" {
			details[i].addMetaItem("domain", domain)
		}
	}

	return details
}

var toStatusCode = map[string]string{
	"invalid_argument": "400",
	"invalid_email":    "400",
	"admin_required":   "403",
	"not_found":        "404",
	"internal":         "500",
}

// Copyright 2022 Dmytro Nozdrin. All rights reserved.
// Use of this source code is governed by the MIT License
// that can be found in the LICENSE file.

package errdetail

import (
	"errors"
)

// Detail represents a set of optional fields which provide more
// information about a parent error.
type Detail struct {
	field       string
	description string
	code        string
	domain      string
	reason      string
	filled      bool
}

// Field is a Detail field getter.
func (d Detail) Field() string {
	return d.field
}

// Description is a Detail description getter.
func (d Detail) Description() string {
	return d.description
}

// Code is a Detail code getter.
func (d Detail) Code() string {
	return d.code
}

// Domain is a Detail domain getter.
// Domain stands here for a specified sphere of activity or knowledge.
func (d Detail) Domain() string {
	return d.domain
}

// Reason is a Detail reason getter.
func (d Detail) Reason() string {
	return d.reason
}

// NewDetail represents a Detail constructor.
func NewDetail(opts ...Option) Detail {
	var detail Detail
	for i := range opts {
		opts[i](&detail)
	}

	return detail
}

// Option is a function type for Detail fields' setters.
type Option func(*Detail)

// WithField is an option for Detail constructs that sets field name
// and marks the detail as not empty.
func WithField(field string) Option {
	return func(d *Detail) {
		d.field = field
		if d.field != "" {
			d.filled = true
		}
	}
}

// WithDescription is an option for Detail constructs that sets
// description and marks the detail as not empty.
func WithDescription(description string) Option {
	return func(d *Detail) {
		d.description = description
		if d.description != "" {
			d.filled = true
		}
	}
}

// WithCode is an option for Detail constructs that sets code
// and marks the detail as not empty.
func WithCode(code string) Option {
	return func(d *Detail) {
		d.code = code
		if d.code != "" {
			d.filled = true
		}
	}
}

// WithDomain is an option for Detail constructs that sets an error
// domain and marks the detail as not empty. Domain stands here for
// a specified sphere of activity or knowledge.
func WithDomain(domain string) Option {
	return func(d *Detail) {
		d.domain = domain
		if d.domain != "" {
			d.filled = true
		}
	}
}

// WithReason is an option for Detail constructs that sets an error
// reason and marks the detail as not empty.
func WithReason(reason string) Option {
	return func(d *Detail) {
		d.reason = reason
		if d.reason != "" {
			d.filled = true
		}
	}
}

type detailed interface {
	Details() []Detail
}

// ExtractDetails extracts details from an error, if any. Otherwise, returns nil.
func ExtractDetails(err error) []Detail {
	var d detailed
	if errors.As(err, &d) {
		return d.Details()
	}

	return nil
}

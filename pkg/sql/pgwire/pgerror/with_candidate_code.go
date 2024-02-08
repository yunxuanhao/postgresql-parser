// Copyright 2019 The Cockroach Authors.
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package pgerror

import (
	"fmt"

	"github.com/cockroachdb/errors"
)

type withCandidateCode struct {
	cause error
	code  string
}

var _ error = (*withCandidateCode)(nil)
var _ errors.SafeDetailer = (*withCandidateCode)(nil)
var _ fmt.Formatter = (*withCandidateCode)(nil)
var _ errors.Formatter = (*withCandidateCode)(nil)

func (w *withCandidateCode) Error() string         { return w.cause.Error() }
func (w *withCandidateCode) Cause() error          { return w.cause }
func (w *withCandidateCode) Unwrap() error         { return w.cause }
func (w *withCandidateCode) SafeDetails() []string { return []string{w.code} }

func (w *withCandidateCode) Format(s fmt.State, verb rune) { errors.FormatError(w, s, verb) }

func (w *withCandidateCode) FormatError(p errors.Printer) (next error) {
	if p.Detail() {
		p.Printf("candidate pg code: %s", w.code)
	}
	return w.cause
}

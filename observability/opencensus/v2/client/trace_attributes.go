/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	"github.com/Yangfisher1/ce-go-sdk/v2/event"
	"github.com/Yangfisher1/ce-go-sdk/v2/observability"
	"github.com/Yangfisher1/opencensus-go/trace"
)

func EventTraceAttributes(e event.EventReader) []trace.Attribute {
	as := []trace.Attribute{
		trace.StringAttribute(observability.SpecversionAttr, e.SpecVersion()),
		trace.StringAttribute(observability.IdAttr, e.ID()),
		trace.StringAttribute(observability.TypeAttr, e.Type()),
		trace.StringAttribute(observability.SourceAttr, e.Source()),
	}
	if sub := e.Subject(); sub != "" {
		as = append(as, trace.StringAttribute(observability.SubjectAttr, sub))
	}
	if dct := e.DataContentType(); dct != "" {
		as = append(as, trace.StringAttribute(observability.DatacontenttypeAttr, dct))
	}
	return as
}

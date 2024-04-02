/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package expression

import (
	cesql "github.com/Yangfisher1/ce-go-sdk/sql/v2"
	cloudevents "github.com/Yangfisher1/ce-go-sdk/v2"
)

type literalExpression struct {
	value interface{}
}

func (l literalExpression) Evaluate(event cloudevents.Event) (interface{}, error) {
	return l.value, nil
}

func NewLiteralExpression(value interface{}) cesql.Expression {
	return literalExpression{value: value}
}

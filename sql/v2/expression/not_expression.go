/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package expression

import (
	cesql "github.com/Yangfisher1/ce-go-sdk/sql/v2"
	"github.com/Yangfisher1/ce-go-sdk/sql/v2/utils"
	cloudevents "github.com/Yangfisher1/ce-go-sdk/v2"
)

type notExpression baseUnaryExpression

func (l notExpression) Evaluate(event cloudevents.Event) (interface{}, error) {
	val, err := l.child.Evaluate(event)
	if err != nil {
		return nil, err
	}

	val, err = utils.Cast(val, cesql.BooleanType)
	if err != nil {
		return nil, err
	}

	return !(val.(bool)), nil
}

func NewNotExpression(child cesql.Expression) cesql.Expression {
	return notExpression{child: child}
}

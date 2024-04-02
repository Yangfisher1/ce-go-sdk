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

type inExpression struct {
	leftExpression cesql.Expression
	setExpression  []cesql.Expression
}

func (l inExpression) Evaluate(event cloudevents.Event) (interface{}, error) {
	leftValue, err := l.leftExpression.Evaluate(event)
	if err != nil {
		return nil, err
	}

	for _, rightExpression := range l.setExpression {
		rightValue, err := rightExpression.Evaluate(event)
		if err != nil {
			return nil, err
		}

		rightValue, err = utils.Cast(rightValue, cesql.TypeFromVal(leftValue))
		if err != nil {
			return nil, err
		}

		if leftValue == rightValue {
			return true, nil
		}
	}

	return false, nil
}

func NewInExpression(leftExpression cesql.Expression, setExpression []cesql.Expression) cesql.Expression {
	return inExpression{leftExpression, setExpression}
}

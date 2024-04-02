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

type existsExpression struct {
	identifier string
}

func (l existsExpression) Evaluate(event cloudevents.Event) (interface{}, error) {
	return utils.ContainsAttribute(event, l.identifier), nil
}

func NewExistsExpression(identifier string) cesql.Expression {
	return existsExpression{identifier: identifier}
}

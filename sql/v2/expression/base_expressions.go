/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package expression

import cesql "github.com/Yangfisher1/ce-go-sdk/sql/v2"

type baseUnaryExpression struct {
	child cesql.Expression
}

type baseBinaryExpression struct {
	left  cesql.Expression
	right cesql.Expression
}

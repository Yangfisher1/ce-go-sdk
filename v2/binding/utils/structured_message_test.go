/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package utils_test

import (
	"bytes"
	"context"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Yangfisher1/ce-go-sdk/v2/binding"
	"github.com/Yangfisher1/ce-go-sdk/v2/binding/format"
	"github.com/Yangfisher1/ce-go-sdk/v2/binding/utils"
	"github.com/Yangfisher1/ce-go-sdk/v2/test"
)

func TestNewStructuredMessage(t *testing.T) {
	testEvent := test.ConvertEventExtensionsToString(t, test.FullEvent())
	jsonBytes := test.MustJSON(t, testEvent)

	message := utils.NewStructuredMessage(format.JSON, ioutil.NopCloser(bytes.NewReader(jsonBytes)))

	require.Equal(t, binding.EncodingStructured, message.ReadEncoding())

	event := test.MustToEvent(t, context.TODO(), message)
	test.AssertEventEquals(t, testEvent, event)

	require.NoError(t, message.Finish(nil))
}

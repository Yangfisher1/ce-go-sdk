/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package buffering

import (
	"context"
	"testing"

	. "github.com/Yangfisher1/ce-go-sdk/v2/binding/test"
	. "github.com/Yangfisher1/ce-go-sdk/v2/test"
)

var err error

func BenchmarkBufferMessageFromStructured(b *testing.B) {
	e := FullEvent()
	input := MustCreateMockStructuredMessage(b, e)
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		outputMessage, _ := BufferMessage(ctx, input)
		err = outputMessage.Finish(nil)
	}
}

func BenchmarkBufferMessageFromBinary(b *testing.B) {
	e := FullEvent()
	input := MustCreateMockBinaryMessage(e)
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		outputMessage, _ := BufferMessage(ctx, input)
		err = outputMessage.Finish(nil)
	}
}

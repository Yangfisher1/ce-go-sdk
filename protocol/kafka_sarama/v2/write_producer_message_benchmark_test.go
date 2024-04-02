/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package kafka_sarama_test

import (
	"context"
	"testing"

	"github.com/Shopify/sarama"

	"github.com/Yangfisher1/ce-go-sdk/protocol/kafka_sarama/v2"
	"github.com/Yangfisher1/ce-go-sdk/v2/binding"
	"github.com/Yangfisher1/ce-go-sdk/v2/binding/format"
	bindingtest "github.com/Yangfisher1/ce-go-sdk/v2/binding/test"
	"github.com/Yangfisher1/ce-go-sdk/v2/event"
	"github.com/Yangfisher1/ce-go-sdk/v2/test"
)

// Avoid DCE
var ProducerMessage *sarama.ProducerMessage

var (
	ctx               context.Context
	initialEvent      event.Event
	structuredMessage binding.Message
	binaryMessage     binding.Message
)

func init() {
	ctx = context.TODO()

	initialEvent = test.FullEvent()

	structuredMessage = &bindingtest.MockStructuredMessage{
		Bytes: func() []byte {
			b, _ := format.JSON.Marshal(&testEvent)
			return b
		}(),
		Format: format.JSON,
	}
	binaryMessage = bindingtest.MustCreateMockBinaryMessage(initialEvent)
}

func BenchmarkEncodeStructuredMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProducerMessage = &sarama.ProducerMessage{}
		Err = kafka_sarama.WriteProducerMessage(ctx, structuredMessage, ProducerMessage)
	}
}

func BenchmarkEncodeBinaryMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProducerMessage = &sarama.ProducerMessage{}
		Err = kafka_sarama.WriteProducerMessage(ctx, binaryMessage, ProducerMessage)
	}
}

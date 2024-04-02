/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package kafka_sarama_test

import (
	"context"
	"testing"

	"github.com/Yangfisher1/ce-go-sdk/protocol/kafka_sarama/v2"
	"github.com/Yangfisher1/ce-go-sdk/v2/binding"
	"github.com/Yangfisher1/ce-go-sdk/v2/event"
)

// Avoid DCE
var M binding.Message
var Event *event.Event
var Err error

func BenchmarkNewStructuredMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		M = kafka_sarama.NewMessageFromConsumerMessage(structuredConsumerMessage)
	}
}

func BenchmarkNewBinaryMessage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		M = kafka_sarama.NewMessageFromConsumerMessage(binaryConsumerMessage)
	}
}

func BenchmarkNewStructuredMessageToEvent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		M = kafka_sarama.NewMessageFromConsumerMessage(structuredConsumerMessage)
		Event, Err = binding.ToEvent(context.TODO(), M)
	}
}

func BenchmarkNewBinaryMessageToEvent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		M = kafka_sarama.NewMessageFromConsumerMessage(binaryConsumerMessage)
		Event, Err = binding.ToEvent(context.TODO(), M)
	}
}

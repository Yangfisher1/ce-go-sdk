/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package nats_jetstream

import (
	"context"
	"io"

	"github.com/Yangfisher1/ce-go-sdk/v2/binding"
	"github.com/Yangfisher1/ce-go-sdk/v2/binding/format"
)

// WriteMsg fills the provided writer with the bindings.Message m.
// Using context you can tweak the encoding processing (more details on binding.Write documentation).
func WriteMsg(ctx context.Context, m binding.Message, writer io.ReaderFrom, transformers ...binding.Transformer) error {
	structuredWriter := &natsMessageWriter{writer}

	_, err := binding.Write(
		ctx,
		m,
		structuredWriter,
		nil,
		transformers...,
	)
	return err
}

type natsMessageWriter struct {
	io.ReaderFrom
}

// StructuredWriter  implements StructuredWriter.SetStructuredEvent
func (w *natsMessageWriter) SetStructuredEvent(_ context.Context, _ format.Format, event io.Reader) error {
	if _, err := w.ReadFrom(event); err != nil {
		return err
	}

	return nil
}

var _ binding.StructuredWriter = (*natsMessageWriter)(nil) // Test it conforms to the interface

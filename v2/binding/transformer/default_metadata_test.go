/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package transformer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/Yangfisher1/ce-go-sdk/v2/binding"
	. "github.com/Yangfisher1/ce-go-sdk/v2/binding/test"
	"github.com/Yangfisher1/ce-go-sdk/v2/event"
	. "github.com/Yangfisher1/ce-go-sdk/v2/test"
)

func TestAddTimeNow(t *testing.T) {
	eventWithoutTime := MinEvent()
	eventCtx := eventWithoutTime.Context.AsV1()
	eventCtx.Time = nil
	eventWithoutTime.Context = eventCtx

	eventWithTime := MinEvent()
	eventWithTime.SetTime(time.Now().Add(2 * time.Hour).UTC())

	assertTimeNow := func(t *testing.T, ev event.Event) {
		require.False(t, ev.Context.GetTime().IsZero())
	}

	RunTransformerTests(t, context.Background(), []TransformerTestArgs{
		{
			Name:         "No change to time to Mock Structured message",
			InputMessage: MustCreateMockStructuredMessage(t, eventWithTime.Clone()),
			WantEvent:    eventWithTime.Clone(),
			Transformers: binding.Transformers{AddTimeNow},
		},
		{
			Name:         "No change to time to Mock Binary message",
			InputMessage: MustCreateMockBinaryMessage(eventWithTime.Clone()),
			WantEvent:    eventWithTime.Clone(),
			Transformers: binding.Transformers{AddTimeNow},
		},
		{
			Name:         "No change to time to Event message",
			InputEvent:   eventWithTime,
			WantEvent:    eventWithTime,
			Transformers: binding.Transformers{AddTimeNow},
		},
		{
			Name:         "Add time.Now() to Mock Binary message",
			InputMessage: MustCreateMockBinaryMessage(eventWithoutTime.Clone()),
			AssertFunc:   assertTimeNow,
			Transformers: binding.Transformers{AddTimeNow},
		},
		{
			Name:         "Add time.Now() to Event message",
			InputEvent:   eventWithoutTime,
			AssertFunc:   assertTimeNow,
			Transformers: binding.Transformers{AddTimeNow},
		},
	})
}

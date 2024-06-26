/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

// Package test provides re-usable functions for binding tests.
package test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/Yangfisher1/ce-go-sdk/v2/binding"
	"github.com/Yangfisher1/ce-go-sdk/v2/protocol"
)

// SendReceive does s.Send(in), then it receives the message in r.Receive() and executes outAssert
// Halt test on error.
func SendReceive(t *testing.T, ctx context.Context, in binding.Message, s protocol.Sender, r protocol.Receiver, outAssert func(binding.Message)) {
	t.Helper()
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		out, result := r.Receive(ctx)
		if !protocol.IsACK(result) {
			require.NoError(t, result)
		}
		outAssert(out)
		finishErr := out.Finish(nil)
		require.NoError(t, finishErr)
	}()

	go func() {
		defer wg.Done()
		mx := sync.Mutex{}
		finished := false
		in = binding.WithFinish(in, func(err error) {
			require.NoError(t, err)
			mx.Lock()
			finished = true
			mx.Unlock()
		})
		result := s.Send(ctx, in)

		mx.Lock()
		if !protocol.IsACK(result) {
			require.NoError(t, result)
		}
		mx.Unlock()

		time.Sleep(5 * time.Millisecond) // let the receiver receive.
		mx.Lock()
		require.True(t, finished)
		mx.Unlock()
	}()

	wg.Wait()
}

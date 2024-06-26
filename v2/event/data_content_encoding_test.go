/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package event_test

import (
	"testing"

	"github.com/Yangfisher1/ce-go-sdk/v2/event"

	"github.com/google/go-cmp/cmp"
)

func TestStringOfBase64(t *testing.T) {
	want := strptr("base64")
	got := event.StringOfBase64()

	if diff := cmp.Diff(want, got); diff != "" {
		t.Errorf("unexpected string (-want, +got) = %v", diff)
	}
}

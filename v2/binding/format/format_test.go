/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package format_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Yangfisher1/ce-go-sdk/v2/binding/format"
	"github.com/Yangfisher1/ce-go-sdk/v2/event"
	"github.com/Yangfisher1/ce-go-sdk/v2/types"
)

func TestJSON(t *testing.T) {
	require := require.New(t)
	e := event.Event{
		Context: event.EventContextV03{
			Type:   "type",
			ID:     "id",
			Source: *types.ParseURIRef("source"),
		}.AsV03(),
	}
	e.SetExtension("ex", "val")
	require.NoError(e.SetData(event.ApplicationJSON, "foo"))
	b, err := format.JSON.Marshal(&e)
	require.NoError(err)
	assertJsonEquals(t, map[string]interface{}{
		"data":            "foo",
		"datacontenttype": "application/json",
		"ex":              "val",
		"id":              "id",
		"source":          "source",
		"specversion":     "0.3",
		"type":            "type",
	}, b)

	var e2 event.Event
	require.NoError(format.JSON.Unmarshal(b, &e2))
	require.Equal(e, e2)
}

func TestLookup(t *testing.T) {
	require := require.New(t)
	require.Nil(format.Lookup("nosuch"))

	{
		f := format.Lookup(event.ApplicationCloudEventsJSON)
		require.Equal(f.MediaType(), event.ApplicationCloudEventsJSON)
		require.Equal(format.JSON, f)
	}

	{
		f := format.Lookup("application/cloudevents+json; charset=utf-8")
		require.Equal(f.MediaType(), event.ApplicationCloudEventsJSON)
		require.Equal(format.JSON, f)
	}

	{
		f := format.Lookup("application/CLOUDEVENTS+json ; charset=utf-8")
		require.Equal(f.MediaType(), event.ApplicationCloudEventsJSON)
		require.Equal(format.JSON, f)
	}
}

func TestMarshalUnmarshal(t *testing.T) {
	require := require.New(t)
	e := event.Event{
		Context: event.EventContextV03{
			Type:   "type",
			ID:     "id",
			Source: *types.ParseURIRef("source"),
		}.AsV03(),
	}
	require.NoError(e.SetData(event.ApplicationJSON, "foo"))
	b, err := format.Marshal(format.JSON.MediaType(), &e)
	require.NoError(err)
	assertJsonEquals(t, map[string]interface{}{
		"data":            "foo",
		"datacontenttype": "application/json",
		"id":              "id",
		"source":          "source",
		"specversion":     "0.3",
		"type":            "type",
	}, b)

	var e2 event.Event
	require.NoError(format.Unmarshal(format.JSON.MediaType(), b, &e2))
	require.Equal(e, e2)

	_, err = format.Marshal("nosuchformat", &e)
	require.EqualError(err, "unknown event format media-type \"nosuchformat\"")
	err = format.Unmarshal("nosuchformat", nil, &e)
	require.EqualError(err, "unknown event format media-type \"nosuchformat\"")
}

type dummyFormat struct{}

func (dummyFormat) MediaType() string                    { return "dummy" }
func (dummyFormat) Marshal(*event.Event) ([]byte, error) { return []byte("dummy!"), nil }
func (dummyFormat) Unmarshal(b []byte, e *event.Event) error {
	e.DataEncoded = []byte("undummy!")
	return nil
}

func TestAdd(t *testing.T) {
	require := require.New(t)
	format.Add(dummyFormat{})
	require.Equal(dummyFormat{}, format.Lookup("dummy"))

	e := event.Event{}
	b, err := format.Marshal("dummy", &e)
	require.NoError(err)
	require.Equal("dummy!", string(b))
	err = format.Unmarshal("dummy", b, &e)
	require.NoError(err)
	require.Equal([]byte("undummy!"), e.Data())
}

func assertJsonEquals(t *testing.T, want map[string]interface{}, got []byte) {
	var gotToCompare map[string]interface{}
	require.NoError(t, json.Unmarshal(got, &gotToCompare))

	// Marshal and unmarshal want to make sure the types are correct
	wantBytes, err := json.Marshal(want)
	require.NoError(t, err)
	var wantToCompare map[string]interface{}
	require.NoError(t, json.Unmarshal(wantBytes, &wantToCompare))

	require.Equal(t, wantToCompare, gotToCompare)
}

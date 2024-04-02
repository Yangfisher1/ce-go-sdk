/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package client

import (
	obshttp "github.com/Yangfisher1/ce-go-sdk/observability/opencensus/v2/http"
	"github.com/Yangfisher1/ce-go-sdk/v2/client"
	"github.com/Yangfisher1/ce-go-sdk/v2/protocol/http"
)

func NewClientHTTP(topt []http.Option, copt []client.Option) (client.Client, error) {
	t, err := obshttp.NewObservedHTTP(topt...)
	if err != nil {
		return nil, err
	}

	copt = append(copt, client.WithTimeNow(), client.WithUUIDs(), client.WithObservabilityService(New()))

	c, err := client.New(t, copt...)
	if err != nil {
		return nil, err
	}

	return c, nil
}

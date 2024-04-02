/*
 Copyright 2021 The CloudEvents Authors
 SPDX-License-Identifier: Apache-2.0
*/

package context

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"go.uber.org/zap"
)

func TestLoggerContext(t *testing.T) {
	var namedLogger *zap.SugaredLogger
	if logger, err := zap.NewProduction(); err != nil {
		t.Fatal(err)
	} else {
		namedLogger = logger.Named("unittest").Sugar()
	}

	nopLogger := zap.NewNop().Sugar()

	testCases := map[string]struct {
		logger *zap.SugaredLogger
		ctx    context.Context
		want   *zap.SugaredLogger
	}{
		"todo context, set logger": {
			ctx:    context.TODO(),
			logger: namedLogger,
			want:   namedLogger,
		},
		"already set logger": {
			ctx:    WithLogger(context.TODO(), nopLogger),
			logger: namedLogger,
			want:   namedLogger,
		},
	}
	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {

			ctx := WithLogger(tc.ctx, tc.logger)
			got := LoggerFrom(ctx)

			if diff := cmp.Diff(tc.want, got, cmpopts.IgnoreUnexported(zap.SugaredLogger{})); diff != "" {
				t.Errorf("unexpected (-want, +got) = %v", diff)
			}
		})
	}
}

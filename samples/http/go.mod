module github.com/Yangfisher1/ce-go-sdk/samples/http

go 1.14

require (
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	github.com/Yangfisher1/ce-go-sdk/binding/format/protobuf/v2 v2.5.0
	github.com/Yangfisher1/ce-go-sdk/observability/opencensus/v2 v2.5.0
	github.com/Yangfisher1/ce-go-sdk/observability/opentelemetry/v2 v2.5.0
	github.com/Yangfisher1/ce-go-sdk/v2 v2.10.0
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.7.3
	github.com/kelseyhightower/envconfig v1.4.0
	go.opencensus.io v0.22.3
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.23.0
	go.opentelemetry.io/otel v1.0.0
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0
	go.opentelemetry.io/otel/sdk v1.0.0
	go.opentelemetry.io/otel/trace v1.0.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/Yangfisher1/ce-go-sdk/v2 => ../../v2

replace github.com/Yangfisher1/ce-go-sdk/binding/format/protobuf/v2 => ../../binding/format/protobuf/v2

replace github.com/Yangfisher1/ce-go-sdk/observability/opencensus/v2 => ../../observability/opencensus/v2

replace github.com/Yangfisher1/ce-go-sdk/observability/opentelemetry/v2 => ../../observability/opentelemetry/v2

module github.com/Yangfisher1/ce-go-sdk/test/benchmark

go 1.14

replace github.com/Yangfisher1/ce-go-sdk/v2 => ../../v2

replace github.com/Yangfisher1/ce-go-sdk/protocol/pubsub/v2 => ../../protocol/pubsub/v2

replace github.com/Yangfisher1/ce-go-sdk/protocol/amqp/v2 => ../../protocol/amqp/v2

replace github.com/Yangfisher1/ce-go-sdk/protocol/stan/v2 => ../../protocol/stan/v2

replace github.com/Yangfisher1/ce-go-sdk/protocol/nats/v2 => ../../protocol/nats/v2

replace github.com/Yangfisher1/ce-go-sdk/protocol/kafka_sarama/v2 => ../../protocol/kafka_sarama/v2

require (
	contrib.go.opencensus.io/exporter/prometheus v0.1.0
	github.com/Azure/go-amqp v0.17.0
	github.com/Shopify/sarama v1.25.0
	github.com/Yangfisher1/ce-go-sdk v1.2.0
	github.com/Yangfisher1/ce-go-sdk/protocol/amqp/v2 v2.0.0
	github.com/Yangfisher1/ce-go-sdk/protocol/kafka_sarama/v2 v2.0.0
	github.com/Yangfisher1/ce-go-sdk/protocol/nats/v2 v2.0.0
	github.com/Yangfisher1/ce-go-sdk/protocol/pubsub/v2 v2.0.0
	github.com/Yangfisher1/ce-go-sdk/protocol/stan/v2 v2.0.0
	github.com/Yangfisher1/ce-go-sdk/v2 v2.5.0
	github.com/google/go-cmp v0.5.6
	github.com/google/uuid v1.1.2
	github.com/gorilla/mux v1.7.3
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.7.0
	go.opencensus.io v0.23.0
)

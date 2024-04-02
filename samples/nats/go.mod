module github.com/Yangfisher1/ce-go-sdk/samples/nats

go 1.14

require (
	github.com/Yangfisher1/ce-go-sdk/protocol/nats/v2 v2.5.0
	github.com/Yangfisher1/ce-go-sdk/v2 v2.10.0
	github.com/google/uuid v1.1.1
	github.com/kelseyhightower/envconfig v1.4.0
)

replace github.com/Yangfisher1/ce-go-sdk/v2 => ../../v2

replace github.com/Yangfisher1/ce-go-sdk/protocol/nats/v2 => ../../protocol/nats/v2

module github.com/Yangfisher1/ce-go-sdk/samples/amqp

go 1.14

require (
	github.com/Azure/go-amqp v0.17.0
	github.com/Yangfisher1/ce-go-sdk/protocol/amqp/v2 v2.10.0
	github.com/Yangfisher1/ce-go-sdk/v2 v2.10.0
	github.com/google/uuid v1.1.1
)

replace github.com/Yangfisher1/ce-go-sdk/v2 => ../../v2

module github.com/Yangfisher1/ce-go-sdk/samples/ws

go 1.14

require (
	github.com/Yangfisher1/ce-go-sdk/protocol/ws/v2 v2.5.0
	github.com/Yangfisher1/ce-go-sdk/v2 v2.10.0
)

replace github.com/Yangfisher1/ce-go-sdk/v2 => ../../v2

replace github.com/Yangfisher1/ce-go-sdk/protocol/ws/v2 => ../../protocol/ws/v2

module github.com/cloudwego/gomall/app/notify

go 1.24.2

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

replace github.com/cloudwego/gomall/app/email/biz/consumer/email => ./biz/consumer/email

replace github.com/cloudwego/gomall/app/notify/infra => ./infra

replace github.com/cloudwego/gomall/app/notify/infra/mq => ./infra/mq

replace github.com/cloudwego/gomall/app/notify/infra/mq/email => ./infra/mq/email

require github.com/nats-io/nats.go v1.42.0

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/nats-io/nkeys v0.4.11 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	golang.org/x/crypto v0.37.0 // indirect
	golang.org/x/sys v0.32.0 // indirect
)

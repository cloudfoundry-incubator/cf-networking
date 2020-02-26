package mbus

import (
	"github.com/nats-io/nats.go"
)

type NatsConnWithUrlProvider struct {
	Url string
}

func (p *NatsConnWithUrlProvider) Connection(opts ...nats.Option) (NatsConn, error) {
	return nats.Connect(p.Url, opts...)
}

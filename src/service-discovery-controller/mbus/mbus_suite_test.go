package mbus_test

import (
	"crypto/tls"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"

	"github.com/nats-io/gnatsd/server"
	gnatsd "github.com/nats-io/gnatsd/test"
)

func TestMbus(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mbus Suite")
}

func RunServerOnPort(port int) *server.Server {
	opts := bootstrapOptions(port)
	return gnatsd.RunServer(&opts)
}

func RunServerWithTLSOnPort(port int, TLSConfig *tls.Config) *server.Server {
	opts := bootstrapOptions(port)
	opts.TLSConfig = TLSConfig
	return gnatsd.RunServer(&opts)
}

func bootstrapOptions(port int) server.Options {
	opts := gnatsd.DefaultTestOptions
	opts.Port = port
	opts.Username = "username"
	opts.Password = "password"
	return opts
}

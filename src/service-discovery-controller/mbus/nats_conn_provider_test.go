package mbus_test

import (
	"code.cloudfoundry.org/cf-networking-helpers/testsupport/ports"

	"crypto/tls"
	. "service-discovery-controller/mbus"
	"time"

	tls_helpers "code.cloudfoundry.org/cf-routing-test-helpers/tls"
	"code.cloudfoundry.org/tlsconfig"

	"github.com/nats-io/gnatsd/server"
	nats "github.com/nats-io/go-nats"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("NatsConnProvider", func() {
	var (
		provider    NatsConnProvider
		gnatsServer *server.Server
		natsCon     *nats.Conn
		port        int
	)

	BeforeEach(func() {
		port = ports.PickAPort()
		gnatsServer = RunServerOnPort(port)
		gnatsServer.Start()

		natsUrl := "nats://username:password@" + gnatsServer.Addr().String()

		provider = &NatsConnWithUrlProvider{
			Url: natsUrl,
		}
	})

	AfterEach(func() {
		if natsCon != nil {
			natsCon.Close()
		}
		gnatsServer.Shutdown()
	})

	It("returns a configured nats connection", func() {
		timeoutOption := nats.Timeout(42 * time.Second)
		conn, err := provider.Connection(timeoutOption)
		Expect(err).NotTo(HaveOccurred())
		var successfulCast bool
		natsCon, successfulCast = conn.(*nats.Conn)
		Expect(successfulCast).To(BeTrue())

		Expect(natsCon.Opts.Timeout).To(Equal(42 * time.Second))
	})

	Context("when provided with a TLS config", func() {
		var (
			natsCAPath             string
			mtlsNATSServerCertPath string
			mtlsNATSServerKeyPath  string
			mtlsNATSClientCert     tls.Certificate
			serverTLSConfig        *tls.Config
			clientTLSConfig        *tls.Config
			tlsPort                int
			err                    error
		)
		BeforeEach(func() {
			natsCAPath, mtlsNATSServerCertPath, mtlsNATSServerKeyPath, mtlsNATSClientCert = tls_helpers.GenerateCaAndMutualTlsCerts()
			serverTLSConfig, err = tlsconfig.Build(
				tlsconfig.WithInternalServiceDefaults(),
				tlsconfig.WithIdentityFromFile(mtlsNATSServerCertPath, mtlsNATSServerKeyPath),
			).Server(
				tlsconfig.WithClientAuthenticationFromFile(natsCAPath),
			)
			Expect(err).NotTo(HaveOccurred())

			tlsPort = ports.PickAPort()
			gnatsServer = RunServerWithTLSOnPort(tlsPort, serverTLSConfig)
			gnatsServer.Start()
		})

		It("returns a nats connection configured with TLS", func() {
			clientTLSConfig, err = tlsconfig.Build(
				tlsconfig.WithInternalServiceDefaults(),
				tlsconfig.WithIdentity(mtlsNATSClientCert),
			).Client(
				tlsconfig.WithAuthorityFromFile(natsCAPath),
			)
			Expect(err).NotTo(HaveOccurred())

			natsUrl := "nats://username:password@" + gnatsServer.Addr().String()

			provider = &NatsConnWithUrlProvider{
				Url: natsUrl,
			}

			timeoutOption := nats.Timeout(42 * time.Second)
			tlsOption := nats.Secure(clientTLSConfig)
			conn, err := provider.Connection(timeoutOption, tlsOption)
			Expect(err).NotTo(HaveOccurred())

			var successfulCast bool
			natsCon, successfulCast = conn.(*nats.Conn)
			Expect(successfulCast).To(BeTrue())

			Expect(natsCon.Opts.Timeout).To(Equal(42 * time.Second))
			Expect(natsCon.Opts.TLSConfig.Certificates).To(ContainElement(clientTLSConfig.Certificates[0]))
		})
	})
})

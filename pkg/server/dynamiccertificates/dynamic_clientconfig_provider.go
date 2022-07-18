package dynamiccertificates

import "crypto/tls"

type GetConfigForClientFunc func(clientHello *tls.ClientHelloInfo) (*tls.Config, error)

// DynamicClientConfigProvider dynamically loads certificates and provides a golang tls.
type DynamicClientConfigProvider interface {
	// DynamicClientConfigProvider wrap the original tls GetConfigForClient function
	// return a new and GetConfigForClient function
	// It should get base tls config and modify it for dynamic SNI
	WrapGetConfigForClient(getConfg GetConfigForClientFunc) GetConfigForClientFunc
}

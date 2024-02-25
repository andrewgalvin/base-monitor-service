package httpclient

import (
	"base-monitor-service/pkg/config"
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

// NewClient creates and returns a new HTTP client configured with the application's settings.
func NewClient(cfg *config.Config, proxy string) (*http.Client, error) {
	tlsConfig := newTLSConfig()

	proxyURL, err := newProxyURL(proxy)
	if err != nil {
		return nil, err // Properly handle the error
	}

	transport := &http.Transport{
		Proxy:           http.ProxyURL(proxyURL),
		TLSClientConfig: tlsConfig,
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   time.Duration(cfg.HTTPTimeout) * time.Second,
	}

	return client, nil
}

// newTLSConfig creates a TLS configuration for the HTTP client.
func newTLSConfig() *tls.Config {
	return &tls.Config{
		ClientAuth:               tls.RequireAndVerifyClientCert,
		MinVersion:               tls.VersionTLS12,
		CipherSuites:             preferredCipherSuites(),
		PreferServerCipherSuites: true,
	}
}

// newProxyURL parses the proxy URL from the configuration.
func newProxyURL(proxy string) (*url.URL, error) {
	if proxy == "" {
		return nil, nil // No proxy configured
	}
	return url.Parse(proxy)
}

// preferredCipherSuites returns a slice of uint16 values representing the preferred cipher suites for the HTTP client.
func preferredCipherSuites() []uint16 {
	return []uint16{
		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
		tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_RSA_WITH_AES_128_CBC_SHA,
	}
}

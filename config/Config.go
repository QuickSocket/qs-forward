package config

import (
	"flag"
	"fmt"
	"strings"
)

type Config struct {
	ClientId      string
	ClientSecret  string
	TLSSkipVerify bool
	WebSocketURL  string

	TargetURL string
}

func NewConfigFromCommandLine() (*Config, error) {
	clientId := flag.String("client-id", "", "The Client ID of the environment to receive callbacks for.")
	clientSecret := flag.String("client-secret", "", "Either of the two possible Client Secrets for the environment.")
	tlsSkipVerify := flag.Bool("tls-skip-verify", false, "When set to true, remote certificates will not be verified during TLS handshake.")
	websocketURL := flag.String("websocket-url", "wss://forward.quicksocket.io", "")
	flag.Parse()

	targetURL := strings.TrimSpace(flag.Arg(0))

	if clientId == nil || *clientId == "" || clientSecret == nil || *clientSecret == "" {
		return nil, fmt.Errorf("a client ID and client secret must be specified")
	}

	if targetURL == "" {
		return nil, fmt.Errorf("a targetURL must be specified")
	}

	return &Config{
		ClientId:      *clientId,
		ClientSecret:  *clientSecret,
		TLSSkipVerify: *tlsSkipVerify,
		WebSocketURL:  *websocketURL,

		TargetURL: targetURL,
	}, nil
}

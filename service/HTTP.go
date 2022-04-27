package service

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	"github.com/QuickSocket/qs-forward/model"
)

type HTTP struct {
	targetURL     string
	tlsSkipVerify bool
	callbackc     <-chan *model.Callback
}

func NewHTTP(targetURL string, tlsSkipVerify bool, callbackc <-chan *model.Callback) *HTTP {
	return &HTTP{
		targetURL:     targetURL,
		tlsSkipVerify: tlsSkipVerify,
		callbackc:     callbackc,
	}
}

func (s *HTTP) Start(logger *log.Logger) error {
	transport := &http.Transport{}
	if s.tlsSkipVerify {
		logger.Printf("WARNING: Not verifying TLS certificates during callback\n")

		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	client := &http.Client{
		Transport: transport,
	}

	for x := range s.callbackc {
		callbackModel := x
		go func() {
			if err := pushCallback(s.targetURL, client, callbackModel); err != nil {
				logger.Printf("%v\n", err)
			}
		}()
	}

	return nil
}

func pushCallback(targetURL string, client *http.Client, callbackModel *model.Callback) error {
	body := bytes.NewBufferString(callbackModel.SerializedData)
	req, err := http.NewRequest(http.MethodPost, targetURL, body)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request for callback: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("QS-Auth-Token-1", callbackModel.AuthToken1)
	req.Header.Add("QS-Auth-Token-2", callbackModel.AuthToken2)
	req.Header.Add("QS-Signature", callbackModel.Signature)

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request for callback: %w", err)
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return fmt.Errorf("warning: HTTP callback returned a status code of %v", res.StatusCode)
	}

	return nil
}

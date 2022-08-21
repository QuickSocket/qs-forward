package main

import (
	"fmt"
	"log"
	"os"

	"github.com/QuickSocket/qs-forward/config"
	"github.com/QuickSocket/qs-forward/model"
	"github.com/QuickSocket/qs-forward/service"
)

const CallbackChannelSize = 1 << 5

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	config, err := config.NewConfigFromCommandLine()
	if err != nil {
		return err
	}

	errc := make(chan error)
	callbackc := make(chan *model.Callback, CallbackChannelSize)
	logger := log.New(os.Stdout, "", log.LstdFlags)

	services := []Service{
		service.NewWebSocket(config.ClientId, config.ClientSecret, config.WebSocketURL, callbackc),
		service.NewHTTP(config.TargetURL, config.TLSSkipVerify, config.Quiet, callbackc),
	}

	for _, service := range services {
		go func(service Service) {
			errc <- service.Start(logger)
		}(service)
	}

	return <-errc
}

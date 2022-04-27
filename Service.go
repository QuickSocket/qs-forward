package main

import "log"

type Service interface {
	Start(logger *log.Logger) error
}

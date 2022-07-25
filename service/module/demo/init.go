package main

import (
	"sync"
)

var (
	mu   sync.Mutex
	demo *Demo
)

func init() {
	NewDemo()
}

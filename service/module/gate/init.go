package main

import (
	"sync"
)

var (
	mu      sync.Mutex
	gate    *Gate
	process *Process
)

func init() {
	NewProcess()
}

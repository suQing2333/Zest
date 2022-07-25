package main

import (
	"sync"
)

var (
	mu      sync.Mutex
	router  *Router
	process *Process
)

func init() {
	NewProcess()
	NewRouter()
}

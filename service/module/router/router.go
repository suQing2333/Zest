package main

import (
	"zest/engine/funcmgr"
	"zest/engine/zslog"
)

type Router struct {
}

func NewRouter() *Router {
	if router == nil {
		mu.Lock()
		defer mu.Unlock()
		if router == nil {
			router = &Router{}
			funcmgr.RegisterFunc(*router, router)
			funcmgr.RegisterSubcmdFunc(1001, "Router.ProtoTest")
		}
	}
	return router
}

func (r *Router) ProtoTest(data []byte) {
	zslog.LogDebug("Router ProtoTest  %v", data)
}

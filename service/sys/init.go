package sys

import (
	"sync"
)

var (
	mu      sync.Mutex
	RpcM    *RPCMgr
	Svr     *Service
	process *Process
)

func init() {
	GetRPCMgr()
	NewProcess()
}

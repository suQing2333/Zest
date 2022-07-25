package main

import (
	"zest/engine/funcmgr"
	"zest/engine/zslog"
)

type Demo struct {
}

func NewDemo() *Demo {
	if demo == nil {
		mu.Lock()
		defer mu.Unlock()
		if demo == nil {
			demo = &Demo{}
			funcmgr.RegisterFunc(*demo, demo)
			funcmgr.RegisterSubcmdFunc(1001, "Demo.ReceivedMessage")
			funcmgr.RegisterSubcmdFunc(1002, "Demo.RetrurnMessage")
		}
	}
	return demo
}

// 打印收到的消息  1001
func (d *Demo) ReceivedMessage(data []byte) {
	zslog.LogDebug("%v", data)
}

// 收到消息后返回消息 1002
func (d *Demo) RetrurnMessage(data []byte) map[string]interface{} {
	res := map[string]interface{}{
		"ClientCmd":       int32(100),
		"ClientSubcmd":    int32(1002),
		"ClientProtoData": []byte("demo return data"),
	}
	return res
}

// RPC
func (d *Demo) RPCAddNum(args ...interface{}) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum = sum + int(args[i].(float64))
	}
	zslog.LogDebug("%v", args)
	return sum
}

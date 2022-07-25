package main

import (
	"zest/engine/funcmgr"
	"zest/engine/itface"
	"zest/engine/netutil"
	"zest/engine/zslog"
	"zest/service/define"
	"zest/service/sys"
)

type Gate struct {
	// UserConnMap map[int64]int64
	ConnMgr itface.IConnMgr
}

func NewGate(connMgr itface.IConnMgr) *Gate {
	if gate == nil {
		mu.Lock()
		defer mu.Unlock()
		if gate == nil {
			gate = &Gate{
				ConnMgr: connMgr,
			}
			funcmgr.RegisterFunc(*gate, gate)
			funcmgr.RegisterSubcmdFunc(1001, "Gate.ProtoTest")
			funcmgr.RegisterSubcmdFunc(1002, "Gate.RPCTest")
		}
	}
	return gate
}

func (g *Gate) Start() {
}

func (g *Gate) ProtoTest(data []byte) {
	zslog.LogDebug("ProtoTest  %v", data)
}

func (g *Gate) RPCTest(data []byte) {
	res := sys.RPCCall("demo", "Demo.RPCAddNum", []interface{}{1, 2, 3, 4, 5})
	zslog.LogDebug("RPCReq get response,res :%v", res)
}

func (g *Gate) SendProtoToClientWithConnID(connID int64, protoData []byte) {
	msg := netutil.NewMsg(define.ServiceToClient, protoData)
	cond := map[string]interface{}{"connid": connID}
	conn, err := g.ConnMgr.GetConn(cond)
	if err != nil {
		zslog.LogInfo("SendMsgToClientWithConnID err :%v", err)
		return
	}
	err = conn.SendBuffMsg(msg)
	if err != nil {
		zslog.LogInfo("SendMsgToClientWithConnID err :%v", err)
		return
	}
}

func (g *Gate) SendMsgToClientWithConnID(connID int64, msg itface.IMessage) {
	cond := map[string]interface{}{"connid": connID}
	conn, err := g.ConnMgr.GetConn(cond)
	if err != nil {
		zslog.LogInfo("SendMsgToClientWithConnID err :%v", err)
		return
	}
	err = conn.SendBuffMsg(msg)
	if err != nil {
		zslog.LogInfo("SendMsgToClientWithConnID err :%v", err)
		return
	}
}

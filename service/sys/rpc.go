package sys

import (
	"encoding/json"
	"zest/engine/funcmgr"
	"zest/engine/pbmgr"
	"zest/engine/uuid"
	"zest/engine/zslog"
	"zest/service/define"
)

type RPCInfo struct {
	RPCID    int64
	Service  string
	CallFunc string
	CallArgs []interface{}
	Reply    interface{}
	done     chan interface{}
}

func RPCCall(service string, callFunc string, callArgs []interface{}) interface{} {
	rpci := &RPCInfo{
		RPCID:    uuid.GetNextVal(),
		Service:  service,
		CallFunc: callFunc,
		CallArgs: callArgs,
		done:     make(chan interface{}),
	}

	RpcM.Add(rpci)
	argsBytes, err := json.Marshal(rpci.CallArgs)
	if err != nil {
		zslog.LogError("RPCInfo Pakage err : ", err)
		return nil
	}
	protoData := pbmgr.GetDataByFullName("protoc.RPCRequest", define.ServiceMoudle(), int64(rpci.RPCID), rpci.CallFunc, argsBytes)
	funcmgr.CallFunc("Service.SendProtoToService", define.RPC, rpci.Service, int32(define.GetCmd(rpci.Service)), int32(define.RPCRequestCMD), protoData)

	select {
	case reply := <-rpci.done:
		rpci.Reply = reply
	}
	RpcM.Remove(rpci.RPCID)
	return rpci.Reply
}

type RPCMgr struct {
	RPCMap map[int64]*RPCInfo
}

func GetRPCMgr() *RPCMgr {
	if RpcM == nil {
		mu.Lock()
		defer mu.Unlock()
		if RpcM == nil {
			RpcM = &RPCMgr{
				RPCMap: make(map[int64]*RPCInfo),
			}
			funcmgr.RegisterFunc(*RpcM, RpcM)
			funcmgr.RegisterSubcmdFunc(define.RPCRequestCMD, "RPCMgr.ProcessRPCRequest")
			funcmgr.RegisterSubcmdFunc(define.RPCResponseCMD, "RPCMgr.ProcessRPCResponse")
		}
	}
	return RpcM
}

func (rpcm *RPCMgr) Add(rpci *RPCInfo) {
	rpcm.RPCMap[rpci.RPCID] = rpci
}

func (rpcm *RPCMgr) Remove(rpcID int64) {
	delete(rpcm.RPCMap, rpcID)
}

func (rpcm *RPCMgr) GetRPCInfo(rpcID int64) *RPCInfo {
	return rpcm.RPCMap[rpcID]
}

func (rpcm *RPCMgr) Write(rpcID int64, reply interface{}) {
	rpci := rpcm.GetRPCInfo(rpcID)
	if rpci == nil {
		zslog.LogWarn("not find RPCInfo rpcId : ", rpcID)
		return
	}
	rpci.done <- reply
}

func (rpcm *RPCMgr) ProcessRPCResponse(data []byte) {
	RPCResponse := pbmgr.NewPBMessage("protoc.RPCResponse", data)
	zslog.LogDebug("%v", data)
	rpcm.Write(RPCResponse.GetValue("RPCID").(int64), RPCResponse.GetValue("Reply").([]byte))
}

func (rpcm *RPCMgr) ProcessRPCRequest(data []byte) []byte {
	RPCRequest := pbmgr.NewPBMessage("protoc.RPCRequest", data)
	Args := []interface{}{}
	json.Unmarshal(RPCRequest.GetValue("CallArgs").([]byte), &Args)
	reply, err := funcmgr.CallFunc(RPCRequest.GetValue("CallFunc").(string), Args...)
	if err != nil {
		zslog.LogError("call func err:", err)
		return nil
	}
	if len(reply) != 1 {
		zslog.LogError("ProcessRPCRequest call func reply err :", reply)
		return nil
	}
	RPCReply, err := json.Marshal(reply[0])
	if err != nil {
		zslog.LogError("RPCInfo Pakage err : ", err)
		return nil
	}
	module := RPCRequest.GetValue("OriService").(string)
	protoData := pbmgr.GetDataByFullName("protoc.RPCResponse", RPCRequest.GetValue("RPCID").(int64), RPCReply)
	Cmd := define.GetCmd(module)
	res := pbmgr.GetDataByFullName("protoc.BaseProto", int32(Cmd), int32(define.RPCResponseCMD), protoData, []byte{})
	return res
}

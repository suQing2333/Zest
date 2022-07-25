package sys

import (
	"encoding/json"
	"reflect"
	"zest/engine/funcmgr"
	"zest/engine/itface"
	"zest/engine/pbmgr"
	"zest/engine/zslog"
	"zest/service/define"
)

func NewProcess() *Process {
	if process == nil {
		mu.Lock()
		defer mu.Unlock()
		if process == nil {
			process = &Process{}
			funcmgr.RegisterFunc(*process, process)
		}
	}
	return process
}

// 默认的处理请求,只适用于非gate,router
type Process struct{}

func (p *Process) ProcessRPC(conn itface.IConnection, data []byte) {
	BaseProto := pbmgr.NewPBMessage("protoc.BaseProto", data)
	Cmd := BaseProto.GetValue("Cmd").(int32)
	subCmd := BaseProto.GetValue("Subcmd").(int32)
	module := define.GetModule(Cmd)

	zslog.LogDebug("ProcessRPC %v", module)

	switch module {
	case define.ServiceMoudle():
		res, err := funcmgr.CallFuncWithSubcmd(subCmd, BaseProto.GetValue("Data").([]byte))
		if err != nil {
			zslog.LogError("ProcessRPC err :", err)
			return
		}
		if len(res) != 1 {
			return
		}
		resData := res[0].([]byte)
		funcmgr.CallFunc("Service.SendBaseProtoWithConn", conn, define.RPC, resData)
	default:
		if subCmd == int32(define.RPCRequestCMD) {
			RPCRequest := pbmgr.NewPBMessage("protoc.RPCRequest", BaseProto.GetValue("Data").([]byte))
			callFunc := RPCRequest.GetValue("CallFunc").(string)
			callArgs := []interface{}{}
			json.Unmarshal(RPCRequest.GetValue("CallArgs").([]byte), &callArgs)
			RPCReply := RPCCall(module, callFunc, callArgs)
			oriModule := RPCRequest.GetValue("OriService").(string)
			protoData := pbmgr.GetDataByFullName("protoc.RPCResponse", RPCRequest.GetValue("RPCID").(int64), RPCReply)
			Cmd := define.GetCmd(oriModule)
			funcmgr.CallFunc("Service.SendProtoWithConn", conn, define.RPC, Cmd, define.RPCResponseCMD, protoData)
		}
	}
}

func (p *Process) ProcessClientToService(conn itface.IConnection, data []byte) {
	BaseProto := pbmgr.NewPBMessage("protoc.BaseProto", data)
	Cmd := BaseProto.GetValue("Cmd").(int32)
	subCmd := BaseProto.GetValue("Subcmd").(int32)
	module := define.GetModule(Cmd)
	switch module {
	case define.ServiceMoudle():
		res, err := funcmgr.CallFuncWithSubcmd(subCmd, BaseProto.GetValue("Data").([]byte))
		if err != nil {
			zslog.LogError("ProcessRPC err :", err)
			return
		}
		if len(res) != 1 {
			return
		}
		// res 结构 map[string]interface{}{"ClientCmd":ClientCmd,"ClientSubcmd":ClientSubcmd,"ClientProtoData":ClientProtoData}
		// 满足该结构则将消息发送到router
		if _, ok := res[0].(map[string]interface{}); !ok {
			return
		}
		resMap := res[0].(map[string]interface{})
		if _, ok := resMap["ClientCmd"]; !ok {
			return
		}
		if _, ok := resMap["ClientSubcmd"]; !ok {
			return
		}
		if _, ok := resMap["ClientProtoData"]; !ok {
			return
		}
		if reflect.TypeOf(resMap["ClientCmd"]).Kind() != reflect.Int32 || reflect.TypeOf(resMap["ClientSubcmd"]).Kind() != reflect.Int32 {
			return
		}
		backInfo := BaseProto.GetValue("BackInfo").([]byte)
		protoData := pbmgr.GetDataByFullName("protoc.BaseProto", resMap["ClientCmd"].(int32), resMap["ClientSubcmd"].(int32), resMap["ClientProtoData"].([]byte), backInfo)
		funcmgr.CallFunc("Service.SendBaseProtoToService", define.ServiceToClient, "router", protoData)
	}
}

func (p *Process) ProcessServiceToClient(conn itface.IConnection, data []byte) {
	// 非gate,router不应该收到该类型消息
	return
}

func processRequest(req itface.IRequest) {
	conn := req.GetConnection()
	data := req.GetData()
	switch req.GetMsgType() {
	case define.RPC:
		zslog.LogDebug("RPC")
		funcmgr.CallFunc("Process.ProcessRPC", conn, data)
	case define.ClientToService:
		zslog.LogDebug("CS")
		funcmgr.CallFunc("Process.ProcessClientToService", conn, data)
	case define.ServiceToClient:
		zslog.LogDebug("SC")
		funcmgr.CallFunc("Process.ProcessServiceToClient", conn, data)
	default:
		return
	}
}

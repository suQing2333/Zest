package main

import (
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

type Process struct{}

func (p *Process) ProcessClientToService(conn itface.IConnection, data []byte) {
	CSProto := pbmgr.NewPBMessage("protoc.CSProto", data)
	Cmd := CSProto.GetValue("Cmd").(int32)
	subCmd := CSProto.GetValue("Subcmd").(int32)

	switch define.GetModule(Cmd) {
	case define.ServiceMoudle():
		res, err := funcmgr.CallFuncWithSubcmd(subCmd, CSProto.GetValue("Data").([]byte))
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
		protoData := pbmgr.GetDataByFullName("protoc.CSProto", resMap["ClientCmd"].(int32), resMap["ClientSubcmd"].(int32), resMap["ClientProtoData"].([]byte))
		funcmgr.CallFunc("Gate.SendProtoToClientWithConnID", conn.GetConnID(), protoData)
	default:
		backData := pbmgr.GetDataByFullName("protoc.BackInfo", define.ServiceMoudle(), define.ServiceID(), conn.GetConnID())
		zslog.LogDebug("backData %v", backData)
		BaseData := pbmgr.GetDataByFullName("protoc.BaseProto", Cmd, subCmd, CSProto.GetValue("Data").([]byte), backData)
		funcmgr.CallFunc("Service.SendBaseProtoToService", define.ClientToService, "router", BaseData)
	}
}

func (p *Process) ProcessServiceToClient(conn itface.IConnection, data []byte) {
	BaseProto := pbmgr.NewPBMessage("protoc.BaseProto", data)
	Cmd := BaseProto.GetValue("Cmd").(int32)
	subCmd := BaseProto.GetValue("Subcmd").(int32)
	BackInfo := pbmgr.NewPBMessage("protoc.BackInfo", BaseProto.GetValue("BackInfo").([]byte))

	module := BackInfo.GetValue("Module").(string)
	if module != define.ServiceMoudle() {
		zslog.LogError("ProcessServiceToClient err,module :%v", module)
		return
	}
	sid := BackInfo.GetValue("Sid").(int32)
	if sid != define.ServiceID() {
		zslog.LogError("ProcessServiceToClient err,sid :%v", sid)
		return
	}
	connID := BackInfo.GetValue("ConnID").(int64)
	CSProto := pbmgr.GetDataByFullName("protoc.CSProto", Cmd, subCmd, BaseProto.GetValue("Data").([]byte))
	funcmgr.CallFunc("Gate.SendProtoToClientWithConnID", connID, CSProto)
}

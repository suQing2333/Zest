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
	BaseProto := pbmgr.NewPBMessage("protoc.BaseProto", data)
	Cmd := BaseProto.GetValue("Cmd").(int32)
	subCmd := BaseProto.GetValue("Subcmd").(int32)

	switch define.GetModule(Cmd) {
	case define.ServiceMoudle():
		zslog.LogDebug("subCmd %v,Data %v", subCmd, BaseProto.GetValue("Data").([]byte))
		res, err := funcmgr.CallFuncWithSubcmd(subCmd, BaseProto.GetValue("Data").([]byte))
		if err != nil {
			zslog.LogError("ProcessClientToService err :", err)
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
		BackInfoMsg := pbmgr.NewPBMessage("process.BackInfo", backInfo)
		module := BackInfoMsg.GetValue("Module").(string)
		sid := BackInfoMsg.GetValue("Sid").(int32)
		protoData := pbmgr.GetDataByFullName("protoc.BaseProto", resMap["ClientCmd"].(int32), resMap["ClientSubcmd"].(int32), resMap["ClientProtoData"].([]byte), backInfo)
		funcmgr.CallFunc("Service.SendBaseProtoToServiceWithSID", define.ServiceToClient, module, sid, protoData)
	default:
		funcmgr.CallFunc("Service.SendBaseProtoToService", define.ClientToService, define.GetModule(Cmd), data)
	}
}

func (p *Process) ProcessServiceToClient(conn itface.IConnection, data []byte) {
	BaseProto := pbmgr.NewPBMessage("protoc.BaseProto", data)
	backInfo := BaseProto.GetValue("BackInfo").([]byte)
	BackInfoMsg := pbmgr.NewPBMessage("protoc.BackInfo", backInfo)
	module := BackInfoMsg.GetValue("Module").(string)
	sid := BackInfoMsg.GetValue("Sid").(int32)
	funcmgr.CallFunc("Service.SendBaseProtoToServiceWithSID", define.ServiceToClient, module, sid, data)
}

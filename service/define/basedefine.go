package define

import (
	"zest/engine/conf"
)

const (
	RPC = iota
	ClientToService
	ServiceToClient
	ServiceTranspond
)

var ServiceCmdMap = map[string]int32{
	"gate":   1,
	"router": 2,
	"demo":   3,
}

var ServiceModuleMap = make(map[int32]string)

func init() {
	BuildModuleMap()
}

func BuildModuleMap() {
	for key, value := range ServiceCmdMap {
		ServiceModuleMap[value] = key
	}
}

func GetCmd(module string) int32 {
	return ServiceCmdMap[module]
}

func GetModule(Cmd int32) string {
	return ServiceModuleMap[Cmd]
}

func ServiceMoudle() string {
	return conf.GetGlobalConf("Service").(string)
}
func ServiceID() int32 {
	return conf.GetGlobalConf("Sid").(int32)
}

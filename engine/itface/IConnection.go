package itface

import (
	"net"
)

// 连接接口定义
type IConnection interface {
	Start() //启动连接
	Stop()  //停止连接

	GetConnID() int64 //获取连接ID
	GetTCPConnection() *net.TCPConn
	SetExtraInfo(extraInfo map[string]interface{})
	GetExtraInfo() map[string]interface{}

	SendMsg(msg IMessage) error // 发送消息到客户端
	SendBuffMsg(msg IMessage) error
}

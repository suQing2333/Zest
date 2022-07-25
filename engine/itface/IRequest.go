package itface

type IRequest interface {
	GetConnection() IConnection //获取请求连接信息
	GetData() []byte            //获取请求消息的数据
	GetMsgType() uint32         //获取请求消息类型
	GetMsg() IMessage           //获取请求消息
}

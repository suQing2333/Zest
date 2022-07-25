package itface

type IMessage interface {
	GetDataLen() uint32
	GetMsgType() uint32
	GetMsgData() []byte
}

package netutil

type Message struct {
	DataLen uint32
	MsgType uint32
	MsgData []byte
}

func NewMsg(msgType uint32, data []byte) *Message {
	return &Message{
		DataLen: uint32(len(data)),
		MsgType: msgType,
		MsgData: data,
	}
}

//GetDataLen 获取消息数据段长度
func (msg *Message) GetDataLen() uint32 {
	return msg.DataLen
}

//GetMsgID 获取消息ID
func (msg *Message) GetMsgType() uint32 {
	return msg.MsgType
}

//GetData 获取消息内容
func (msg *Message) GetMsgData() []byte {
	return msg.MsgData
}

//SetDataLen 设置消息数据段长度
func (msg *Message) SetDataLen(dataLen uint32) {
	msg.DataLen = dataLen
}

//SetMsgID 设计消息ID
func (msg *Message) SetMsgID(msgType uint32) {
	msg.MsgType = msgType
}

//SetData 设计消息内容
func (msg *Message) SetData(data []byte) {
	msg.MsgData = data
}

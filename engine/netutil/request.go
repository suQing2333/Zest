package netutil

import (
	"zest/engine/itface"
)

type Request struct {
	conn itface.IConnection
	msg  itface.IMessage
}

func NewRequest(conn itface.IConnection, msg *Message) *Request {
	req := &Request{
		conn: conn,
		msg:  msg,
	}
	return req
}

func (r *Request) GetMsg() itface.IMessage {
	return r.msg
}

func (r *Request) GetConnection() itface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetMsgData()
}

func (r *Request) GetMsgType() uint32 {
	return r.msg.GetMsgType()
}

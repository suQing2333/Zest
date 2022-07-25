package netutil

import (
	"bytes"
	"encoding/binary"
	"zest/engine/itface"
	// "fmt"
)

func Pack(msg itface.IMessage) (out []byte, err error) {
	outbuff := bytes.NewBuffer([]byte{})
	if err = binary.Write(outbuff, binary.LittleEndian, msg.GetDataLen()); err != nil {
		return
	}
	if err := binary.Write(outbuff, binary.LittleEndian, msg.GetMsgType()); err != nil {
		return nil, err
	}
	if err = binary.Write(outbuff, binary.LittleEndian, msg.GetMsgData()); err != nil {
		return
	}
	out = outbuff.Bytes()
	return out, nil
}

func Unpack(headdata []byte) (head *Message, err error) {
	headBuf := bytes.NewReader(headdata)
	head = &Message{}
	// 读取Len
	if err = binary.Read(headBuf, binary.LittleEndian, &head.DataLen); err != nil {
		return nil, err
	}
	//读取Type
	if err := binary.Read(headBuf, binary.LittleEndian, &head.MsgType); err != nil {
		return nil, err
	}
	// head.Data = make([]byte, head.Len)

	// if err = binary.Read(headBuf, binary.LittleEndian, &head.Data); err != nil {
	// 	return nil, err
	// }

	return head, nil
}

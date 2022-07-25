package pbmgr

import (
	"fmt"
	"testing"
	// "zest/engine/zslog"
	// "encoding/json"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoregistry"
	"zest/service/protoc"
)

func TestProto(t *testing.T) {

	data := []byte{}
	BaseBegin := &stProto.BaseProto{
		Cmd:    1,
		Subcmd: 2,
		Data:   data,
	}

	//protobuf编码
	pData, err := proto.Marshal(BaseBegin)
	if err != nil {
		panic(err)
	}
	fmt.Println(pData)

	BaseAfter := &stProto.BaseProto{}
	proto.Unmarshal(pData, BaseAfter)
	fmt.Printf("%T\n", BaseAfter.Cmd)

	fmt.Println(protoregistry.GlobalTypes)
	// any, err := json.Marshal([]interface{}{1, "234", true, 4, 5})

	// fmt.Println(any)
	// var args []interface{}
	// err = json.Unmarshal(any, &args)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(args)
	// RPCInfo := &stProto.RPCInfo{
	// 	Service:  "demo",
	// 	RPCID:    10000,
	// 	CallFunc: "Service.TestCall",
	// 	CallArgs: any,
	// }
	// pData, err = proto.Marshal(RPCInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(pData)
	msg := GetMessageByFullName("protoc.CSProto", pData)
	fmt.Printf("%T  %v\n", msg, msg)
	fmt.Println(GetMessageFieldsValue(msg, "Cmd"))
	SetMessageFieldsValue(msg, "Cmd", int32(2))
	fmt.Println(GetMessageFieldsValue(msg, "Cmd"))
	// GetDataByFullName("protoc.BaseProto", int32(1), int32(2), []byte{})

}

package pbmgr

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"zest/engine/zslog"
	_ "zest/service/protoc"
)

// 在proto.Message外封装一层

type PBMessage struct {
	pm        proto.Message
	isSuccess bool // 创建proto.Message是否成功
}

func NewPBMessage(fullName string, data []byte) *PBMessage {
	pm := GetMessageByFullName(fullName, data)
	isSuccess := true
	if pm == nil {
		isSuccess = false
	}
	pbm := &PBMessage{
		pm:        pm,
		isSuccess: isSuccess,
	}
	return pbm
}

func (pbm *PBMessage) GetValue(fieldName string) interface{} {
	if !pbm.isSuccess {
		return nil
	}
	return GetMessageFieldsValue(pbm.pm, fieldName)

}

func (pbm *PBMessage) SetValue(fieldName string, value interface{}) {
	if !pbm.isSuccess {
		return
	}
	SetMessageFieldsValue(pbm.pm, fieldName, value)
}

// 通过fullname解析data 返回message
// 使用该方法生成的message虽然类型相同,但是已经不能通过 .*** 或者Get***() 等方式去获取fields的value
func GetMessageByFullName(fullName string, data []byte) proto.Message {
	msgName := protoreflect.FullName(fullName)
	msgType, err := protoregistry.GlobalTypes.FindMessageByName(msgName)
	if err != nil {
		zslog.LogWarn("proto name %v not found ", fullName)
		return nil
	}
	message := msgType.New().Interface()
	err = proto.Unmarshal(data, message)
	if err != nil {
		return nil
	}
	return message
}

// 通过fullname args 生成 []byte 注意args需要与proto中的field类型相同
func GetDataByFullName(fullName string, args ...interface{}) []byte {
	msgName := protoreflect.FullName(fullName)
	msgType, err := protoregistry.GlobalTypes.FindMessageByName(msgName)
	if err != nil {
		zslog.LogWarn("proto name %v not found ", fullName)
		return nil
	}
	message := msgType.New()
	msgFields := message.Descriptor().Fields()
	msgArgsLen := msgFields.Len()
	for i := 0; i < msgArgsLen; i++ {
		message.Set(msgFields.Get(i), protoreflect.ValueOf(args[i]))
	}
	data, err := proto.Marshal(message.Interface())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return data
}

// 获取proto.Message中fieldName字段的value
func GetMessageFieldsValue(pm proto.Message, fieldName string) interface{} {
	message := pm.ProtoReflect()
	msgFields := message.Descriptor().Fields()
	msgField := msgFields.ByTextName(fieldName)
	// 获取的字段不存在
	if msgField == nil {
		return nil
	}
	value := message.Get(msgField)
	return value.Interface()
}

// 设置proto.Message fieldName 字段 ,注意value的类型
func SetMessageFieldsValue(pm proto.Message, fieldName string, value interface{}) {
	message := pm.ProtoReflect()
	msgFields := message.Descriptor().Fields()
	msgField := msgFields.ByTextName(fieldName)
	if msgField == nil {
		return
	}
	message.Set(msgField, protoreflect.ValueOf(value))
}

// 获取message是否有field字段
func MessageHasField(pm proto.Message, fieldName string) bool {
	message := pm.ProtoReflect()
	msgFields := message.Descriptor().Fields()
	msgField := msgFields.ByTextName(fieldName)
	if msgField == nil {
		return false
	}
	return true
}

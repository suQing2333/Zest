# zest Server

Zest 是一个基于Golang的服务器框架

## 简介
本项目主要是个人学习项目,加入了部分个人理解
1. 函数管理器
  funcmgr,个人喜欢能够通过字符串就可以调用对应函数,但是golang原生的函数变量限制比较大,在定义变量的时候就得定义好参数类型和返回值类型.
  注册函数如下:
  ```go
  func RegisterFunc(singleton interface{}, singletionPointer interface{}) {
  	typeName := reflect.TypeOf(singleton).Name()
  	if _, ok := funcMap[typeName]; !ok {
  		funcMap[typeName] = make(map[string]reflect.Value)
  	}
  	vf := reflect.ValueOf(singletionPointer)
  	vft := vf.Type()
  	mNum := vf.NumMethod()
  	for i := 0; i < mNum; i++ {
  		methodName := vft.Method(i).Name
  		funcMap[typeName][methodName] = vf.Method(i)
  	}
  }
  ```
  首先必须保证注册函数为结构体函数,且必须为单例模式,因为注册函数只能定位到注册对象下的所有函数,如果有两个同名结构体,可以通过注册的先后顺序以达到后注册覆盖先注册的效果,一定程度上能实现对函数的重写
  调用函数只需要以"struct.func"的形式传入结构体名与函数名,和函数所需参数即可调用
  调用函数如下:
  ```go
  func CallFunc(callFunc string, args ...interface{}) ([]interface{}, error) {
	parseFunc := strings.Split(callFunc, ".")
	if len(parseFunc) != 2 {
		err := fmt.Errorf("parse callFunc arg num err")
		return nil, err
	}
	obj := parseFunc[0]
	method := parseFunc[1]
	if _, ok := funcMap[obj]; !ok {
		err := fmt.Errorf("CallFunc err, not obj : %v\n", obj)
		return nil, err
	}
	if _, ok := funcMap[obj][method]; !ok {
		err := fmt.Errorf("CallFunc err, not method : %v\n", method)
		return nil, err
	}
	parms := []reflect.Value{}

	for i := 0; i < len(args); i++ {
		parms = append(parms, reflect.ValueOf(args[i]))
	}
	resValue := funcMap[obj][method].Call(parms)
	res := make([]interface{}, len(resValue))
	for index, value := range resValue {
		res[index] = value.Interface()
	}
	return res, nil
}
  ```
  调用函数最终会返回函数所有返回值的interface{}类型
2. protobuf管理器
  pbmgr,同样的,个人也比较希望能够通过message名与序列化后的[]byte生成对应的message对象,或者通过message名与参数生成序列化后的[]byte.
  生成message函数如下:
  ```go
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
  ```
  但是该方法生成的对象用检测对象类型的确是对应fullName的类型,但是只要使用 .*** 或者Get***()等方法就会返回类型错误的panic(推测是类型的确是proto.Message,只不过具备了对应结构体的属性,所以得按照proto.Message的方法来使用返回的对象,没想到太好的办法)
  生成[]byte函数如下:
  ```go
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
  ```
  由于message在定义成员变量的时候是要求要输入字段的序列号(不知道具体学名是啥)的,所以调用该函数也应该按照字段的升序序列来写入要赋值的变量
  对于proto.Message也增加了两个函数以处理字段的获取与赋值
  函数如下:
  ```go
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
  ```
  为了更方便的调用实现了一个PBMessage的结构体
  ```go
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
  ```
  可以直接通过NewPBMessage生成对象,GetValue与SetValue来获取和设置字段

todo

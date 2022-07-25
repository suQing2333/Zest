package funcmgr

import (
	"fmt"
	"reflect"
	"strings"
)

var funcMap = make(map[string]map[string]reflect.Value)
var subcmdMap = make(map[int32]string)

// 注册函数
// 可以通过注册顺序覆盖注册函数
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

// 调用函数
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

// 注册subcmd必须是funcMap中已注册的函数
func RegisterSubcmdFunc(subcmd int, callFunc string) error {
	parseFunc := strings.Split(callFunc, ".")
	if len(parseFunc) != 2 {
		err := fmt.Errorf("parse callFunc arg num err")
		return err
	}
	obj := parseFunc[0]
	method := parseFunc[1]
	if _, ok := funcMap[obj]; !ok {
		err := fmt.Errorf("RegisterSubcmdFunc err, not obj : %v\n", obj)
		return err
	}
	if _, ok := funcMap[obj][method]; !ok {
		err := fmt.Errorf("CallFunc err, not method : %v\n", method)
		return err
	}
	subcmdMap[int32(subcmd)] = callFunc
	return nil
}

// 通过subcmd调用函数
func CallFuncWithSubcmd(subcmd int32, args ...interface{}) ([]interface{}, error) {
	if _, ok := subcmdMap[subcmd]; !ok {
		err := fmt.Errorf("not found match callfunc,subcmd : %v", subcmd)
		return nil, err
	}
	callFunc := subcmdMap[subcmd]
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
	fmt.Println(subcmdMap)
	resValue := funcMap[obj][method].Call(parms)
	res := make([]interface{}, len(resValue))
	for index, value := range resValue {
		res[index] = value.Interface()
	}
	return res, nil
}

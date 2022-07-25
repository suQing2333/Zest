package timer

import (
	"reflect"
	"zest/engine/common"
	"zest/engine/zslog"
)

var GtID = 1

//Timer 定时器实现
type Timer struct {
	//延迟调用函数
	tID       int
	delayFunc func(...interface{})
	funcArgs  []interface{}
	funcName  string
	delayTime int64
	//本次调用时间
	callTime int64
	//创建时间
	creatTime int64
	//是否循环调用
	loop bool
}

//创建一个定时器
func NewTimer(f func(v ...interface{}), args []interface{}, delayTime int64, loop bool) *Timer {
	nowTime := common.UnixMilli()
	name := reflect.TypeOf(f).Name()
	t := &Timer{
		tID:       GtID,
		delayFunc: f,
		funcArgs:  args,
		funcName:  name,
		delayTime: delayTime,
		callTime:  nowTime + delayTime,
		creatTime: nowTime,
		loop:      loop,
	}
	GtID++
	tc.AddTimer(t)
	return t
}

func (t *Timer) Call() {
	defer func() {
		if err := recover(); err != nil {
			zslog.LogError("call func err : %v", err)
		}
	}()

	//调用定时器超时函数
	t.delayFunc(t.funcArgs...)
	if t.loop {
		nowTime := common.UnixMilli()
		t.callTime = nowTime + t.delayTime
	}
}

func (t *Timer) Stop() {
	tc.RemoveTimer(t)
}

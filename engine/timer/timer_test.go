package timer

import (
	"fmt"
	"testing"
	"time"
	"zest/engine/common"
)

func myFunc(v ...interface{}) {
	nowTime := common.UnixMilli()
	fmt.Printf("No.%d function calld. delay %d second(s) nowTime %d \n", v[0].(int), v[1].(int), nowTime)
}

func TestTimer(t *testing.T) {
	nowTime := common.UnixMilli()
	fmt.Printf("test start time : %d \n", nowTime)
	tc := NewTimeClock()

	t1 := NewTimer(myFunc, []interface{}{2, 5}, 5000, false)
	_ = tc.AddTimer(t1)

	t2 := NewTimer(myFunc, []interface{}{3, 10}, 10000, true)
	_ = tc.AddTimer(t2)

	time.Sleep(10 * time.Minute)
}

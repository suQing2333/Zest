package timer

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"zest/engine/common"
	"zest/engine/zslog"
)

const (
	MinScale = 50
)

var tc *TimeClock

type TimeClock struct {
	timeScale  int64                    // 时间刻度间隔
	timerQueue map[int64]map[int]*Timer // 定时任务队列
	sync.RWMutex
}

func init() {
	tc = NewTimeClock()
	go tc.run()
}

func NewTimeClock() *TimeClock {
	tc := &TimeClock{
		timeScale:  MinScale,
		timerQueue: make(map[int64]map[int]*Timer),
	}
	go tc.run()
	return tc
}

func (tc *TimeClock) AddTimer(t *Timer) error {
	defer func() error {
		if err := recover(); err != nil {
			errstr := fmt.Sprintf("addTimer function err : %s", err)
			zslog.LogError(errstr)
			return errors.New(errstr)
		}
		return nil
	}()
	now := common.UnixMilli()
	if t.callTime > now {
		var scaleIndex int64
		if t.callTime%tc.timeScale == 0 {
			scaleIndex = t.callTime / tc.timeScale
		} else {
			scaleIndex = t.callTime/tc.timeScale + 1
		}
		if _, ok := tc.timerQueue[scaleIndex]; !ok {
			tc.timerQueue[scaleIndex] = make(map[int]*Timer)
		}
		tc.timerQueue[scaleIndex][t.tID] = t
	} else {
		return nil
	}
	return nil
}

func (tc *TimeClock) RemoveTimer(t *Timer) {
	tc.Lock()
	defer tc.Unlock()
	var scaleIndex int64
	if t.callTime%tc.timeScale == 0 {
		scaleIndex = t.callTime / tc.timeScale
	} else {
		scaleIndex = t.callTime/tc.timeScale + 1
	}

	if _, ok := tc.timerQueue[scaleIndex][t.tID]; ok {
		delete(tc.timerQueue[scaleIndex], t.tID)
	}
}

func (tc *TimeClock) run() {
	for {
		time.Sleep(time.Millisecond)
		tc.Lock()
		now := common.UnixMilli()
		curScale := now / tc.timeScale
		// 属于该时间刻度上的函数都要被执行
		if curTimers, ok := tc.timerQueue[curScale]; ok {
			for _, t := range curTimers {
				t.Call()
				tc.AddTimer(t)
			}
			delete(tc.timerQueue, curScale)
		}
		tc.Unlock()
	}
}

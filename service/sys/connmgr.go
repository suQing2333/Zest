package sys

import (
	"fmt"
	"math/rand"
	"sync"
	"zest/engine/itface"
	"zest/engine/zslog"
)

type ConnMgr struct {
	conns      map[string]map[int]itface.IConnection
	count      int
	msgHandler func(req itface.IRequest)
	connLock   sync.RWMutex
}

func NewConnMgr() *ConnMgr {
	cm := &ConnMgr{
		conns: make(map[string]map[int]itface.IConnection),
		count: 0,
	}
	return cm
}
func (cm *ConnMgr) SetMsgHandler(funcHook func(itface.IRequest)) {
	cm.msgHandler = funcHook
}

func (cm *ConnMgr) GetMsgHandler() func(itface.IRequest) {
	return cm.msgHandler
}

func (cm *ConnMgr) GetConn(selectInfo map[string]interface{}) (itface.IConnection, error) {
	if _, ok := selectInfo["sname"]; !ok {
		err := fmt.Errorf("can not selecet conn,not find service name")
		return nil, err
	}
	if _, ok := selectInfo["sid"]; !ok {
		err := fmt.Errorf("can not selecet conn,not find sid")
		return nil, err
	}
	sName := selectInfo["sname"].(string)
	sid := selectInfo["sid"].(int)
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	if _, ok := cm.conns[sName]; !ok {
		err := fmt.Errorf("connmgr can not add conn,not find service name,sname :%v", sName)
		return nil, err
	}
	return cm.conns[sName][sid], nil
}

func (cm *ConnMgr) Add(conn itface.IConnection) error {
	extraInfo := conn.GetExtraInfo()
	if extraInfo == nil {
		err := fmt.Errorf("can not add conn,extraInfo is nil")
		return err
	}
	if _, ok := extraInfo["sname"]; !ok {
		err := fmt.Errorf("can not add conn,not find service name")
		return err
	}
	if _, ok := extraInfo["sid"]; !ok {
		err := fmt.Errorf("can not add conn,not find sid")
		return err
	}
	sName := extraInfo["sname"].(string)
	sid := extraInfo["sid"].(int)
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	if _, ok := cm.conns[sName]; !ok {
		cm.conns[sName] = make(map[int]itface.IConnection)
	}
	cm.conns[sName][sid] = conn
	cm.count++
	zslog.LogDebug("%v", cm.conns)
	return nil

}

func (cm *ConnMgr) Remove(conn itface.IConnection) error {
	if conn == nil {
		err := fmt.Errorf("can not remove conn,conn is nil")
		return err
	}
	extraInfo := conn.GetExtraInfo()
	if extraInfo == nil {
		err := fmt.Errorf("can not remove conn,extraInfo is nil")
		return err
	}
	if _, ok := extraInfo["sname"]; !ok {
		err := fmt.Errorf("can not remove conn,not find service name")
		return err
	}
	if _, ok := extraInfo["sid"]; !ok {
		err := fmt.Errorf("can not remove conn,not find sid")
		return err
	}
	sName := extraInfo["sname"].(string)
	sid := extraInfo["sid"].(int)
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	if _, ok := cm.conns[sName]; !ok {
		err := fmt.Errorf("connmgr can not remove conn,not find service name,sname :%v", sName)
		return err
	}
	conn.Stop()
	delete(cm.conns[sName], sid)
	cm.count--
	return nil
}

func (cm *ConnMgr) Count() int {
	return cm.count

}

func (cm *ConnMgr) ClearConn() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	for _, sConns := range cm.conns {
		for _, conn := range sConns {
			conn.Stop()
			cm.Remove(conn)
		}
	}
}

func (cm *ConnMgr) RandomSelectConnWithCond(cond map[string]interface{}) (itface.IConnection, error) {
	if _, ok := cond["sname"]; !ok {
		err := fmt.Errorf("random select conn with cond err,cond not key sname")
		return nil, err
	}
	sName := cond["sname"].(string)
	if _, ok := cm.conns[sName]; !ok {
		err := fmt.Errorf("connmgr not key : %v\n", sName)
		return nil, err
	}
	if len(cm.conns[sName]) == 0 {
		err := fmt.Errorf("connmgr service: %v size = 0\n", sName)
		return nil, err
	}
	j := 0
	keys := make([]int, len(cm.conns[sName]))
	for k := range cm.conns[sName] {
		keys[j] = k
		j++
	}
	index := rand.Intn(len(keys))
	return cm.conns[sName][keys[index]], nil
}

// 获取所有满足条件的conns
func (cm *ConnMgr) GetMeetCondConns(cond map[string]interface{}) ([]itface.IConnection, error) {
	if _, ok := cond["sname"]; !ok {
		err := fmt.Errorf("get meet cond conns err,cond not key sname")
		return nil, err
	}
	sName := cond["sname"].(string)
	connMap := cm.conns[sName]
	connList := []itface.IConnection{}
	for _, conn := range connMap {
		connList = append(connList, conn)
	}
	return connList, nil
}

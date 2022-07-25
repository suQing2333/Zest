package netutil

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"zest/engine/itface"
)

// 连接管理器
type ConnMgr struct {
	conns      map[int64]itface.IConnection
	msgHandler func(req itface.IRequest)
	connLock   sync.RWMutex
}

func NewConnManager() *ConnMgr {
	return &ConnMgr{
		conns: make(map[int64]itface.IConnection),
	}
}

func (cm *ConnMgr) SetMsgHandler(funcHook func(itface.IRequest)) {
	cm.msgHandler = funcHook
}

func (cm *ConnMgr) GetMsgHandler() func(itface.IRequest) {
	return cm.msgHandler
}

// 添加连接
func (cm *ConnMgr) Add(conn itface.IConnection) error {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	cm.conns[conn.GetConnID()] = conn
	return nil
}

func (cm *ConnMgr) Remove(conn itface.IConnection) error {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	conn.Stop()
	delete(cm.conns, conn.GetConnID())
	return nil
}

func (cm *ConnMgr) RemoveWithConnID(connID int64) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	if conn, ok := cm.conns[connID]; ok {
		conn.Stop()
		delete(cm.conns, connID)
	}
}

func (cm *ConnMgr) Count() int {
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()
	length := len(cm.conns)
	return length
}

func (cm *ConnMgr) GetConn(selectInfo map[string]interface{}) (itface.IConnection, error) {
	if _, ok := selectInfo["connid"]; !ok {
		err := fmt.Errorf("can not selecet conn,not find connid")
		return nil, err
	}
	connID := selectInfo["connid"].(int64)
	cm.connLock.RLock()
	defer cm.connLock.RUnlock()

	if conn, ok := cm.conns[connID]; ok {
		return conn, nil
	}

	return nil, errors.New("connection not found")
}

// 清除所有连接
func (cm *ConnMgr) ClearConn() {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()

	for connID, conn := range cm.conns {
		conn.Stop()
		delete(cm.conns, connID)
	}
}

func (cm *ConnMgr) RandomSelectConnWithCond(cond map[string]interface{}) (itface.IConnection, error) {
	cm.connLock.Lock()
	defer cm.connLock.Unlock()
	return cm.conns[int64(rand.Intn(len(cm.conns)))], nil
}

// 获取所有满足条件的conns
func (cm *ConnMgr) GetMeetCondConns(cond map[string]interface{}) ([]itface.IConnection, error) {
	connList := []itface.IConnection{}
	return connList, nil
}

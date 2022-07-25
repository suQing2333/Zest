package netutil

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
	"zest/engine/itface"
	"zest/engine/uuid"
)

type Connection struct {
	Conn   *net.TCPConn
	ConnID int64
	ctx    context.Context
	cancel context.CancelFunc
	// 缓冲管道
	msgBuffChan chan []byte
	isClosed    bool
	ConnMgr     itface.IConnMgr
	ExtraInfo   map[string]interface{} // 额外信息
	sync.RWMutex
}

func NewConnection(conn *net.TCPConn, connMgr itface.IConnMgr) *Connection {
	c := &Connection{
		Conn:        conn,
		ConnID:      uuid.GetNextVal(),
		ConnMgr:     connMgr,
		isClosed:    false,
		msgBuffChan: make(chan []byte, 1024),
		ExtraInfo:   nil,
	}
	c.ConnMgr.Add(c)
	return c
}

func (c *Connection) GetConnID() int64 {
	return c.ConnID
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) SetExtraInfo(extraInfo map[string]interface{}) {
	c.ExtraInfo = extraInfo
}

func (c *Connection) GetExtraInfo() map[string]interface{} {
	return c.ExtraInfo
}

// 写数据
func (c *Connection) StartWriter() {
	fmt.Println("Start writer connID :", c.ConnID)
	for {
		select {
		case data, ok := <-c.msgBuffChan:
			if ok {
				if _, err := c.Conn.Write(data); err != nil {
					fmt.Println("Send Buff Data error:", err)
					return
				}
			} else {
				fmt.Println("msgBuffChan is Closed")
				break
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c *Connection) StartReader() {
	fmt.Println("Start reader connID :", c.ConnID)
	defer c.Stop()

	// 创建拆包解包的对象
	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			//读取客户端的Msg head
			headData := make([]byte, 8)
			if _, err := io.ReadFull(c.Conn, headData); err != nil {
				fmt.Println("read msg head error ", err)
				return
			}
			msg, err := Unpack(headData)
			if err != nil {
				fmt.Println("unpack error ", err)
				return
			}

			//根据解析出的 dataLen 读取 data，放在msg.Data中
			var data []byte
			if msg.GetDataLen() > 0 {
				data = make([]byte, msg.GetDataLen())
				if _, err := io.ReadFull(c.Conn, data); err != nil {
					fmt.Println("read msg data error ", err)
					return
				}
			}
			msg.SetData(data)
			req := NewRequest(c, msg)
			msgHandler := c.ConnMgr.GetMsgHandler()
			msgHandler(req)
		}
	}
}

func (c *Connection) Start() {
	c.ctx, c.cancel = context.WithCancel(context.Background())

	go c.StartWriter()
	go c.StartReader()

	select {
	case <-c.ctx.Done():
		c.finalizer()
		return
	}
}

func (c *Connection) Stop() {
	c.cancel()
}

// 发送消息到客户端
func (c *Connection) SendMsg(msg itface.IMessage) error {
	c.RLock()
	defer c.RUnlock()
	if c.isClosed == true {
		return errors.New("connection closed when send msg")
	}

	//将data封包，并且发送
	out, err := Pack(msg)
	if err != nil {
		return errors.New("Pack error msg ")
	}

	//写回客户端
	_, err = c.Conn.Write(out)
	return err
}

//SendBuffMsg  发生BuffMsg
func (c *Connection) SendBuffMsg(msg itface.IMessage) error {
	c.RLock()
	defer c.RUnlock()
	idleTimeout := time.NewTimer(5 * time.Millisecond)
	defer idleTimeout.Stop()

	if c.isClosed == true {
		return errors.New("Connection closed when send buff msg")
	}

	//将data封包，并且发送
	out, err := Pack(msg)
	if err != nil {
		return errors.New("Pack error msg ")
	}

	// 发送超时
	select {
	case <-idleTimeout.C:
		return errors.New("send buff msg timeout")
	case c.msgBuffChan <- out: //写回客户端
		return nil
	}

	return nil
}

// 关闭连接
func (c *Connection) finalizer() {
	c.Lock()
	defer c.Unlock()

	if c.isClosed == true {
		return
	}

	// 关闭socket链接
	_ = c.Conn.Close()
	c.ConnMgr.Remove(c)
	//关闭该链接全部管道
	close(c.msgBuffChan)
	//设置标志位
	c.isClosed = true
}

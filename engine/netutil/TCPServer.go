package netutil

import (
	"fmt"
	"net"
	"zest/engine/itface"
	// "zest/engine/zslog"
)

type TCPServer struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	ConnMgr   itface.IConnMgr
	AllowOut  bool
}

func NewTCPServer(service string, prot int, allowOut bool) *TCPServer {
	s := &TCPServer{
		Name:      service,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      prot,
		ConnMgr:   NewConnManager(),
		AllowOut:  allowOut,
	}
	return s
}

// 启动一个协程去监听端口, 如果端口冲突直接panic
func (s *TCPServer) Start() {
	fmt.Println("TCPServer start ")
	// zslog.LogInfo("[START] TCPServer Name:%v ,IP : %v, Port: %v", s.Name, s.IP, s.Port)
	go func() {
		address, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr err: ", err)
			panic("resolve tcp addr err")
			return
		}
		listener, err := net.ListenTCP(s.IPVersion, address)
		if err != nil {
			fmt.Println("listen tcp err:", err)
			panic("listen tcp err")
			return
		}

		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err: ", err)
				continue
			}
			if !s.AllowOut && !HasLocalIPAddr(conn.RemoteAddr().String()) {
				conn.Close()
				continue
			}
			dealConn := NewConnection(conn, s.ConnMgr)
			go dealConn.Start()
		}

	}()
}

func (s *TCPServer) Stop() {
	// 清除所有连接
	s.ConnMgr.ClearConn()
}

func (s *TCPServer) Serve() {
	s.Start()

	select {}
}

func (s *TCPServer) SetMsgHandler(hookFunc func(itface.IRequest)) {
	s.ConnMgr.SetMsgHandler(hookFunc)
}

func (s *TCPServer) GetConnMgr() itface.IConnMgr {
	return s.ConnMgr
}

func HasLocalIPAddr(addr string) bool {
	ip, _, err := net.SplitHostPort(addr)
	if err != nil {
		fmt.Println("Split Host Port err : ", err)
		return false
	}
	return HasLocalIP(net.ParseIP(ip))
}

// HasLocalIP 检测 IP 地址是否是内网地址
// 通过直接对比ip段范围效率更高
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}

	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}

	return ip4[0] == 10 || // 10.0.0.0/8
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) || // 172.16.0.0/12
		(ip4[0] == 169 && ip4[1] == 254) || // 169.254.0.0/16
		(ip4[0] == 192 && ip4[1] == 168) // 192.168.0.0/16
}

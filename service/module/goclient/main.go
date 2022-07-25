package main

import (
	"fmt"
	"net"
	"os"
	"time"
	"zest/engine/netutil"
	"zest/engine/pbmgr"
	_ "zest/service/protoc"
)

func main() {
	server := "81.69.250.214:10000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)

	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: ", err)
		os.Exit(1)
	}
	for i := 0; i < 1; i++ {
		go func() {
			//建立服务器连接
			conn, err := net.DialTCP("tcp", nil, tcpAddr)

			if err != nil {
				fmt.Println(conn.RemoteAddr().String(), os.Stderr, "Fatal error:", err)
				// os.Exit(1)
				return
			}

			conn.SetKeepAlive(true)
			conn.SetKeepAlivePeriod(3 * time.Second)

			fmt.Println("connection success")
			go reader(conn)
			for {
				sender(conn)
				time.Sleep(time.Second * 10)
			}
		}()
	}
	select {}
}

func reader(conn *net.TCPConn) {
	for {
		buffer := make([]byte, 1024)
		_, err := conn.Read(buffer)
		if err != nil {
			return
		}
		fmt.Println(buffer)
	}
}

func sender(conn *net.TCPConn) {
	data := pbmgr.GetDataByFullName("protoc.CSProto", int32(1), int32(1002), []byte("this is a test"))
	fmt.Println("send proto ", data)

	msg := netutil.NewMsg(uint32(1), data)

	out, err := netutil.Pack(msg)
	if err != nil {
		fmt.Println(err)
	}
	_, err = conn.Write(out)

	if err != nil {
		fmt.Println("send msg err", err)
	}
}

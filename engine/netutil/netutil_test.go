package netutil

import (
	"fmt"
	"testing"
	"zest/engine/itface"
)

func TestNet(t *testing.T) {

	// testTCP()
}

func testPacket() {
	s1 := "test"
	dataBytes := []byte(s1)
	msg := NewMsg(1, dataBytes)
	out, _ := Pack(msg)
	fmt.Println(out)

	head, _ := Unpack(out[0:4])
	fmt.Println(head)
}

func testTCP() {
	s := NewTCPServer("gate", 10000, true)
	s.SetMsgHandler(testMsgHandler)
	s.Serve()
}

func testMsgHandler(req itface.IRequest) {

}

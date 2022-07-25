package sys

import (
	"fmt"
	"testing"
	// "zest/engine/funcmgr"
)

func testFunc(args ...interface{}) interface{} {
	fmt.Println("call testFunc args: ", args)
	return map[int]int{1: 2, 3: 4}
}

func TestSys(t *testing.T) {
	// testEtcd()
	s := NewService(10000, "demo")
	fmt.Printf("%v\n", s)
	s.Start()

	select {}
}

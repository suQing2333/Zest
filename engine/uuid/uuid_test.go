package uuid

import (
	"fmt"
	"testing"
	"time"
)

func TestUUID(t *testing.T) {
	for i := 0; i < 5; i++ {
		fmt.Println(GetNextVal())
	}
	time.Sleep(10 * time.Second)
	fmt.Println(GetNextVal())
}

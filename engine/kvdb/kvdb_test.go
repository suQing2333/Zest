package kvdb

import (
	"fmt"
	"testing"
)

func TestController(t *testing.T) {
	testRedis(t)
	testMongo(t)
}

func testRedis(t *testing.T) {
	Redis("hset", "test1", "5", "6")
}
func testMongo(t *testing.T) {
	// fmt.Println(mc)
	Connect()
	Close()
}

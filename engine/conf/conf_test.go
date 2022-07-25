package conf

import (
	"fmt"
	"testing"
)

func TestConf(t *testing.T) {
	SetConfigFile("//127.0.0.1/suqing/zest/zest.yaml")
	LoadConfig()
	fmt.Println(GetString("mongo.ip"))
}

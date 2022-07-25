package funcmgr

import (
	"fmt"
	"testing"
)

type Routers struct {
}

func (this *Routers) Ftest1(msg string) {
	fmt.Println("Login:", msg)
}

func (this *Routers) Ftest2(msg1 int, msg2 string, msg3 bool) (string, string) {
	fmt.Println(msg1, msg2, msg3)
	return "test string res", "more res"
}

func (this *Routers) Ftest3(msg1 string) map[string]interface{} {
	return map[string]interface{}{"msg": msg1}
}

func TestFuncMgr(t *testing.T) {
	testCall()
}

func testCall() {
	// var ruTest Routers
	ruTest := &Routers{}
	RegisterFunc(*ruTest, ruTest)
	// args := []interface{}{10, "this is a test", false}
	res, err := CallFunc("Routers.Ftest1", "this is a test")

	if err != nil {
		fmt.Println(err)
	}
	if res != nil {
		fmt.Println("123")
	}
	fmt.Println(res)

	res, err = CallFunc("Routers.Ftest3", "this is a test")
	if err != nil {
		fmt.Println(err)
	}
	if res != nil {
		fmt.Println("123")
	}
	fmt.Println(res)

}

package common

import (
	"fmt"
	"testing"
)

func TestController(t *testing.T) {
	// testTypeCast()
	// testShellTools()
	// testFunction()
	testSerial()
}

func testShellTools() {
	err := Command("ipconfig")
	fmt.Println(err)
}

func testTypeCast() {
	var tmp1 int
	tmp1 = 100
	fmt.Println(Typeof(tmp1))
	tmp2 := Float(tmp1)
	tmp3 := String(tmp1)
	fmt.Println(Typeof(tmp2))
	fmt.Println(Typeof(tmp3))
}

func testFunction() {
	path, err := GetProgrammePath()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ProgrammePath", path)

	a := 5
	b := 6
	fmt.Println(ThreeUnary(a > b, a, b))
	// Println(ThreeUnary(a < b, a, b))
}

func testSerial() {
	testSerial := map[interface{}]interface{}{
		"testString":         "testString",
		2:                    5,
		"testBool":           true,
		"testMap":            map[string]int{"t": 123},
		"testInterfaceSlice": []interface{}{1, 2, 3, 4, 5},
	}

	out, err := Serialization(testSerial)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(out)
	res, err := Deserialization(out)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

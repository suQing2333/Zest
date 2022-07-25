package zslog

import "testing"

func TestGWLog(t *testing.T) {
	test := "123456"
	LogDebug("test1 = %v", test)

	test2 := []int{1, 2, 3, 4, 5}
	LogInfo("test2 = %v", test2)
	LogInfo("this is a Info")

	LogWarn("this is a Warn")

	LogError("this is a error")

	// LogPanic("this is a foreseeable panic")

	// LogDebug("testPanic panic = %v", test2[5])

}

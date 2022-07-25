package common

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"sync"
)

// 调用shell命令,注意window与linux使用的命令不同
func Command(cmd string) error {
	// c := exec.Command("cmd", "/C", cmd)  // windows
	c := exec.Command("bash", "-c", cmd) // mac or linux
	stdout, err := c.StdoutPipe()
	if err != nil {
		return err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		reader := bufio.NewReader(stdout)
		for {
			readString, err := reader.ReadString('\n')
			if err != nil || err == io.EOF {
				return
			}
			byte2String := ConvertByte2String([]byte(readString), "GB18030")
			fmt.Println(byte2String)
		}
	}()
	err = c.Start()
	wg.Wait()
	return err
}

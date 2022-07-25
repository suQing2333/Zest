// +build linux

package binutil

import (
	"fmt"
	"github.com/sevlyar/go-daemon"
	"os"
	"syscall"
)

func RedirectStderr(f *os.File) {
	err := syscall.Dup2(int(f.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		return
	}
}

func Daemonize() *daemon.Context {
	context := new(daemon.Context)
	child, err := context.Reborn()

	if err != nil {
		// daemonize failed
		fmt.Println(err)
	}

	if child != nil {
		os.Exit(0)
		return nil
	} else {
		return context
	}
}

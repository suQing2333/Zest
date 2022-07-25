// +build windows

package binutil

import (
	"fmt"
	"os"
	"syscall"
)

// redirect panic
func setStdHandle(stdhandle int32, handle syscall.Handle) error {
	kernel32 := syscall.MustLoadDLL("kernel32.dll")
	procSetStdHandle := kernel32.MustFindProc("SetStdHandle")
	r0, _, e1 := syscall.Syscall(procSetStdHandle.Addr(), 2, uintptr(stdhandle), uintptr(handle), 0)
	if r0 == 0 {
		if e1 != 0 {
			return error(e1)
		}
		return syscall.EINVAL
	}
	return nil
}

func RedirectStderr(f *os.File) error {
	err := setStdHandle(syscall.STD_ERROR_HANDLE, syscall.Handle(f.Fd()))
	if err != nil {
		return error(err)
	}
	os.Stderr = f
	return nil
}

type nopRelease int

func (_ nopRelease) Release() {

}

func Daemonize() nopRelease {
	// Windows can not daemonize
	fmt.Println("can not run in daemon mode in windows, -d ignored")
	return nopRelease(0)
}

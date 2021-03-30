// common
package exec

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"os/exec"
	"path/filepath"
)

// 错误处理函数
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func IsProcessExist(pid int) (bool) {
	process, err := os.FindProcess(pid)
	if err != nil {
		//fmt.Printf("Failed to find process: %s\n", err)
	} else {
		if runtime.GOOS == "windows" {
			//fmt.Printf("The process %d exists!\r\n", pid)
			return true
		}

		err := process.Signal(syscall.Signal(0))
		//fmt.Printf("process.Signal on pid %d returned: %v\n", pid, err)
		if err == nil {
			//fmt.Printf("The process %d exists!\r\n", pid)
			return true
		}
	}

	return false
}

func GetExePath() (string, error) {
	exeFile, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	exePath, err := filepath.Abs(exeFile)
	if err != nil {
		return "", err
	} else {
		return filepath.Dir(exePath), nil
	}
}
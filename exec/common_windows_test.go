// common
package exec

import (
	"strings"
	"testing"
)

func TestExecShell(t *testing.T) {
	// 测试反馈结果是否正常
	toDis := "hello"
	cmdStr := "echo " + toDis

	rtn, err := ExecShell(cmdStr)

	if err == nil {
		rtn = strings.TrimSpace(rtn)
		if strings.Compare(toDis, rtn) != 0 {
			t.Errorf("ExecShell(%s) failed. Got %s, expected %s.", cmdStr, rtn, toDis)
		}
	} else {
		t.Error(err)
	}
}

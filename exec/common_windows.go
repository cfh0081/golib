// common
package exec

import (
	"os/exec"
)

func ExecShell(s string) (string, error) {
	cmd := exec.Command("cmd", "/C", s)
	stdout, err := cmd.Output()

	var result string = ""
	if err == nil {
		result = string(stdout)
	}

	return result, err
}

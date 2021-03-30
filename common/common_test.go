// common_test.go
package common

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestExists(t *testing.T) {
	file := "./hello.exe"
	path, _ := os.Getwd()
	fmt.Println(path)
	fmt.Println(filepath.Abs(file))

	r := Exists(file)
	if r {
		t.Errorf("Exists(%s) failed. Got %t, expected false.", file, r)
	}

	file = "./common.go"
	r = Exists(file)
	if !r {
		t.Errorf("Exists(%s) failed. Got %t, expected true.", file, r)
	}
}

func TestPureName(t *testing.T) {
	name := "what.bat"
	expect := "what"
	pure := PureName(name)
	r := pure == expect
	if !r {
		t.Errorf("PureName(%s) failed. Got %s, expected %s.", name, pure, expect)
	}

	name = "name"
	expect = "name"
	pure = PureName(name)
	r = pure == expect
	if !r {
		t.Errorf("PureName(%s) failed. Got %s, expected %s.", name, pure, expect)
	}
}

package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunFuncName(t *testing.T) {

	name := RunFuncName()
	expected := `github.com/cfh0081/golib/common.TestRunFuncName`
	assert.Equal(t, name, expected)
}

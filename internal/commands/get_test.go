package commands

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCmd(t *testing.T) {
	setUp()
	defer tearDown()

	cmd := GetCmdFactory()
	buf := bytes.NewBufferString("")
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"FOO"})
	cmd.Execute()
	out, _ := ioutil.ReadAll(buf)

	replacer := strings.NewReplacer("-", "", "+", "", " ", "", "\n", "")
	expectedResult := "|NAME|LOGIN|PASSWORD|DESCRIPTION||FOO|BAR|PASS|DESC|"
	assert.Equal(t, expectedResult, replacer.Replace(string(out)))
}

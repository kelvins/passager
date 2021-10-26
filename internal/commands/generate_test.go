package commands

import (
	"bytes"
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func GenerateCmdFixture(args []string) string {
	cmd := GenerateCmdFactory()
	buf := bytes.NewBufferString("")
	cmd.SetOut(buf)
	cmd.SetArgs(args)
	cmd.Execute()
	out, _ := ioutil.ReadAll(buf)
	return string(out)
}

func TestGenerateCmdDefault(t *testing.T) {
	out := GenerateCmdFixture([]string{})
	assert.Equal(t, 19, len(out))
}

func TestGenerateCmdChangeLength(t *testing.T) {
	out := GenerateCmdFixture([]string{"-l", "24"})
	assert.Equal(t, 25, len(out))
}

func TestGenerateCmdNoNumbers(t *testing.T) {
	out := GenerateCmdFixture([]string{"--no-numbers"})
	result, err := regexp.MatchString("[0-9]", out)
	assert.Nil(t, err)
	assert.False(t, result)
}

func TestGenerateCmdNoSymbols(t *testing.T) {
	out := GenerateCmdFixture([]string{"--no-symbols"})
	result, err := regexp.MatchString("[!@#$]", out)
	assert.Nil(t, err)
	assert.False(t, result)
}

package commands

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestRootCmd(t *testing.T) {
	cmd := RootCmdFactory()
	buf := bytes.NewBufferString("")
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"--version"})
	cmd.Execute()
	out, _ := ioutil.ReadAll(buf)
	assert.Equal(t, string(out), fmt.Sprintf("Passager v%s", Version))
}

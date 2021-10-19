package commands

import (
	"fmt"
	"testing"
	"bytes"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
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

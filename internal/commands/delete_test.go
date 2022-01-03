package commands

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/kelvins/passager/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestDeleteCmd(t *testing.T) {
	setUp()
	defer tearDown()

	cred, err := models.ReadAll("FOO")
	assert.Nil(t, err)
	assert.Equal(t, len(cred), 1)

	cmd := DeleteCmdFactory()
	buf := bytes.NewBufferString("")
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"FOO"})
	cmd.Execute()
	out, _ := ioutil.ReadAll(buf)
	assert.Equal(t, "Credential FOO successfully deleted!\n", string(out))

	cred, err = models.ReadAll("FOO")
	assert.Nil(t, err)
	assert.Equal(t, len(cred), 0)
}

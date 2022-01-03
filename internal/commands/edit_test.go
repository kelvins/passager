package commands

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/kelvins/passager/internal/models"
	"github.com/kelvins/passager/pkg/crypto"
	"github.com/stretchr/testify/assert"
)

func TestEditCmd(t *testing.T) {
	setUp()
	defer tearDown()

	cmd := EditCmdFactory()
	buf := bytes.NewBufferString("")
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"FOO", "-l", "BAR2", "-p", "PASS2", "-d", "DESC2"})
	cmd.Execute()
	out, _ := ioutil.ReadAll(buf)
	assert.Equal(t, "Credential FOO successfully updated!\n", string(out))

	cred, err := models.ReadFirst("FOO")
	assert.Nil(t, err)
	assert.Equal(t, "FOO", cred.Name)
	assert.Equal(t, "BAR2", cred.Login)
	assert.Equal(t, "PASS2", crypto.Decrypt(cred.Password, crypto.EncryptionKey()))
	assert.Equal(t, "DESC2", cred.Description)
}

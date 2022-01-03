package commands

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"

	"github.com/kelvins/passager/internal/models"
	"github.com/kelvins/passager/pkg/crypto"
	"github.com/stretchr/testify/assert"
)

func setUp() {
	os.Setenv("PASSAGER_DATABASE", "/tmp/passager-test.db")

	credential := models.Credential{
		Name:        "FOO",
		Login:       "BAR",
		Password:    crypto.Encrypt("PASS", crypto.EncryptionKey()),
		Description: "DESC",
	}
	models.Create(&credential)
}

func tearDown() {
	os.Unsetenv("PASSAGER_DATABASE")
	os.Remove("/tmp/passager-test.db")
}

func TestAddCmd(t *testing.T) {
	setUp()
	defer tearDown()
	cmd := AddCmdFactory()
	buf := bytes.NewBufferString("")
	cmd.SetOut(buf)
	cmd.SetArgs([]string{"TEST", "-l", "FOO", "-p", "BAR"})
	cmd.Execute()
	out, _ := ioutil.ReadAll(buf)
	assert.Equal(t, "Credential TEST successfully created!\n", string(out))
	cred, err := models.ReadAll("TEST")
	assert.Nil(t, err)
	assert.Equal(t, len(cred), 1)
	assert.Equal(t, "TEST", cred[0].Name)
	assert.Equal(t, "FOO", cred[0].Login)
	assert.Equal(t, "BAR", crypto.Decrypt(cred[0].Password, crypto.EncryptionKey()))
}

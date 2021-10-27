package commands

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/kelvins/passager/internal/models"
	"github.com/kelvins/passager/pkg/crypto"
	"github.com/stretchr/testify/assert"
)

func TestGetCmd(t *testing.T) {
	setUp()
	defer tearDown()

	cred, err := models.ReadAll("FOO")
	assert.Nil(t, err)
	assert.Equal(t, len(cred), 0)

	credential := models.Credential{
		Name:        "FOO",
		Login:       "BAR",
		Password:    crypto.Encrypt("PASS", crypto.EncryptionKey()),
		Description: "DESC",
	}
	models.Create(&credential)

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

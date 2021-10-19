package models

import (
	"github.com/stretchr/testify/assert"
	"os"
	"log"
	"fmt"
	"testing"
)

func setUp() {
	os.Setenv("PASSAGER_DATABASE", "/tmp/passager-test.db")
}

func tearDown() {
	os.Remove("/tmp/passager-test.db")
}

func fakeCredential() Credential {
	return Credential{
		Name:        "name",
		Login:       "login",
		Password:    "password",
		Description: "description",
	}
}

func assertEqualCredential(t *testing.T, first, second Credential) {
	assert.Equal(t, first.Name, second.Name)
	assert.Equal(t, first.Login, second.Login)
	assert.Equal(t, first.Password, second.Password)
	assert.Equal(t, first.Description, second.Description)
}

func TestDatabasePath(t *testing.T) {
	os.Unsetenv("PASSAGER_DATABASE")
	dirname, _ := os.UserHomeDir()
	dbPath := fmt.Sprintf("%s/.passager.db", dirname)
	assert.Equal(t, databasePath(), dbPath)
}

func TestCreate(t *testing.T) {
	setUp()
	defer tearDown()
	credential := fakeCredential()
	Create(&credential)

	conn := openConnection()
	var loadedCredential Credential
	conn.First(&loadedCredential, "name = name")
	assertEqualCredential(t, credential, loadedCredential)
}

func TestReadAll(t *testing.T) {
	setUp()
	defer tearDown()
	conn := openConnection()
	credential := fakeCredential()
	conn.Create(&credential)

	credentials, err := ReadAll("name")
	if err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, len(credentials), 1, "length should be one")
	assertEqualCredential(t, credential, credentials[0])
}

func TestDelete(t *testing.T) {
	setUp()
	defer tearDown()
	conn := openConnection()
	credential := fakeCredential()
	conn.Create(&credential)

	var loadedCredential Credential
	conn.First(&loadedCredential, "name = name")
	assertEqualCredential(t, credential, loadedCredential)

	err := Delete("name")
	if err != nil {
		log.Fatal(err)
	}
	err = conn.First(&loadedCredential, "name = name").Error
	assert.Equal(t, fmt.Sprint(err), "record not found", "message should be record not found")
}

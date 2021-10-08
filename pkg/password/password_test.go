package password

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	pass1 := Generate(10, true, true)
	pass2 := Generate(10, true, true)
	if pass1 == pass2 {
		t.Errorf("Password %s and %s should be different", pass1, pass2)
	}
}

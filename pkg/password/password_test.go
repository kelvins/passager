package password

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTimeNowSeed(t *testing.T) {
	password1 := Generate(10, true, true)
	password2 := Generate(10, true, true)
	assert.NotEqual(t, password1, password2, "passwords should not be equal")
}

func TestGenerateStritngLength(t *testing.T) {
	testCases := []struct{
		length int8
		numbers bool
		symbols bool
	}{
		{10, false, false},
		{20, true, false},
		{30, false, true},
		{40, true, true},
	}
	for _, tc := range testCases {
		assert.Equal(t, int(tc.length), len(Generate(tc.length, tc.numbers, tc.symbols)), "length should be equal")
	}
}

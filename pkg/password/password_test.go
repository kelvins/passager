package password

import (
	"github.com/stretchr/testify/assert"
	"regexp"
	"strings"
	"testing"
)

func TestGenerateTimeNowSeed(t *testing.T) {
	password1 := Generate(10, true, true)
	password2 := Generate(10, true, true)
	assert.NotEqual(t, password1, password2, "passwords should not be equal")
}

func TestGenerateStritngLength(t *testing.T) {
	testCases := []struct {
		length  int8
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

func TestGenerateLetters(t *testing.T) {
	var isStringLetters = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
	data := strings.Join(generateLetters(20), "")
	assert.True(t, isStringLetters(data), "data should have only letters")
}

func TestGenerateNumbers(t *testing.T) {
	var isStringNumbers = regexp.MustCompile(`^[0-9]+$`).MatchString
	data := strings.Join(generateNumbers(20), "")
	assert.True(t, isStringNumbers(data), "data should have only numbers")
}

func TestGenerateSymbols(t *testing.T) {
	var isStringSymbols = regexp.MustCompile(`^[!@#$]+$`).MatchString
	data := strings.Join(generateSymbols(20), "")
	assert.True(t, isStringSymbols(data), "data should have only symbols")
}

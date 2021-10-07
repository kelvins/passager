package password

import (
	"time"
	"math"
	"strings"
	"math/rand"
)

func generateRandomSequence(sequence string, length int8) []string {
	result := make([]byte, length)
	for i := range result {
		result[i] = sequence[rand.Intn(len(sequence))]
	}
	return strings.Split(string(result), "")
}

func generateLetters(length int8) []string {
	sequence := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return generateRandomSequence(sequence, length)
}

func generateNumbers(length int8) []string {
	sequence := "0123456789"
	return generateRandomSequence(sequence, length)
}

func generateSymbols(length int8) []string {
	sequence := "!#$@"
	return generateRandomSequence(sequence, length)
}

func Generate(length int8, numbers, symbols bool) string {
	rand.Seed(time.Now().UnixNano())

	numbersRate := 0.3
	symbolsRate := 0.1

	numbersLength := int8(math.Round(float64(length) * numbersRate))
	symbolsLength := int8(math.Round(float64(length) * symbolsRate))
	lettersLength := length - (numbersLength + symbolsLength)

	var result []string

	result = append(result, generateLetters(lettersLength)...)
	result = append(result, generateNumbers(numbersLength)...)
	result = append(result, generateSymbols(symbolsLength)...)

	rand.Shuffle(len(result), func(i, j int) {result[i], result[j] = result[j], result[i]})

	return strings.Join(result, "")
}

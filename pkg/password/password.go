package password

import (
	"math"
	"math/rand"
	"strings"
	"time"
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

	var result []string

	var numbersLength int8
	if numbers {
		numbersRate := 0.3
		numbersLength = int8(math.Round(float64(length) * numbersRate))
		result = append(result, generateNumbers(numbersLength)...)
	}

	var symbolsLength int8
	if symbols {
		symbolsRate := 0.1
		symbolsLength = int8(math.Round(float64(length) * symbolsRate))
		result = append(result, generateSymbols(symbolsLength)...)
	}

	lettersLength := length - (numbersLength + symbolsLength)
	result = append(result, generateLetters(lettersLength)...)

	rand.Shuffle(len(result), func(i, j int) { result[i], result[j] = result[j], result[i] })

	return strings.Join(result, "")
}

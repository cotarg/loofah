package main

import (
	"strings"
	"testing"
)

func TestEmojifier(t *testing.T) {
	for _, testCase := range []struct {
		input          string
		expectedOutput string
	}{
		{input: "a", expectedOutput: "😻"},
		{input: "aba", expectedOutput: "😻b😻"},
		{input: "cbc", expectedOutput: "cbc"},
	} {
		output := emojifier(testCase.input)
		if output != testCase.expectedOutput {
			t.Errorf("emojifier failed, output %s does not match expected %s", output, testCase.expectedOutput)
		}
	}
}

func BenchmarkEmojifier1000(b *testing.B) {
	input := strings.Repeat("a", 1000)
	for i := 0; i < b.N; i++ {
		emojifier(input)
	}
}

func BenchmarkEmojifier10000(b *testing.B) {
	input := strings.Repeat("a", 10000)
	for i := 0; i < b.N; i++ {
		emojifier(input)
	}
}

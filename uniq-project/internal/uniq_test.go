package uniq

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUniq(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		options  Options
		expected string
	}{
		{
			name:     "Basic unique lines",
			input:    "a\nb\na\nc\nb\n",
			options:  Options{},
			expected: "a\nb\nc\n",
		},
		{
			name:     "Count occurrences",
			input:    "a\nb\na\nc\nb\n",
			options:  Options{Count: true},
			expected: "2 a\n2 b\n1 c\n",
		},
		{
			name:     "Only duplicates",
			input:    "a\nb\na\nc\nb\n",
			options:  Options{Repeated: true},
			expected: "a\nb\n",
		},
		{
			name:     "Only unique",
			input:    "a\nb\na\nc\nb\n",
			options:  Options{Unique: true},
			expected: "c\n",
		},
		{
			name:     "Ignore case",
			input:    "A\nb\na\nB\nc\n",
			options:  Options{IgnoreCase: true},
			expected: "A\nb\nc\n",
		},
		{
			name:     "Skip fields",
			input:    "1 apple\n2 banana\n1 apricot\n3 apple\n",
			options:  Options{NumFields: 1},
			expected: "apple\nbanana\napricot\n",
		},
		{
			name:     "Skip chars",
			input:    "apple\nbanana\napricot\napple\n",
			options:  Options{NumChars: 1},
			expected: "pple\nanana\npricot\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := strings.NewReader(tt.input)
			var output bytes.Buffer

			err := Run(input, &output, tt.options)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, output.String())
		})
	}
}

func TestProcessLine(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		options  Options
		expected string
	}{
		{
			name:     "Skip 1 field",
			input:    "hello world test",
			options:  Options{NumFields: 1},
			expected: "world test",
		},
		{
			name:     "Skip 2 chars",
			input:    "hello",
			options:  Options{NumChars: 2},
			expected: "llo",
		},
		{
			name:     "Ignore case",
			input:    "Hello",
			options:  Options{IgnoreCase: true},
			expected: "hello",
		},
		{
			name:     "Skip fields and chars",
			input:    "a b c d",
			options:  Options{NumFields: 2, NumChars: 1},
			expected: " d",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := processLine(tt.input, tt.options)
			assert.Equal(t, tt.expected, result)
		})
	}
}
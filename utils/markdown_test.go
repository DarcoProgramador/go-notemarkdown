package utils_test

import (
	"bytes"
	"testing"

	"github.com/Darcoprogramador/go-notemarkdown/utils"
)

func TestMdToHTML(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:     "Test 1",
			input:    []byte("# Hello, World!"),
			expected: []byte("<h1 id=\"hello-world\">Hello, World!</h1>\n"),
		},
		{
			name:     "Test 2",
			input:    []byte("## Hello, World!"),
			expected: []byte("<h2 id=\"hello-world\">Hello, World!</h2>\n"),
		},
		{
			name:     "Test 3",
			input:    []byte("### Hello, World!"),
			expected: []byte("<h3 id=\"hello-world\">Hello, World!</h3>\n"),
		},
		{
			name:     "Test 4",
			input:    []byte("#### Hello, World!"),
			expected: []byte("<h4 id=\"hello-world\">Hello, World!</h4>\n"),
		},
		{
			name:     "Test 5",
			input:    []byte("##### Hello, World!"),
			expected: []byte("<h5 id=\"hello-world\">Hello, World!</h5>\n"),
		},
		{
			name:     "Test 6",
			input:    []byte("###### Hello, World!"),
			expected: []byte("<h6 id=\"hello-world\">Hello, World!</h6>\n"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := utils.MdToHTML(tt.input)

			if !bytes.Equal(got, tt.expected) {
				t.Errorf("mdToHTML() = %s, want %s", got, tt.expected)
			}
		})
	}
}

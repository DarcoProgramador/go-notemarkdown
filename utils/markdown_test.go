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

func TestMdToHtmlHighlight(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{
			name:  "Test 1",
			input: []byte("```python\nprint(\"Hello, World!\")\n```"),
			expected: []byte(`<pre style="color:#f8f8f2;background-color:#272822;"><code><span style="display:flex;"><span style="white-space:pre;-webkit-user-select:none;user-select:none;margin-right:0.4em;padding:0 0.4em 0 0.4em;color:#7f7f7f">1</span><span>print(<span style="color:#e6db74">&#34;Hello, World!&#34;</span>)
</span></span></code></pre>`),
		},
		{
			name:  "Test 2",
			input: []byte("```javascript\nconsole.log(\"Hello, World!\")\n```"),
			expected: []byte(`<pre style="color:#f8f8f2;background-color:#272822;"><code><span style="display:flex;"><span style="white-space:pre;-webkit-user-select:none;user-select:none;margin-right:0.4em;padding:0 0.4em 0 0.4em;color:#7f7f7f">1</span><span><span style="color:#a6e22e">console</span>.<span style="color:#a6e22e">log</span>(<span style="color:#e6db74">&#34;Hello, World!&#34;</span>)
</span></span></code></pre>`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := utils.MdToHtmlHighlight(tt.input)

			if !bytes.Equal(got, tt.expected) {
				t.Errorf("mdToHtmlHighlight() = %s, want %s", got, tt.expected)
			}
		})
	}
}

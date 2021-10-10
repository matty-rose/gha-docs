/*
Copyright © 2021 Matt Rose <matthewrose153@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package formatter_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matty-rose/gha-docs/pkg/formatter"
)

func TestNewMarkdownDocument(t *testing.T) {
	t.Parallel()

	doc := formatter.NewMarkdownDocument()
	assert.Equal(t, "", doc.Render())
}

func TestMarkdownWriteText(t *testing.T) {
	t.Parallel()

	testStrings := []string{"hello", "these are the", "test strings!"}
	for _, s := range testStrings {
		doc := formatter.NewMarkdownDocument()
		doc.WriteText(s)

		assert.Equal(t, s, doc.Render())
	}
}

func TestMarkdownWriteHeading(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		text        string
		level       formatter.MarkdownHeadingLevel
		expectedDoc string
	}{
		{text: "test", level: formatter.H1, expectedDoc: "# test\n"},
		{text: "this heading", level: formatter.H3, expectedDoc: "### this heading\n"},
		{text: "function", level: formatter.H6, expectedDoc: "###### function\n"},
	}

	for _, tc := range testCases {
		doc := formatter.NewMarkdownDocument()
		doc.WriteHeading(tc.text, tc.level)

		assert.Equal(t, tc.expectedDoc, doc.Render())
	}
}

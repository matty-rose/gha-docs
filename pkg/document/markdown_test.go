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
package document_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matty-rose/gha-docs/pkg/document"
)

func TestNewMarkdownDocument(t *testing.T) {
	t.Parallel()

	doc := document.NewMarkdownDocument()
	assert.Equal(t, "", doc.Render())
}

func TestMarkdownWriteText(t *testing.T) {
	t.Parallel()

	testStrings := []string{"hello", "these are the", "test strings!"}
	for _, s := range testStrings {
		doc := document.NewMarkdownDocument()
		doc.WriteText(s)

		assert.Equal(t, s, doc.Render())
	}
}

func TestMarkdownWriteNewLine(t *testing.T) {
	t.Parallel()

	doc := document.NewMarkdownDocument().WriteNewLine()

	assert.Equal(t, "\n", doc.Render())
}

func TestMarkdownWriteHeading(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		text        string
		level       document.MarkdownHeadingLevel
		expectedDoc string
	}{
		{text: "test", level: document.H1, expectedDoc: "# test\n"},
		{text: "this heading", level: document.H3, expectedDoc: "### this heading\n"},
		{text: "function", level: document.H6, expectedDoc: "###### function\n"},
	}

	for _, tc := range testCases {
		doc := document.NewMarkdownDocument()
		doc.WriteHeading(tc.text, tc.level)

		assert.Equal(t, tc.expectedDoc, doc.Render())
	}
}

func TestMarkdownWriteTable(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		columns     []string
		rows        [][]string
		expectedDoc string
		errExpected bool
	}{
		{"single_column_no_rows", []string{"one"}, nil, "| one |\n| --- |\n", false},
		{"multi_column_no_rows", []string{"more", "than", "one"}, nil, "| more | than | one |\n| --- | --- | --- |\n", false},
		{
			"multi_column_multi_row",
			[]string{"two", "columns"},
			[][]string{{"and", "some"}, {"more", "rows"}},
			"| two | columns |\n| --- | --- |\n| and | some |\n| more | rows |\n",
			false,
		},
		{
			"single_column_multi_row_error",
			[]string{"main"},
			[][]string{{"sometimes", "there's", "mismatches"}},
			"",
			true,
		},
		{
			"multi_column_multi_row_error",
			[]string{"main", "backup"},
			[][]string{{"a", "row"}, {"some", "other", "row"}},
			"",
			true,
		},
	}

	for _, tc := range testCases {
		doc := document.NewMarkdownDocument()
		_, err := doc.WriteTable(tc.columns, tc.rows)
		errReceived := err != nil

		if tc.errExpected != errReceived {
			t.Fatalf(
				"Test %s: Error expected <%v>, Error got <%v>",
				tc.name,
				tc.errExpected,
				err,
			)
		}

		assert.Equal(t, tc.expectedDoc, doc.Render())
	}
}

func TestCreateMarkdownLink(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		title        string
		url          string
		expectedLink string
	}{
		{
			"title",
			"www.google.com",
			"[title](www.google.com)",
		},
		{
			"a longer title",
			"www.google.com/a/nested/page",
			"[a longer title](www.google.com/a/nested/page)",
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expectedLink, document.NewMarkdownDocument().CreateLink(tc.title, tc.url))
	}
}

func TestMarkdownFormatCode(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		text         string
		expectedCode string
	}{
		{
			"",
			"",
		},
		{
			"title",
			"`title`",
		},
		{
			"a longer title",
			"`a longer title`",
		},
	}

	for _, tc := range testCases {
		assert.Equal(t, tc.expectedCode, document.NewMarkdownDocument().FormatCode(tc.text))
	}
}

func TestMarkdownWriteCodeBlockMarker(t *testing.T) {
	t.Parallel()

	doc := document.NewMarkdownDocument()
	doc.WriteCodeBlockMarker()
	assert.Equal(t, "```\n", doc.Render())
}

func TestMarkdownWriteCodeBlockMarkerFormat(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		format   string
		expected string
	}{
		{
			"yaml",
			"```yaml\n",
		},
		{
			"python",
			"```python\n",
		},
		{
			"sh",
			"```sh\n",
		},
	}

	for _, tc := range testCases {
		doc := document.NewMarkdownDocument()
		doc.WriteCodeBlockMarkerWithFormat(tc.format)
		assert.Equal(t, tc.expected, doc.Render())
	}
}

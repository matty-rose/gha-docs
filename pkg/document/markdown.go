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
package document

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

type MarkdownDocument struct {
	builder *strings.Builder
}

// NewMarkdownAction returns a new markdown action that wraps the provided composite action data structure,
// and contains a blank string builder.
func NewMarkdownDocument() *MarkdownDocument {
	mda := new(MarkdownDocument)
	mda.builder = new(strings.Builder)

	return mda
}

func (m MarkdownDocument) Render() string {
	return m.builder.String()
}

func (m *MarkdownDocument) WriteText(text string) *MarkdownDocument {
	m.builder.WriteString(text)
	return m
}

func (m *MarkdownDocument) WriteNewLine() *MarkdownDocument {
	m.WriteText("\n")
	return m
}

type MarkdownHeadingLevel int

const (
	H1 MarkdownHeadingLevel = iota + 1
	H2
	H3
	H4
	H5
	H6
)

func (mdh MarkdownHeadingLevel) Value() int {
	return int(mdh)
}

func (m *MarkdownDocument) WriteHeading(text string, level MarkdownHeadingLevel) *MarkdownDocument {
	heading := fmt.Sprintf("%s %s", strings.Repeat("#", level.Value()), text)
	m.WriteText(heading)
	m.WriteNewLine()

	return m
}

func (m *MarkdownDocument) writeTableHeader(columns []string) *MarkdownDocument {
	m.WriteText("|")

	for _, column := range columns {
		m.WriteText(fmt.Sprintf(" %s |", column))
	}

	m.WriteNewLine()
	m.WriteText("|")

	for range columns {
		m.WriteText(" --- |")
	}

	return m
}

func (m *MarkdownDocument) writeTableRows(rows [][]string) *MarkdownDocument {
	for _, row := range rows {
		m.WriteText("|")

		for _, value := range row {
			m.WriteText(fmt.Sprintf(" %s |", value))
		}

		m.WriteNewLine()
	}

	return m
}

func (m *MarkdownDocument) WriteTable(columns []string, rows [][]string) (*MarkdownDocument, error) {
	for _, row := range rows {
		if len(row) != len(columns) {
			return nil, errors.New("each row must have the same number of values as the number of columns")
		}
	}

	m.writeTableHeader(columns)
	m.WriteNewLine()
	m.writeTableRows(rows)

	return m, nil
}

func (m MarkdownDocument) CreateLink(title, url string) string {
	return fmt.Sprintf("[%s](%s)", title, url)
}

func (m MarkdownDocument) FormatCode(text string) string {
	if text == "" {
		return text
	}

	return fmt.Sprintf("`%s`", text)
}

const CodeBlockMarker string = "```"

func (m *MarkdownDocument) WriteCodeBlockMarker() *MarkdownDocument {
	m.WriteText(CodeBlockMarker)
	m.WriteNewLine()

	return m
}

func (m *MarkdownDocument) WriteCodeBlockMarkerWithFormat(format string) *MarkdownDocument {
	m.WriteText(fmt.Sprintf("%s%s", CodeBlockMarker, format))
	m.WriteNewLine()

	return m
}

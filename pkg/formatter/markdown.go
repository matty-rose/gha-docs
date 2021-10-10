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
package formatter

import (
	"fmt"
	"strings"
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

func (m *MarkdownDocument) WriteText(text string) {
	m.builder.WriteString(text)
}

func (m *MarkdownDocument) WriteNewLine() {
	m.WriteText("\n")
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

func (m *MarkdownDocument) WriteHeading(text string, level MarkdownHeadingLevel) {
	heading := fmt.Sprintf("%s %s", strings.Repeat("#", int(level)), text)
	m.WriteText(heading)
	m.WriteNewLine()
}

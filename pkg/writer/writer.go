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
package writer

import (
	"os"
	"strconv"

	"github.com/matty-rose/gha-docs/pkg/formatter"
	"github.com/matty-rose/gha-docs/pkg/types"
)

func Write(action *types.CompositeAction) error {
	doc := formatter.NewMarkdownDocument()

	doc.WriteHeading(action.Name, 1)
	doc.WriteText(action.Description)
	doc.WriteNewLine()

	if len(action.Inputs) != 0 {
		var inputs [][]string
		for _, inp := range action.Inputs {
			inputs = append(inputs, []string{inp.Name, inp.Description, strconv.FormatBool(inp.Required), inp.Description})
		}

		doc.WriteNewLine()
		doc.WriteHeading("Inputs", 2)
		_, _ = doc.WriteTable(
			[]string{"Name", "Description", "Required", "Default"},
			inputs,
		)
	}

	if len(action.Outputs) != 0 {
		var outputs [][]string
		for _, out := range action.Outputs {
			outputs = append(outputs, []string{out.Name, out.Description, out.Value})
		}

		doc.WriteNewLine()
		doc.WriteHeading("Outputs", 2)
		_, _ = doc.WriteTable(
			[]string{"Name", "Description", "Value"},
			outputs,
		)
	}

	_, err := os.Stdout.Write([]byte(doc.Render() + "\n"))

	return err
}

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
package generator

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/matty-rose/gha-docs/pkg/document"
	"github.com/matty-rose/gha-docs/pkg/types"
)

type markdownGenerator struct {
	config Config
}

func (mdg markdownGenerator) Generate(action *types.CompositeAction) string {
	doc := document.NewMarkdownDocument()

	doc.WriteHeading(action.Name, 1)
	doc.WriteTextLn(action.Description)

	doc.WriteNewLine()
	doc.WriteHeading("Inputs", 2)

	if len(action.Inputs) != 0 {
		mdg.generateInputTable(action, doc)
	} else {
		doc.WriteTextLn("No inputs.")
	}

	doc.WriteNewLine()
	doc.WriteHeading("Outputs", 2)

	if len(action.Outputs) != 0 {
		mdg.generateOutputTable(action, doc)
	} else {
		doc.WriteTextLn("No outputs.")
	}

	doc.WriteNewLine()
	doc.WriteHeading("External Actions", 2)

	if len(action.Uses) != 0 {
		mdg.generateExternalActionTable(action, doc)
	} else {
		doc.WriteTextLn("No external actions.")
	}

	doc.WriteNewLine()
	doc.WriteHeading("Example Usage", 2)
	mdg.generateExampleUsageBlock(action, doc)

	return doc.Render()
}

func (mdg markdownGenerator) generateInputTable(act *types.CompositeAction, doc *document.MarkdownDocument) {
	columns := []string{"Name", "Description", "Required", "Default"}

	var rows [][]string

	sort.Slice(act.Inputs, func(a, b int) bool {
		return act.Inputs[a].Name < act.Inputs[b].Name
	})

	for _, inp := range act.Inputs {
		rows = append(
			rows,
			[]string{
				inp.Name,
				inp.Description,
				strconv.FormatBool(inp.Required),
				doc.FormatCode(inp.Default),
			},
		)
	}

	_, _ = doc.WriteTable(columns, rows)
}

func (mdg markdownGenerator) generateOutputTable(act *types.CompositeAction, doc *document.MarkdownDocument) {
	columns := []string{"Name", "Description", "Value"}

	var rows [][]string

	sort.Slice(act.Outputs, func(a, b int) bool {
		return act.Outputs[a].Name < act.Outputs[b].Name
	})

	for _, out := range act.Outputs {
		rows = append(rows, []string{out.Name, out.Description, doc.FormatCode(out.Value)})
	}

	_, _ = doc.WriteTable(columns, rows)
}

func (mdg markdownGenerator) generateExternalActionTable(act *types.CompositeAction, doc *document.MarkdownDocument) {
	columns := []string{"Name", "Creator", "Version", "Step Name", "Step ID"}

	var rows [][]string

	sort.Slice(act.Uses, func(a, b int) bool {
		return act.Uses[a].Name < act.Uses[b].Name
	})

	for _, act := range act.Uses {
		rows = append(
			rows,
			[]string{
				doc.CreateLink(act.Name, act.GetLink()),
				act.Creator,
				act.Version,
				act.StepName,
				act.StepID,
			},
		)
	}

	_, _ = doc.WriteTable(columns, rows)
}

func (mdg markdownGenerator) generateExampleUsageBlock(act *types.CompositeAction, doc *document.MarkdownDocument) {
	doc.WriteCodeBlockMarkerWithFormat("yaml")
	doc.WriteTextLn(fmt.Sprintf("- name: %s", act.Name))

	switch *mdg.config.ExampleUsageMode {
	case Remote:
		// TODO: Some way of getting actual owner/repo name here?
		doc.WriteTextLn("  uses: owner/repo@latest")
	case Local:
		doc.WriteTextLn("  uses: ./path/to/action.yml")
	}

	if len(act.Inputs) == 0 {
		doc.WriteCodeBlockMarker()
		return
	}

	doc.WriteTextLn("  with:")

	sort.Slice(act.Inputs, func(a, b int) bool {
		return act.Inputs[a].Name < act.Inputs[b].Name
	})

	for idx, inp := range act.Inputs {
		doc.WriteTextLn(fmt.Sprintf("    # %s", inp.Description))
		doc.WriteTextLn(fmt.Sprintf("    %s:", inp.Name))

		if idx != len(act.Inputs)-1 {
			doc.WriteNewLine()
		}
	}

	doc.WriteCodeBlockMarker()
}

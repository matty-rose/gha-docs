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
package generator_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matty-rose/gha-docs/pkg/generator"
	"github.com/matty-rose/gha-docs/pkg/types"
)

func newMarkdownConfig(mode generator.UsageMode) generator.Config {
	return generator.Config{Format: "markdown", ExampleUsageMode: &mode}
}

func TestGenerateMarkdownNameDescription(t *testing.T) {
	g, err := generator.New(newMarkdownConfig(generator.Remote))
	if err != nil {
		t.Fatal(err)
	}

	action := types.CompositeAction{Name: "test", Description: "also test"}

	expected := getMarkdownNameDesc()

	content := g.Generate(&action)

	assert.Equal(t, expected, content)
}

func TestGenerateMarkdownInputsRemote(t *testing.T) {
	g, err := generator.New(newMarkdownConfig(generator.Remote))
	if err != nil {
		t.Fatal(err)
	}

	action := types.CompositeAction{
		Name:        "test",
		Description: "also test",
		Inputs: []types.Input{
			{
				Name:        "a",
				Description: "a",
				Required:    false,
				Default:     "a",
			},
			{
				Name:        "b",
				Description: "b",
				Required:    true,
			},
		},
	}

	expected := getMarkdownInputs(generator.Remote)

	content := g.Generate(&action)

	assert.Equal(t, expected, content)
}

func TestGenerateMarkdownInputsLocal(t *testing.T) {
	g, err := generator.New(newMarkdownConfig(generator.Local))
	if err != nil {
		t.Fatal(err)
	}

	action := types.CompositeAction{
		Name:        "test",
		Description: "also test",
		Inputs: []types.Input{
			{
				Name:        "a",
				Description: "a",
				Required:    false,
				Default:     "a",
			},
			{
				Name:        "b",
				Description: "b",
				Required:    true,
			},
		},
	}

	expected := getMarkdownInputs(generator.Local)

	content := g.Generate(&action)

	assert.Equal(t, expected, content)
}

func TestGenerateMarkdownOutputs(t *testing.T) {
	g, err := generator.New(newMarkdownConfig(generator.Remote))
	if err != nil {
		t.Fatal(err)
	}

	action := types.CompositeAction{
		Name:        "test",
		Description: "also test",
		Outputs: []types.Output{
			{
				Name:        "a",
				Description: "a",
				Value:       "a",
			},
			{
				Name:        "b",
				Description: "b",
				Value:       "b",
			},
		},
	}

	expected := getMarkdownOutputs()

	content := g.Generate(&action)

	assert.Equal(t, expected, content)
}

func TestGenerateMarkdownExternal(t *testing.T) {
	g, err := generator.New(newMarkdownConfig(generator.Remote))
	if err != nil {
		t.Fatal(err)
	}

	action := types.CompositeAction{
		Name:        "test",
		Description: "also test",
		Uses: []types.ExternalAction{
			{
				Creator: "actions",
				Name:    "cache",
				Version: "v2.1.6",
			},
			{
				Creator:  "actions",
				Name:     "setup-python",
				Version:  "v2",
				StepName: "Set up python",
			},
		},
	}

	expected := getMarkdownExternal()

	content := g.Generate(&action)

	assert.Equal(t, expected, content)
}

func TestGenerateMarkdownFull(t *testing.T) {
	g, err := generator.New(newMarkdownConfig(generator.Remote))
	if err != nil {
		t.Fatal(err)
	}

	action := types.CompositeAction{
		Name:        "test",
		Description: "also test",
		Inputs: []types.Input{
			{
				Name:        "a",
				Description: "a",
				Required:    false,
				Default:     "a",
			},
			{
				Name:        "b",
				Description: "b",
				Required:    true,
			},
		},
		Outputs: []types.Output{
			{
				Name:        "a",
				Description: "a",
				Value:       "a",
			},
			{
				Name:        "b",
				Description: "b",
				Value:       "b",
			},
		},
	}

	expected := getMarkdownFull()

	content := g.Generate(&action)

	assert.Equal(t, expected, content)
}

func getUsageModeOutputString(mode generator.UsageMode) string {
	switch mode {
	case generator.Remote:
		return "owner/repo@latest"
	case generator.Local:
		return "./path/to/action.yml"
	default:
		return "owner/repo@latest"
	}
}

func getMarkdownNameDesc() string {
	return `# test
also test

## Inputs
No inputs.

## Outputs
No outputs.

## External Actions
No external actions.

## Example Usage
` + "```yaml" + `
- name: test
  uses: owner/repo@latest
` + "```" + `
`
}

func getMarkdownInputs(mode generator.UsageMode) string {
	return fmt.Sprintf(`# test
also test

## Inputs
| Name | Description | Required | Default |
| --- | --- | --- | --- |
| a | a | false | `+"`a`"+` |
| b | b | true |  |

## Outputs
No outputs.

## External Actions
No external actions.

## Example Usage
`+"```yaml"+`
- name: test
  uses: %s
  with:
    # a
    a:

    # b
    b:
`+"```"+`
`, getUsageModeOutputString(mode))
}

func getMarkdownOutputs() string {
	return `# test
also test

## Inputs
No inputs.

## Outputs
| Name | Description | Value |
| --- | --- | --- |
| a | a | ` + "`a`" + ` |
| b | b | ` + "`b`" + ` |

## External Actions
No external actions.

## Example Usage
` + "```yaml" + `
- name: test
  uses: owner/repo@latest
` + "```" + `
`
}

func getMarkdownExternal() string {
	return `# test
also test

## Inputs
No inputs.

## Outputs
No outputs.

## External Actions
| Name | Creator | Version | Step Name | Step ID |
| --- | --- | --- | --- | --- |
| [cache](https://github.com/actions/cache/tree/v2.1.6) | actions | v2.1.6 |  |  |
| [setup-python](https://github.com/actions/setup-python/tree/v2) | actions | v2 | Set up python |  |

## Example Usage
` + "```yaml" + `
- name: test
  uses: owner/repo@latest
` + "```" + `
`
}

func getMarkdownFull() string {
	return `# test
also test

## Inputs
| Name | Description | Required | Default |
| --- | --- | --- | --- |
| a | a | false | ` + "`a`" + ` |
| b | b | true |  |

## Outputs
| Name | Description | Value |
| --- | --- | --- |
| a | a | ` + "`a`" + ` |
| b | b | ` + "`b`" + ` |

## External Actions
No external actions.

## Example Usage
` + "```yaml" + `
- name: test
  uses: owner/repo@latest
  with:
    # a
    a:

    # b
    b:
` + "```" + `
`
}

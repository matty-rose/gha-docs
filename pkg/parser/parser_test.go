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
package parser_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matty-rose/gha-docs/pkg/parser"
)

func TestParseNameDescription(t *testing.T) {
	t.Parallel()

	action, err := parser.Parse("./testdata/name_description.yaml")
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "test", action.Name)
	assert.Equal(t, "test", action.Description)
}

func TestParseInputs(t *testing.T) {
	t.Parallel()

	action, err := parser.Parse("./testdata/inputs.yaml")
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, action.Inputs, 2)
}

func TestParseOutputs(t *testing.T) {
	t.Parallel()

	action, err := parser.Parse("./testdata/outputs.yaml")
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, action.Outputs, 1)
}

func TestParseUses(t *testing.T) {
	t.Parallel()

	action, err := parser.Parse("./testdata/uses.yaml")
	if err != nil {
		t.Fatal(err)
	}

	assert.Len(t, action.Uses, 4)
}

func TestInvalidFiles(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		file           string
		expectedErrMsg string
	}{
		{"./testdata/doesnt_exist.yaml", "couldn't read given yaml file"},
		{"./testdata/invalid.yaml", "failed unmarshalling yaml data"},
		{"./testdata/invalid_inputs.yaml", "failed parsing action input into struct"},
		{"./testdata/invalid_outputs.yaml", "failed parsing action output into struct"},
		{"./testdata/invalid_uses.yaml", "step does not have a valid structure"},
	}

	for _, tc := range testCases {
		action, err := parser.Parse(tc.file)

		assert.Nil(t, action)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), tc.expectedErrMsg)
	}
}

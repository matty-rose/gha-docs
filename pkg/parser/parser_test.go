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
	assertion := assert.New(t)

	action, err := parser.Parse("./testdata/name_description.yaml")
	if err != nil {
		t.Fatal(err)
	}

	assertion.Equal("test", action.Name)
	assertion.Equal("test", action.Description)
}

func TestParseInputs(t *testing.T) {
	assertion := assert.New(t)

	action, err := parser.Parse("./testdata/inputs.yaml")
	if err != nil {
		t.Fatal(err)
	}

	assertion.Len(action.Inputs, 2)
}

func TestParseOutputs(t *testing.T) {
	assertion := assert.New(t)

	action, err := parser.Parse("./testdata/outputs.yaml")
	if err != nil {
		t.Fatal(err)
	}

	assertion.Len(action.Outputs, 1)
}

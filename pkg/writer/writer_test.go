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
package writer_test

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/matty-rose/gha-docs/pkg/writer"
)

func TestStdoutWriter(t *testing.T) {
	t.Parallel()

	testCases := []string{"word", "some more words", `multiline
    shenanigans`}

	for _, tc := range testCases {
		old := os.Stdout // keep backup of the real stdout

		r, w, err := os.Pipe()
		if err != nil {
			t.Fatal(err)
		}

		os.Stdout = w

		outC := make(chan string)
		// copy the output in a separate goroutine so printing can't block indefinitely
		go func() {
			var buf bytes.Buffer

			io.Copy(&buf, r)
			outC <- buf.String()
		}()

		writer.Write(writer.WriteInputs{Content: tc, OutputFile: "", Inject: false})

		// back to normal state
		w.Close()

		os.Stdout = old // restoring the real stdout
		out := <-outC
		assert.Equal(t, tc, out)
	}
}

func TestFileWriter(t *testing.T) {
	t.Parallel()

	testFile := "test.txt"
	defer os.Remove(testFile)

	testCases := []string{"word", "some more words", `multiline
    shenanigans`}

	for _, tc := range testCases {
		writer.Write(writer.WriteInputs{Content: tc, OutputFile: testFile, Inject: false})

		got, err := os.ReadFile(testFile)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc, string(got))
	}
}

func TestFileWriterInject(t *testing.T) {
	t.Parallel()

	content := "dummy\n"

	testCases := []struct {
		outputFile      string
		expectedContent string
	}{
		{"./testdata/non_existent.md", "dummy\n"},
		{"./testdata/empty.md", "dummy\n"},
		{"./testdata/only_markers.md", fmt.Sprintf("%s\ndummy\n%s\n", writer.BeginInjection, writer.EndInjection)},
		{"./testdata/existing_and_markers.md",
			fmt.Sprintf(
				"something already existing\n%s\ndummy\n%s\nsomething else not overwritten\n",
				writer.BeginInjection,
				writer.EndInjection,
			),
		},
	}

	for _, tc := range testCases {
		writer.Write(writer.WriteInputs{Content: content, OutputFile: tc.outputFile, Inject: true})

		got, err := os.ReadFile(tc.outputFile)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc.expectedContent, string(got))
	}
}

func TestFileWriterInjectInvalid(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		outputFile     string
		expectedErrMsg string
	}{
		{"./testdata/no_begin_marker.md", "missing begin injection marker"},
		{"./testdata/no_end_marker.md", "missing end injection marker"},
		{"./testdata/end_before_begin.md", "end injection marker is before begin"},
	}

	for _, tc := range testCases {
		err := writer.Write(writer.WriteInputs{Content: "dummy", OutputFile: tc.outputFile, Inject: true})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), tc.expectedErrMsg)
	}
}

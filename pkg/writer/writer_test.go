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

		writer.Write(tc, "")

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
		writer.Write(tc, testFile)

		got, err := os.ReadFile(testFile)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, tc, string(got))
	}
}

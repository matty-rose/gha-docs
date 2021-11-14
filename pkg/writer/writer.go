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
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

type stdoutWriter struct{}

func (sw stdoutWriter) Write(content []byte) (int, error) {
	return os.Stdout.Write(content)
}

type fileWriter struct {
	file string
}

func (fw fileWriter) Write(content []byte) (n int, err error) {
	f, err := os.OpenFile(fw.file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("couldn't open file: %s", fw.file))
	}

	defer func() {
		cerr := f.Close()
		if err == nil {
			err = cerr
		}
	}()

	n, err = f.Write(content)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("couldn't write to file: %s", fw.file))
	}

	// Uses named return values
	return
}

func Write(content string, outputFile string) error {
	var w io.Writer

	if outputFile != "" {
		w = fileWriter{outputFile}
	} else {
		w = stdoutWriter{}
	}

	_, err := io.WriteString(w, content)

	return err
}

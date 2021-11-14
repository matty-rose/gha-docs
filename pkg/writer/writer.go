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
	"strings"

	"github.com/pkg/errors"
)

const (
	BeginInjection string = "<!-- BEGIN GHA DOCS -->"
	EndInjection   string = "<!-- END GHA DOCS -->"
)

type stdoutWriter struct{}

func (sw stdoutWriter) Write(content []byte) (int, error) {
	return os.Stdout.Write(content)
}

type fileWriter struct {
	file   string
	inject bool
}

func (fw fileWriter) Write(content []byte) (int, error) {
	if !fw.inject {
		return fw.writeFile(content)
	}

	existingFileContent, err := os.ReadFile(fw.file)
	if err != nil || len(existingFileContent) == 0 {
		// Even if inject flag is passed, if file doesn't exist OR is empty, then write as per normal.
		return fw.writeFile(content)
	}

	return fw.injectContent(string(existingFileContent), string(content))
}

func (fw fileWriter) injectContent(existing, newContent string) (int, error) {
	beginIdx := strings.Index(existing, BeginInjection)
	endIdx := strings.Index(existing, EndInjection)

	if beginIdx == -1 {
		return 0, errors.New(fmt.Sprintf("missing begin injection marker: %s", BeginInjection))
	}

	if endIdx == -1 {
		return 0, errors.New(fmt.Sprintf("missing end injection marker: %s", EndInjection))
	}

	if endIdx < beginIdx {
		return 0, errors.New("end injection marker is before begin injection marker")
	}

	injectedContent := existing[:beginIdx+len(BeginInjection)] + "\n" + newContent + existing[endIdx:]

	return fw.writeFile([]byte(injectedContent))
}

func (fw fileWriter) writeFile(content []byte) (int, error) {
	err := os.WriteFile(fw.file, []byte(content), 0644)
	if err != nil {
		return 0, errors.Wrap(err, fmt.Sprintf("couldn't write to file: %s", fw.file))
	}

	return len(content), err
}

type WriteInputs struct {
	Content    string
	OutputFile string
	Inject     bool
}

func Write(inputs WriteInputs) error {
	var w io.Writer

	if inputs.OutputFile != "" {
		w = fileWriter{inputs.OutputFile, inputs.Inject}
	} else {
		w = stdoutWriter{}
	}

	_, err := io.WriteString(w, inputs.Content)

	return err
}

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
package cmd

import (
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/matty-rose/gha-docs/pkg/generator"
	"github.com/matty-rose/gha-docs/pkg/parser"
)

// Format flag
var format string

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate [PATH]",
	Short: "Generate documentation for a composite GitHub action.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		action, err := parser.Parse(args[0])
		if err != nil {
			return errors.Wrap(err, "couldn't parse the action file")
		}

		var g generator.Generator
		g, err = generator.New(format)
		if err != nil {
			return errors.Wrap(err, "couldn't construct the generator")
		}

		out := g.Generate(action)
		_, err = os.Stdout.Write([]byte(out))
		return err
	},
}

func init() {
	generateCmd.PersistentFlags().StringVarP(
		&format,
		"format",
		"f",
		"markdown",
		"Format to generate documentation in - currently only markdown is supported.",
	)
	rootCmd.AddCommand(generateCmd)
}

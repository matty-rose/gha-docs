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
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/thediveo/enumflag"

	"github.com/matty-rose/gha-docs/pkg/generator"
	"github.com/matty-rose/gha-docs/pkg/parser"
	"github.com/matty-rose/gha-docs/pkg/writer"
)

// Format flag
var format string

// Output file flag
var outputFile string

// Inject flag
var inject bool

// Usage mode flag
var usageMode generator.UsageMode = generator.Remote

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
		g, err = generator.New(generator.Config{
			Format:           format,
			ExampleUsageMode: &usageMode,
		})
		if err != nil {
			return errors.Wrap(err, "couldn't construct the generator")
		}

		content := g.Generate(action)

		err = writer.Write(writer.WriteInputs{
			Content:    content,
			OutputFile: outputFile,
			Inject:     inject,
		})
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
	generateCmd.PersistentFlags().StringVarP(
		&outputFile,
		"output-file",
		"o",
		"",
		"File to write generated documentation to.",
	)
	generateCmd.PersistentFlags().BoolVarP(
		&inject,
		"inject",
		"i",
		false,
		"Set flag to inject generated documentation between markers. Ignored if not writing to a file. Defaults to false.",
	)
	generateCmd.PersistentFlags().VarP(
		enumflag.New(&usageMode, "mode", generator.UsageModeIDs, enumflag.EnumCaseInsensitive),
		"usage-mode",
		"u",
		"Sets the usage mode when generating example usage block. Must be one of 'remote' or 'local'.",
	)
	rootCmd.AddCommand(generateCmd)
}

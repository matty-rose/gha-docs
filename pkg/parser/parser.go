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
package parser

import (
	"fmt"
	"io/ioutil"

	"github.com/matty-rose/gha-docs/pkg/types"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
)

func Parse(filename string) (*types.CompositeAction, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't read given yaml file")
	}

	data := make(map[interface{}]interface{})

	err = yaml.Unmarshal(file, &data)
	if err != nil {
		return nil, errors.Wrap(err, "failed unmarshalling yaml data")
	}

	var action types.CompositeAction

	parseMetadata(&action, data)

	err = parseInputs(&action, data)
	if err != nil {
		return nil, err
	}

	err = parseOutputs(&action, data)
	if err != nil {
		return nil, err
	}

	return &action, nil
}

func parseMetadata(action *types.CompositeAction, data map[interface{}]interface{}) {
	action.SetName(data["name"].(string))
	action.SetDescription(data["description"].(string))
}

func parseInputs(action *types.CompositeAction, data map[interface{}]interface{}) error {
	inputs, ok := data["inputs"].(map[string]interface{})
	if !ok {
		// TODO: Replace with logrus
		fmt.Println("no inputs found")
	}

	for name, input := range inputs {
		inp := types.Input{Name: name}

		err := mapstructure.Decode(input, &inp)
		if err != nil {
			return errors.Wrap(err, "failed parsing action input into struct")
		}

		action.AddInput(inp)
	}

	return nil
}

func parseOutputs(action *types.CompositeAction, data map[interface{}]interface{}) error {
	outputs, ok := data["outputs"].(map[string]interface{})
	if !ok {
		// TODO: Replace with logrus
		fmt.Println("no outputs found")
	}

	for name, output := range outputs {
		out := types.Output{Name: name}

		err := mapstructure.Decode(output, &out)
		if err != nil {
			return errors.Wrap(err, "failed parsing action output into struct")
		}

		action.AddOutput(out)
	}

	return nil
}

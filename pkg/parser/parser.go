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
	"io/ioutil"
	"path/filepath"
	"regexp"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/matty-rose/gha-docs/pkg/types"
)

func Parse(filename string) (*types.CompositeAction, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't read given yaml file")
	}

	data := make(map[interface{}]interface{})

	if err := yaml.Unmarshal(file, &data); err != nil {
		return nil, errors.Wrap(err, "failed unmarshalling yaml data")
	}

	var action types.CompositeAction

	parseMetadata(&action, data)

	if err := parseInputs(&action, data); err != nil {
		return nil, err
	}

	if err := parseOutputs(&action, data); err != nil {
		return nil, err
	}

	if err := parseExternalActions(&action, data); err != nil {
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
		logrus.Debug("no inputs found")
	}

	for name, input := range inputs {
		inp := types.Input{Name: name}

		if err := mapstructure.Decode(input, &inp); err != nil {
			return errors.Wrap(err, "failed parsing action input into struct")
		}

		action.AddInput(inp)
	}

	return nil
}

func parseOutputs(action *types.CompositeAction, data map[interface{}]interface{}) error {
	outputs, ok := data["outputs"].(map[string]interface{})
	if !ok {
		logrus.Debug("no outputs found")
	}

	for name, output := range outputs {
		out := types.Output{Name: name}

		if err := mapstructure.Decode(output, &out); err != nil {
			return errors.Wrap(err, "failed parsing action output into struct")
		}

		action.AddOutput(out)
	}

	return nil
}

func tryMatchRemoteUses(text string) ([][]string, bool) {
	regex := *regexp.MustCompile(`(.+)\/(.+)@(.+)`)

	res := regex.FindAllStringSubmatch(text, -1)
	if res != nil {
		return res, true
	}

	return res, false
}

func parseUses(ext *types.ExternalAction, uses string) error {
	if remote, ok := tryMatchRemoteUses(uses); ok {
		ext.Creator = remote[0][1]
		ext.Name = remote[0][2]
		ext.Version = remote[0][3]
		ext.Local = false

		return nil
	}

	// If remote regex doesn't match, assume its a local reference to an action
	logrus.Debug("Matching local action reference")

	ext.Local = true
	ext.LocalPath = &uses
	ext.Name = filepath.Base(uses)

	return nil
}

func parseExternalActions(action *types.CompositeAction, data map[interface{}]interface{}) error {
	runs, ok := data["runs"].(map[string]interface{})
	if !ok {
		logrus.Debug("no runs found")
	}

	steps, ok := runs["steps"].([]interface{})
	if !ok {
		logrus.Debug("no steps found")
	}

	for _, s := range steps {
		step, ok := s.(map[string]interface{})
		if !ok {
			return errors.New("step does not have a valid structure")
		}

		if _, ok := step["uses"].(string); !ok {
			logrus.Debug("step uses key does not exist, or isn't a string, skipping")
			continue
		}

		var ext types.ExternalAction

		if stepName, ok := step["name"].(string); ok {
			ext.StepName = stepName
		}

		if stepID, ok := step["id"].(string); ok {
			ext.StepID = stepID
		}

		err := parseUses(&ext, step["uses"].(string))
		if err != nil {
			return errors.Wrap(err, "couldn't parse the value in the 'uses' field")
		}

		logrus.Debug(ext)

		action.AddExternalAction(ext)
	}

	return nil
}

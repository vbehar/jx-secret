package schema

import (
	"fmt"
	"io/ioutil"

	"github.com/jenkins-x/jx-helpers/pkg/files"
	"github.com/pkg/errors"

	"gopkg.in/yaml.v1"
)

// LoadSurveySchema loads a specific secret mapping YAML file
func LoadSurveySchema(fileName string) (*SurveySchema, error) {
	config := &SurveySchema{}

	exists, err := files.FileExists(fileName)
	if err != nil {
		return config, errors.Wrapf(err, "failed to check file exists %s", fileName)
	}
	if !exists {
		return config, nil
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to load file %s due to %s", fileName, err)
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML file %s due to %s", fileName, err)
	}

	err = config.Validate()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to validate secret mapping YAML file %s", fileName)
	}
	return config, nil
}

package dotgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

// TODO: Custom Marshalling according to Plugin Type.
type Config struct {
	Data []map[string]interface{}
}

// ReadConfig returns the read config data from the file.
func ReadConfig(path string) (map[string]interface{}, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})

	switch ext := filepath.Ext(path); ext {
	case ".json":
		err = json.Unmarshal(file, &data)
	case ".yaml", ".yml":
		err = yaml.Unmarshal(file, &data)
	default:
		return nil, fmt.Errorf("extension %s is unsupported", ext)
	}

	if err != nil {
		return nil, fmt.Errorf("err while parsing config: %w", err)
	}

	return data, nil
}

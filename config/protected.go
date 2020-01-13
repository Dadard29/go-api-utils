package config

import (
	"encoding/json"
	"io/ioutil"
)

// 1 level
func newEnv() *env {
	return &env{config: map[string]string{}}
}

// 3 levels
func newFile(path string) (*file, error) {
	f := &file{config: map[string]map[string]map[string]string{}}
	err := f.loadFile(path)
	return f, err
}

func (f *file) loadFile(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &f.config)
	return err
}
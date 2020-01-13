package config

import (
	"errors"
	"fmt"
	"os"
)

func (a *Accessor) ReloadConfigFile() {
	var err error
	a.file, err = newFile(a.configFilePath)
	a.logger.CheckErr(err)
}

func (a *Accessor) GetValueFromFile(category string, subcategory string, key string) (string, error) {
	if value, check := a.file.config[category][subcategory][key]; check {
		a.logger.Debug(fmt.Sprintf("retrieved %s.%s.%s property: %s", category, subcategory, key, value))
		return value, nil
	} else {
		msg := fmt.Sprintf("unknow property: %s.%s.%s", category, subcategory, key)
		a.logger.Error(msg)
		return "", errors.New(msg)
	}
}

func (a *Accessor) GetSubcategoryFromFile(category string, subcategory string) (map[string]string, error) {
	if value, check := a.file.config[category][subcategory]; check {
		a.logger.Debug(fmt.Sprintf("retrieved %s.%s subcategory", category, subcategory))
		return value, nil
	} else {
		msg := fmt.Sprintf("unknow subcategory: %s.%s", category, subcategory)
		a.logger.Error(msg)
		return map[string]string{}, errors.New(msg)
	}
}

func (a *Accessor) GetCategoryFromFile(category string) (map[string]map[string]string, error) {
	if value, check := a.file.config[category]; check {
		a.logger.Debug(fmt.Sprintf("retrieved %s category", category))
		return value, nil
	} else {
		msg := fmt.Sprintf("unknow category: %s", category)
		a.logger.Error(msg)
		return map[string]map[string]string{}, errors.New(msg)
	}
}

func (a *Accessor) GetEnv(name string) string {
	if value, check := a.env.config[name]; check {
		return value
	} else {
		value := os.Getenv(name)
		a.env.config[name] = value
		return value
	}
}

/*
 * Author:     Vignesh Kumar
 * Copyright:  vigneshuvi
 * Date:       01/07/2019
 * This file contains the Main.
 */
package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

import Model "./model"

type configuration struct {
	Env Model.ConfigModel
}

var instance *configuration
var once sync.Once
var envConfigFile = "./config/environments/"

// Help to set Config.
func SetConfig(env string) *configuration {
	once.Do(func() {
		instance = &configuration{getConfig(env)}
	})
	return instance
}

// Get the config.
func getConfig(env string) Model.ConfigModel {
	// Generate the config url
	var confFile = fmt.Sprintf("%s%s.json", envConfigFile, env)
	envConfig, err := loadConfiguration(confFile)
	if err != nil {
		fmt.Println(err)
	}

	return envConfig
}

// Load the config file
func loadConfiguration(file string) (Model.ConfigModel, error) {
	// we initialize our Configs environment
	var env Model.ConfigModel

	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		return env, err
	}

	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&env)
	return env, err
}

package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// ServerConfig stores appconfig data
type ServerConfig struct {
	Port           uint16
	Services       map[string]string
	DefaultService string `json:"default_service"`
}

// NewServerConfig returns ServerConfig file with provided json file content
func NewServerConfig(path string) (ServerConfig, error) {
	var result ServerConfig

	raw, err := ioutil.ReadFile(path)
	if err != nil {
		errstr := fmt.Sprintf("Error while reading config file: %v", err)
		return result, errors.New(errstr)
	}
	json.Unmarshal(raw, &result)

	return result, nil
}

// GetPortString returns port number as a port string, e.g. ":8080"
func (sc *ServerConfig) GetPortString() string {
	portstr := fmt.Sprintf(":%v", sc.Port)
	return portstr
}

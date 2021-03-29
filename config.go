package linkedin

import (
	"fmt"
	"github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	Linkedin *LinkedinConfig `json:"linkedin"`
}

// LinkedinConfig ...
type LinkedinConfig struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
	Authorization struct {
		Id     string `json:"id"`
		Secret string `json:"secret"`
	} `json:"authorization"`
	Hosts struct {
		Api string `json:"api"`
	} `json:"hosts"`
}

// NewConfig ...
func NewConfig() (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	return appConfig, simpleConfig, err
}

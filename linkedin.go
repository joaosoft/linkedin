package linkedin

import (
	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

type Linkedin struct {
	client        manager.IGateway
	config        *LinkedinConfig
	pm            *manager.Manager
	logger        logger.ILogger
	isLogExternal bool
}

// NewLinkedin ...
func NewLinkedin(options ...LinkedinOption) (*Linkedin, error) {
	config, simpleConfig, err := NewConfig()
	pm := manager.NewManager(manager.WithRunInBackground(false))

	client, err := pm.NewSimpleGateway()
	if err != nil {
		return nil, err
	}

	service := &Linkedin{
		client: client,
		pm:     pm,
		config: config.Linkedin,
		logger: logger.NewLogDefault("linkedin", logger.WarnLevel),
	}

	if service.isLogExternal {
		service.pm.Reconfigure(manager.WithLogger(service.logger))
	}

	if err != nil {
		service.logger.Error(err.Error())
	} else if config.Linkedin != nil {
		service.pm.AddConfig("config_app", simpleConfig)
		level, _ := logger.ParseLevel(config.Linkedin.Log.Level)
		service.logger.Debugf("setting log level to %s", level)
		service.logger.Reconfigure(logger.WithLevel(level))
	}

	service.Reconfigure(options...)

	return service, nil
}
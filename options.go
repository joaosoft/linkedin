package linkedin

import (
	"github.com/joaosoft/logger"
	"github.com/joaosoft/manager"
)

// LinkedinOption ...
type LinkedinOption func(linkedin *Linkedin)

// Reconfigure ...
func (linkedin *Linkedin) Reconfigure(options ...LinkedinOption) {
	for _, option := range options {
		option(linkedin)
	}
}

// WithConfiguration ...
func WithConfiguration(config *LinkedinConfig) LinkedinOption {
	return func(linkedin *Linkedin) {
		linkedin.config = config
	}
}

// WithLogger ...
func WithLogger(logger logger.ILogger) LinkedinOption {
	return func(linkedin *Linkedin) {
		linkedin.logger = logger
		linkedin.isLogExternal = true
	}
}

// WithLogLevel ...
func WithLogLevel(level logger.Level) LinkedinOption {
	return func(linkedin *Linkedin) {
		linkedin.logger.SetLevel(level)
	}
}

// WithManager ...
func WithManager(mgr *manager.Manager) LinkedinOption {
	return func(linkedin *Linkedin) {
		linkedin.pm = mgr
	}
}

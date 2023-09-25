package config

import (
	"sync"
)

type (
	Bootstrap struct {
		// flag ...
		flag string

		// config ...
	}

	BootstrapOption func(*Bootstrap)
)

var (
	c    *Bootstrap
	once sync.Once
)

func NewBootstrap(opts ...BootstrapOption) *Bootstrap {
	if c != nil {
		return c
	}

	b := &Bootstrap{
		flag: "config/config.yaml",
	}

	for _, o := range opts {
		o(b)
	}

	once.Do(func() {
		c = b
	})

	return b
}

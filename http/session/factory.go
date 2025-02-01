package session

import (
	ntHttp "github.com/go-nt/nt/http"
)

func NewSession(ctx *ntHttp.Context) *Driver {

	var d Driver

	switch config.driver {
	case "file":
		d = new(DriverFile)
	case "redis":
		d = new(DriverRedis)
	}

	d.Init(config, ctx)

	return &d
}

//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package backend

import (
	"github.com/fengshux/blog2/backend/conf"
	"github.com/fengshux/blog2/backend/controller"
	"github.com/fengshux/blog2/backend/db"
	"github.com/fengshux/blog2/backend/service"
	"github.com/google/wire"
)

func NewRegister() (*Register, error) {
	panic(wire.Build(
		conf.ProviderSet,
		db.ProviderSet,
		service.ProviderSet,
		controller.ProviderSet,
		newRegister,
	))
}

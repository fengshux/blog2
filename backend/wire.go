//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package backend

import (
	
	"github.com/google/wire"
	"github.com/fengshux/blog2/backend/controller"
	"github.com/fengshux/blog2/backend/db"
	"github.com/fengshux/blog2/backend/service"
)

func NewRegister() (*Register, error) {
	panic(wire.Build(
		db.ProviderSet,
		service.ProviderSet,
		controller.ProviderSet,
		newRegister,
	))
}

// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package backend

import (
	"github.com/fengshux/blog2/backend/controller"
	"github.com/fengshux/blog2/backend/db"
	"github.com/fengshux/blog2/backend/service"
)

// Injectors from wire.go:

func NewRegister() (*Register, error) {
	gormDB := db.NewPG()
	user := service.NewUser(gormDB)
	controllerUser := controller.NewUser(user)
	register := newRegister(controllerUser)
	return register, nil
}

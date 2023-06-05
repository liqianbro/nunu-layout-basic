// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/go-nunu/nunu-layout-base/internal/dao"
	"github.com/go-nunu/nunu-layout-base/internal/handler"
	"github.com/go-nunu/nunu-layout-base/internal/server"
	"github.com/go-nunu/nunu-layout-base/internal/service"
	"github.com/go-nunu/nunu-layout-base/pkg/log"
	"github.com/spf13/viper"
)

// Injectors from wire.go:

// wire.go 初始化模块
func NewApp(viperViper *viper.Viper, logger *log.Logger) (*gin.Engine, func(), error) {
	handlerHandler := handler.NewHandler(logger)
	serviceService := service.NewService(logger)
	daoDao := dao.NewDao(logger)
	userDao := dao.NewUserDao(daoDao)
	userService := service.NewUserService(serviceService, userDao)
	userHandler := handler.NewUserHandler(handlerHandler, userService)
	engine, cleanup := server.NewServerHTTP(logger, userHandler)
	return engine, func() {
		cleanup()
	}, nil
}

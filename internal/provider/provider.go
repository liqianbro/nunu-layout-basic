package provider

import (
	"github.com/go-nunu/nunu-layout-base/internal/dao"
	"github.com/go-nunu/nunu-layout-base/internal/handler"
	"github.com/go-nunu/nunu-layout-base/internal/server"
	"github.com/go-nunu/nunu-layout-base/internal/service"
	"github.com/google/wire"
)

var ServerSet = wire.NewSet(server.NewServerHTTP)

var DaoSet = wire.NewSet(
	dao.NewDb,
	dao.NewDao,
	dao.NewUserDao,
)

var ServiceSet = wire.NewSet(
	service.NewService,
	service.NewUserService,
)

var HandlerSet = wire.NewSet(
	handler.NewHandler,
	handler.NewUserHandler,
)

// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package demo

import (
	"github.com/aide-family/moon/cmd/server/demo/internal/biz"
	"github.com/aide-family/moon/cmd/server/demo/internal/data"
	"github.com/aide-family/moon/cmd/server/demo/internal/data/repoimpl"
	"github.com/aide-family/moon/cmd/server/demo/internal/democonf"
	"github.com/aide-family/moon/cmd/server/demo/internal/server"
	"github.com/aide-family/moon/cmd/server/demo/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *democonf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	grpcServer := server.NewGRPCServer(bootstrap)
	httpServer := server.NewHTTPServer(bootstrap)
	dataData, cleanup, err := data.NewData(bootstrap)
	if err != nil {
		return nil, nil, err
	}
	greeterRepo := data.NewGreeterRepo(dataData)
	cacheRepo := repoimpl.NewCacheRepo(dataData)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo, cacheRepo)
	greeterService := service.NewGreeterService(greeterUsecase)
	serverServer := server.RegisterService(grpcServer, httpServer, greeterService)
	app := newApp(serverServer, logger)
	return app, func() {
		cleanup()
	}, nil
}

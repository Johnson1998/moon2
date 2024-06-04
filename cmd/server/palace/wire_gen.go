// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package palace

import (
	"github.com/aide-family/moon/cmd/server/palace/internal/biz"
	"github.com/aide-family/moon/cmd/server/palace/internal/data"
	"github.com/aide-family/moon/cmd/server/palace/internal/data/microserver"
	"github.com/aide-family/moon/cmd/server/palace/internal/data/microserver/microserverrepoimpl"
	"github.com/aide-family/moon/cmd/server/palace/internal/data/repoimpl"
	"github.com/aide-family/moon/cmd/server/palace/internal/palaceconf"
	"github.com/aide-family/moon/cmd/server/palace/internal/server"
	"github.com/aide-family/moon/cmd/server/palace/internal/service"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/authorization"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/datasource"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/resource"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/team"
	"github.com/aide-family/moon/cmd/server/palace/internal/service/user"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(bootstrap *palaceconf.Bootstrap, logger log.Logger) (*kratos.App, func(), error) {
	grpcServer := server.NewGRPCServer(bootstrap)
	dataData, cleanup, err := data.NewData(bootstrap)
	if err != nil {
		return nil, nil, err
	}
	captcha := repoimpl.NewCaptchaRepository(dataData)
	captchaBiz := biz.NewCaptchaBiz(captcha)
	repositoryUser := repoimpl.NewUserRepository(dataData)
	repositoryTeam := repoimpl.NewTeamRepository(dataData)
	cache := repoimpl.NewCacheRepository(dataData)
	teamRole := repoimpl.NewTeamRoleRepository(dataData)
	authorizationBiz := biz.NewAuthorizationBiz(repositoryUser, repositoryTeam, cache, teamRole)
	authorizationService := authorization.NewAuthorizationService(captchaBiz, authorizationBiz)
	httpServer := server.NewHTTPServer(bootstrap, authorizationService)
	greeterRepo := data.NewGreeterRepo(dataData)
	greeterUsecase := biz.NewGreeterUsecase(greeterRepo)
	greeterService := service.NewGreeterService(greeterUsecase)
	healthService := service.NewHealthService()
	userBiz := biz.NewUserBiz(repositoryUser)
	userService := user.NewUserService(userBiz)
	repositoryResource := repoimpl.NewResourceRepository(dataData)
	resourceBiz := biz.NewResourceBiz(repositoryResource)
	resourceService := resource.NewResourceService(resourceBiz)
	teamBiz := biz.NewTeamBiz(repositoryTeam)
	teamService := team.NewTeamService(teamBiz)
	teamRoleBiz := biz.NewTeamRoleBiz(teamRole)
	roleService := team.NewRoleService(teamRoleBiz)
	repositoryDatasource := repoimpl.NewDatasourceRepository(dataData)
	datasourceMetric := repoimpl.NewDatasourceMetricRepository(dataData)
	houYiConn, cleanup2, err := microserver.NewHouYiConn(bootstrap)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	microrepositoryDatasourceMetric := microserverrepoimpl.NewDatasourceMetricRepository(houYiConn)
	lock := repoimpl.NewLockRepository(dataData)
	datasourceBiz := biz.NewDatasourceBiz(repositoryDatasource, datasourceMetric, microrepositoryDatasourceMetric, lock)
	datasourceService := datasource.NewDatasourceService(datasourceBiz)
	teamMenu := repoimpl.NewTeamMenuRepository(dataData)
	rabbitConn, cleanup3, err := microserver.NewRabbitRpcConn(bootstrap)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	msg := microserverrepoimpl.NewMsgRepository(rabbitConn)
	menuBiz := biz.NewMenuBiz(teamMenu, msg)
	menuService := resource.NewMenuService(menuBiz)
	serverServer := server.RegisterService(grpcServer, httpServer, greeterService, healthService, userService, authorizationService, resourceService, teamService, roleService, datasourceService, menuService)
	app := newApp(bootstrap, serverServer, logger)
	return app, func() {
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}

package main

import (
	"fmt"
	"github.com/reoden/todolist/app/user/repository/db/dao"
	"github.com/reoden/todolist/app/user/service"
	"github.com/reoden/todolist/config"
	"github.com/reoden/todolist/idl/pb"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	config.Init()
	dao.InitDB()

	etcdReg := registry.NewRegistry(registry.Addrs(fmt.Sprintf("%s:%s", config.EtcdHost, config.EtcdPort)))

	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address(config.UserServiceAddress),
		micro.Registry(etcdReg),
	)

	microService.Init()

	_ = pb.RegisterUserServiceHandler(microService.Server(), service.GetUserSrv())
	_ = microService.Run()
}

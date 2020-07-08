package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	_ "micro/AppInit"
	"micro/Services"
	"micro/ServicesImpl"
)

func main() {
	consul := etcd.NewRegistry(
		registry.Addrs("localhost:2379"),
	)
	myService := micro.NewService(
		micro.Name("yuxiang.li.myapp"),
		micro.Address("localhost:8001"),
		micro.Registry(consul),
	)
	//此处利用继承，传入自己定于的struct，去执行自己定于的call
	//Services.RegisterTestServiceHandler(myService.Server(), new(ServicesImpl.TestService))
	Services.RegisterUserServiceHandler(myService.Server(), new(ServicesImpl.UserServices))
	myService.Run()
}

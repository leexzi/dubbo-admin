package application

const (
	serviceFile = `package service

import (
	"context"
)

import (
	"dubbo.apache.org-go-app/api"

	"dubbo.apache.org.apache.org/dubbo.apache.org-go/v3/common/logger"
	"dubbo.apache.org.apache.org/dubbo.apache.org-go/v3/config"
)

type GreeterServerImpl struct {
	api.UnimplementedGreeterServer
}

func (s *GreeterServerImpl) SayHello(ctx context.Context, in *api.HelloRequest) (*api.User, error) {
	logger.Infof("Dubbo-go GreeterProvider get user name = %s\n", in.Name)
	return &api.User{Name: "Hello " + in.Name, Id: "12345", Age: 21}, nil
}

func init() {
	config.SetProviderService(&GreeterServerImpl{})
}

`
)

func init() {
	fileMap["serviceFile"] = &fileGenerator{
		path:    "./pkg/service",
		file:    "service.go",
		context: serviceFile,
	}
}

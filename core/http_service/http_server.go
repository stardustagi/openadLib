package http_service

import (
	"github.com/hashicorp/go-hclog"
	"github.com/stardustagi/openadLib/core/errors"
)

type IHttpService interface {
	SetupHttpSrv(c Config) errors.StackError
	StartService() errors.StackError
	StopService() errors.StackError
	AddRoute()
}

type HttpService struct {
	logger *hclog.Logger
	config Config
}

func (h HttpService) SetupHttpSrv(c Config) errors.StackError {
	//TODO implement me
	panic("implement me")
}

func (h HttpService) StartService() errors.StackError {
	//TODO implement me
	panic("implement me")
}

func (h HttpService) StopService() errors.StackError {
	//TODO implement me
	panic("implement me")
}

func (h HttpService) AddRoute() {
	//TODO implement me
	panic("implement me")
}

func NewHttpService() IHttpService {
	return &HttpService{}
}

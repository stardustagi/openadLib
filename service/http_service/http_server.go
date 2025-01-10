package http_service

import (
	"fmt"
	"github.com/go-kit/kit/sd/etcd"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-hclog"
	"github.com/labstack/echo/v4"
	"github.com/stardustagi/openadLib/core/errors"
	"sync"
)

// IHttpService 抽象接口
type IHttpService interface {
	SetupHttpSrv(c Config, middleware ...echo.MiddlewareFunc) *errors.StackError
	Start() *errors.StackError
	Stop() *errors.StackError
	GetGroup(group string) *echo.Group
	NewSchema() *Schema[any, any]
	InitRoute()
}

// HttpService 主服务
type HttpService struct {
	logger        *hclog.Logger
	config        Config
	stop          chan bool
	srv           *echo.Echo
	groupInstance map[string]*echo.Group
	wg            sync.WaitGroup
	etcClient     etcd.Client
}

func (h *HttpService) SetGroup(group string, g *echo.Group) {
	h.groupInstance[group] = g
}

func AddRoute[Req any, Resp any](method, path string, group *echo.Group, handler func(req Req) (resp Resp, err *errors.StackError)) {
	group.Add(method, path, func(c echo.Context) error {
		var req Req
		if err := c.Bind(&req); err != nil {
			return c.JSON(500, "bind request error")
		}
		if err := c.Validate(req); err != nil {
			return c.JSON(500, "validate request error")
		}
		if resp, err := handler(req); err != nil {
			return c.JSON(err.Code(), err.Msg())
		} else {
			return c.JSON(200, resp)
		}
	})
}

func (h *HttpService) SetupHttpSrv(c Config, middleware ...echo.MiddlewareFunc) *errors.StackError {
	h.srv = echo.New()
	h.srv.Validator = &CustomValidator{Validator: validator.New()}
	h.config = c
	for _, v := range c.Group {
		h.groupInstance[v] = h.srv.Group(v, middleware...)
		//设置中间件
	}
	return nil
}

func (h *HttpService) GetGroup(group string) *echo.Group {
	return h.groupInstance[group]
}

func (h *HttpService) NewSchema() *Schema[any, any] {
	return &Schema[any, any]{}
}

func (h *HttpService) Start() *errors.StackError {
	addr := fmt.Sprintf("%s:%d", h.config.IP, h.config.Port)
	err := h.srv.Start(addr)
	if err != nil {
		return errors.New("http service start error", 500, err)
	}
	return nil
}

func (h *HttpService) Stop() *errors.StackError {
	//TODO implement me
	panic("implement me")
}

func (h *HttpService) InitRoute() {
	//TODO implement me
	panic("implement	me")
}
func NewHttpService() IHttpService {
	return &HttpService{
		groupInstance: make(map[string]*echo.Group),
	}
}

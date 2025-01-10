package http_service

import (
	"github.com/labstack/echo/v4"
	"github.com/stardustagi/openadLib/core/errors"
	"testing"
)

type TestReq struct {
}
type TestResp struct {
	Message string `json:"message"`
}

func TestHttpServer(t *testing.T) {
	t.Log("HttpServerTest")
	c := Config{
		IP:    "127.0.0.1",
		Port:  10001,
		Group: []string{"/v1", "/v2"},
	}
	srv := NewHttpService()
	if err := srv.SetupHttpSrv(c, Cors(), RequestID(), Request()); err != nil {
		panic(err)
	}
	AddRoute[TestReq, TestResp]("POST", "/test", srv.GetGroup("/v1"), func(req TestReq) (resp TestResp, err *errors.StackError) {
		// http事件处理
		return TestResp{
			Message: "test",
		}, nil
	})
	AddRoute[TestReq, TestResp]("POST", "/test1", srv.GetGroup("/v2"), func(req TestReq) (resp TestResp, err *errors.StackError) {
		// http 事件处理
		return TestResp{
			Message: "test1",
		}, nil
	})
	g := &echo.Group{
		prefix: "/test3",
	}
	srv.SetGroup("/test3", g)
	if err := srv.Start(); err != nil {
		panic(err)
	}
}

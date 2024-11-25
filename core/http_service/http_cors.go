package http_service

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Cors 处理跨域请求,支持options访问
func Cors() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
			c.Response().Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Response().Header().Set("Access-Control-Allow-Headers", "*")
			c.Response().Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
			//放行所有OPTIONS方法
			method := c.Request().Method
			if method == "OPTIONS" {
				return c.NoContent(http.StatusNoContent)
			}
			// 处理请求
			return next(c)
		}
	}
}

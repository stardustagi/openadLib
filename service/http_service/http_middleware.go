package http_service

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stardustagi/openadLib/core/logger"
	"io"
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

type bodyDumpResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w *bodyDumpResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// Request Response 记录请求日志
func Request() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 取得request body
			requestBody := ""
			b, err := io.ReadAll(c.Request().Body)
			if err != nil {
				requestBody = "failed to request body"
			} else {
				requestBody = string(b)
				c.Request().Body = io.NopCloser(bytes.NewBuffer(b))
			}
			host := c.Request().Host
			uri := c.Request().RequestURI
			method := c.Request().Method
			agent := c.Request().UserAgent()

			// 取得 response body
			respBody := new(bytes.Buffer)
			mw := io.MultiWriter(c.Response().Writer, respBody)
			writer := &bodyDumpResponseWriter{Writer: mw, ResponseWriter: c.Response().Writer}
			c.Response().Writer = writer
			status := c.Response().Status
			ip := echo.ExtractIPFromXFFHeader()(c.Request())
			logger.Info(fmt.Sprintf("request :'%s %s' %d %s '-' '%s' %s '%s'", method, uri, status, ip, agent, host, requestBody))
			err = next(c)
			if err != nil {
				return err
			}
			logger.Info(fmt.Sprintf("response: '%s'", respBody.String()))
			return nil
		}
	}
}

// RequestID Middleware to add a unique ID to each request
func RequestID() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := uuid.New().String()
			c.Set("RequestID", requestID)
			c.Response().Header().Set("X-Request-ID", requestID)
			return next(c)
		}
	}
}

package http_service

import "github.com/stardustagi/openadLib/core/errors"

type HttpMethod string

const (
	GET     HttpMethod = "GET"
	POST    HttpMethod = "POST"
	PUT     HttpMethod = "PUT"
	DELETE  HttpMethod = "DELETE"
	PATCH   HttpMethod = "PATCH"
	OPTIONS HttpMethod = "OPTIONS"
	HEAD    HttpMethod = "HEAD"
)

// Schema 接口属性
type Schema[I any, O any] struct {
	Description string                                       `json:"description"`
	Authorized  string                                       `json:"authorized"`
	Input       I                                            `json:"input,omitempty"`
	Output      O                                            `json:"output,omitempty"`
	Method      HttpMethod                                   `jons:"method"`
	Handler     func(req I) (resp O, err *errors.StackError) `json:"-"`
}

func NewSchema[I any, O any](description string, authorized string, input I, output O, handler func(req I) (resp O, err *errors.StackError)) *Schema[I, O] {
	return &Schema[I, O]{
		Description: description,
		Authorized:  authorized,
		Input:       input,
		Output:      output,
		Handler:     handler,
	}
}

func (s *Schema[I, O]) GetRequestBody() I {
	return s.Input
}

func (s *Schema[I, O]) GetResponseBody() O {
	return s.Output
}

package http

import (
	"net/http"

	"github.com/dzeqkon/pkg/errors"
)

type Response struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
}

func WriteResponse(err error, data interface{}) (int, Response) {
	if err != nil {
		coder := errors.ParseCoder(err)
		return coder.HTTPStatus(), Response{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
		}
	}
	return http.StatusOK, data.(Response)
}

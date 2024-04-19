package http

import "github.com/dzeqkon/pkg/errors"

type Response struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Reference string `json:"reference,omitempty"`
}

func WriteResponse(err error, data interface{}) Response {
	if err != nil {
		coder := errors.ParseCoder(err)
		return Response{
			Code:      coder.Code(),
			Message:   coder.String(),
			Reference: coder.Reference(),
		}
	}
	return data.(Response)
}

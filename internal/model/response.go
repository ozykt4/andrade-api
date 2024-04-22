package model

import "strings"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func NewSuccessResponse(data any, args ...any) *Response {
	response := &Response{
		Status: 200,
		Data:   data,
	}

	populateRes(response, args...)

	return response
}

func NewErrorResponse(err error, args ...any) *Response {
	response := &Response{
		Status: 400,
	}

	if err != nil {
		response.Message = err.Error()
	}

	populateRes(response, args...)

	return response
}

func populateRes(response *Response, args ...any) *Response {
	if len(args) >= 1 {
		response.Status = args[0].(int)
	}

	if len(args) == 2 {
		response.Message = strings.Join([]string{response.Message, args[1].(string)}, " ")
	}

	return response
}

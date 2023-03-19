package utils

import (
	"encoding/json"
	"net/http"
)

const (
	contentType              = "Content-Type"
	contentTypeValueJson     = "application/json; charset=utf-8"
	contentTypeValueHtml     = "text/html; charset=utf-8"
	xContentTypeOptions      = "X-Content-Type-Options"
	xContentTypeOptionsValue = "nosniff"

	StatusInternalErr  = "STATUS_INTERNAL_ERROR"
	StatusBadRequest   = "STATUS_BAD_REQUEST"
	StatusUnauthorized = "STATUS_UNAUTHORIZED"
	StatusForbidden    = "STATUS_FORBIDDEN"
	StatusNotFound     = "STATUS_NOT_FOUND"
)

// Response is object for response http
type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Err  string      `json:"error"`
	Msg  string      `json:"message"`
}

func SetResponseJSON(code int, data interface{}, err string, msg string) *Response {
	return &Response{
		Code: code,
		Data: data,
		Err:  err,
		Msg:  msg,
	}
}

// JSONResponseWithErr is method for return http json err
func (r *Response) JSONResponseWithErr(w http.ResponseWriter) {
	w.Header().Set(contentType, contentTypeValueJson)
	w.Header().Set(xContentTypeOptions, xContentTypeOptionsValue)
	w.WriteHeader(r.Code)
	json.NewEncoder(w).Encode(r)
}

// JSONResponse JSONResponseWithErr is method for return http json
func (r *Response) JSONResponse(w http.ResponseWriter) {
	w.Header().Set(contentType, contentTypeValueJson)
	w.Header().Set(xContentTypeOptions, xContentTypeOptionsValue)
	w.WriteHeader(r.Code)
	json.NewEncoder(w).Encode(r)
}

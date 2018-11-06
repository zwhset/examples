package defs

import (
	"encoding/json"
	"log"
	"net/http"
)

// 定义错误
var (
	ResJSON = Response{Code: "0000", Message: "Success"}

	ErrJSONParseFailed = ErrResponse{HttpSC: 403, Error: Err{Error: "Json Parse Failed.", ErrorCode: "0001"}}
	ErrNotAuthUser     = ErrResponse{HttpSC: 401, Error: Err{Error: "User authentication failed", ErrorCode: "0002"}}
)

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error  Err
}

// 正确返回方法
func HttpResponse(w http.ResponseWriter, rs interface{}) {
	// json
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(rs); err != nil {
		log.Printf(err.Error())
		ErrHttpResponse(w, ErrJSONParseFailed)
	}
}

// 错误返回方法
func ErrHttpResponse(w http.ResponseWriter, err ErrResponse) {
	log.Printf("%#v", err)
	http.Error(w, err.Error.Error, err.HttpSC)
}

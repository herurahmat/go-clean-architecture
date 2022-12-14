package helper

import (
	"encoding/json"
	"fmt"
	"log"
	http2 "net/http"
	"runtime/debug"
)

type Pages struct {
	Total      uint64 `json:"total"`
	PerPage    uint64 `json:"per_page"`
	Current    uint64 `json:"current"`
	TotalPages uint64 `json:"total_pages"`
}

type View struct {
	Status       bool        `json:"status"`
	Code         string      `json:"code"`
	Message      string      `json:"message"`
	ErrorMessage interface{} `json:"error_message"`
	Data         interface{} `json:"data"`
	Pagination   Pages       `json:"pagination"`
}

type response struct {
	view View
}

func (r *response) JSON() []byte {
	result, err := json.Marshal(r.view)
	if err != nil {
		return nil
	}
	return result
}

func ResponseHttp(w http2.ResponseWriter, view interface{}, httpStatus int) error {
	res, err := json.Marshal(view)
	if err != nil {
		return err
	}

	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, X-Auth-Token, x-api-Key, X-localization, channel, Channel, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	if _, err := w.Write(res); err != nil {
		return err
	}

	return nil
}

func Recover(w http2.ResponseWriter, code, errorMessage string, httpStatus int) {
	msg := recover()
	if msg != nil {
		log.Println(string(debug.Stack()))

		ResponseHttp(w, View{
			Code:         code,
			ErrorMessage: fmt.Sprintf("%v", msg),
			Message:      errorMessage,
		}, httpStatus)
	}
}

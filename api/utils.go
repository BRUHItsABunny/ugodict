package api

import (
	"github.com/BRUHItsABunny/gOkHttp/constants"
	"net/http"
)

func DefaultHeaders() http.Header {
	return http.Header{
		"user-agent":   {"okhttp/3.12.1"},
		"content-type": {constants.MIMEApplicationJSON},
		"accept":       {constants.MIMEApplicationJSON},
	}
}

package handlers

import (
	"basicGo/config"
	"net/http"
	"fmt"
)

func Login(context config.Context)http.HandlerFunc {
	return func(r http.Response,rq http.Request) {
		url := rq.URL
		fmt.Println(url.Path)
	}
}
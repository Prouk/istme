package main

// generate exemple
//go:generate jade -writer -fmt -pkg=static -d=static\generated .\static\jade\index.jade

import (
	"context"
	"main/routes"
	"net/http"
	"strings"
)

var ctx = context.Background()

type MyHandler struct{}

func (handler MyHandler) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	arg := strings.Split(req.RequestURI, "/")
	switch arg[1] {
	case "":
		routes.Home(wr, req)
	case "img":
		routes.Files(wr, req)
	default:
		routes.Home(wr, req)
	}
}

func main() {
	handler := &MyHandler{}
	http.ListenAndServe(":8080", handler)
}

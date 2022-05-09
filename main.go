package main

// generate exemple
//go:generate jade -writer -fmt -pkg=static -d=static\\generated .\\static\\jade\\index.jade

import (
	"context"
	"main/routes"
	"net/http"
	"strings"
)

var ctx = context.Background()

type MyHandler struct{}

func (handler MyHandler) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	url := strings.Split(req.RequestURI, "/")
	arg := strings.Split(url[1], "?")
	if req.Method == "GET" {
		switch arg[0] {
		case "":
			routes.Home(wr, req)
		case "style":
			routes.MainFiles(wr, req)
		case "script":
			routes.MainFiles(wr, req)
		case "img":
			routes.Files(wr, req)
		case "apiFfxiv":
			routes.FfxivApi(wr, req)
		default:
			routes.Home(wr, req)
		}
	} else {
		switch arg[0] {
		case "apiTest":
			routes.SimpleApiTest(wr, req)
		default:
			routes.ErrorApiTest(wr, req)
		}
	}
}

func main() {
	handler := &MyHandler{}
	http.ListenAndServe(":8080", handler)
}

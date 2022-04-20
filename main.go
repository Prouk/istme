package main

// generate exemple
//jade -writer -fmt -pkg=static -d=static\generated .\static\jade\index.jade

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
	if req.Method == "GET" {
		switch arg[1] {
		case "":
			routes.Home(wr, req)
		case "style":
			routes.MainFiles(wr, req)
		case "script":
			routes.MainFiles(wr, req)
		case "img":
			routes.Files(wr, req)
		default:
			routes.Home(wr, req)
		}
	} else {
		switch arg[1] {
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

package routes

import (
	"io/ioutil"
	"net/http"
)

func Files(wr http.ResponseWriter, req *http.Request) {
	fileBytes, err := ioutil.ReadFile("static" + req.RequestURI)
	if err != nil {
		panic(err)
	}
	wr.WriteHeader(http.StatusOK)
	wr.Header().Set("Content-Type", "application/octet-stream")
	wr.Write(fileBytes)
}

func MainFiles(wr http.ResponseWriter, req *http.Request) {
	http.ServeFile(wr, req, "static"+req.RequestURI)
}

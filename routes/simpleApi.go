package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Msg  string
	Code int
}

func SimpleApiTest(wr http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	stringTest := req.Form.Get("text")
	file, err := os.OpenFile("testFile.txt", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err2 := file.WriteString(stringTest)
	if err2 != nil {
		log.Fatal(err2)
	}
	respTest := Response{
		"Response : Written " + stringTest + " in the testFile.txt of the server folder",
		001,
	}
	jsresp, err := json.Marshal(respTest)
	wr.Header().Set("Content-Type", "application/json")
	wr.Write(jsresp)
}

func ErrorApiTest(wr http.ResponseWriter, req *http.Request) {
	errorTest := Response{
		"Unknown API request",
		000,
	}
	jsresp, err := json.Marshal(errorTest)
	if err != nil {
		http.Error(wr, err.Error(), http.StatusInternalServerError)
		return
	}
	wr.Header().Set("Content-Type", "application/json")
	wr.Write(jsresp)
}

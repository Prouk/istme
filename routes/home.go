package routes

import (
	"main/static/generated"
	"net/http"
)

func Home(wr http.ResponseWriter, req *http.Request) {
	static.Jade_index(wr)
}

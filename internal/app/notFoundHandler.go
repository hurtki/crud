package app

import (
	"net/http"
)

func NotFoundHandle(res http.ResponseWriter, req *http.Request) {
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(http.StatusNotFound)
	res.Write([]byte(`
{
	"error": "resource not found"
}
	`))
}

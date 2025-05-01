package httpserver

import "net/http"

func NewHttpServer(port string, handler http.Handler) error {
	return http.ListenAndServe(":8080", handler)
}
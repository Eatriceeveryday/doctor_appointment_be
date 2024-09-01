package http

import "net/http"

func CreateNewServer(router *http.ServeMux) http.Server {
	server := http.Server{
		Handler: router,
		Addr:    ":8080",
	}

	return server
}

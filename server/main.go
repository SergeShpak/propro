package main

import (
	"net/http"
	"time"
)

func main() {
	srv := createHttpServer()
	srv.ListenAndServe()
}

func createHttpServer() *http.Server {
	r := newRouter()
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
		Addr:         ":8080",
	}
	return srv
}

func listenAndServe(srv *http.Server) error {
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

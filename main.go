package main

import (
	//"time"
	"net/http"
	_api "tongren/api"
)
func main() {
	router := _api.InitRouter()
	server := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		//ReadTimeout:    setting.ReadTimeout * time.Second,
		//WriteTimeout:   setting.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}

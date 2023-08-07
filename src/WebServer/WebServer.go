package webserver

import (
	"net/http"
	"time"
)

func WebServerInit(){
	server := &http.Server{
		Addr			: ":8080",
		MaxHeaderBytes	: 1 << 20,
		WriteTimeout	: 10 * time.Second,
		ReadTimeout		: 10 * time.Second,
		Handler			: Routes(),
	}

	server.ListenAndServe()
}
package webserver

import (
	model "myanimelisttop/src/WebServer/Model"
	"net/http"
)

func Routes() *http.ServeMux{
	mux := http.NewServeMux()

	mux.HandleFunc("/test", Test)
	mux.HandleFunc("/", model.Index)

	return mux
}

func Test(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("OK"))
}
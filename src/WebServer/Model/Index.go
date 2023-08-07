package model

import (
	"encoding/json"
	soup "myanimelisttop/src/Soup"
	"net/http"
	"strconv"
	"sync"
)

func Index(w http.ResponseWriter, r *http.Request){
	var pipe = make(chan []soup.Anime)
	limit, _ := strconv.ParseInt(r.URL.Query().Get("limit"), 10, 0)

	if (limit < 0) {
		limit = 0
	}

	m := sync.Mutex{}

	go soup.GetRequest(int(limit), pipe, &m)

	data := <- pipe

	JSON, err := json.Marshal(data)

	if err != nil{
		panic(err)
	}
	
	w.Header().Add("Content-Type", "application/json")
	w.Write(JSON)
}
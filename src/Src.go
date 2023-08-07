package src

import (
	"fmt"
	httprequest "myanimelisttop/src/HttpRequest"
	webserver "myanimelisttop/src/WebServer"
)



func Run(){
	var res []byte = httprequest.GetMyAnimeList(1)
	webserver.WebServerInit()

	fmt.Println(string(res))
}
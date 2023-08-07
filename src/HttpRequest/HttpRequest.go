package httprequest

import (
	"fmt"
	"io"
	"net/http"
)

func GetMyAnimeList(start int) []byte{
	var url string = fmt.Sprintf("https://myanimelist.net/topanime.php?limit=%d", start)
	response, err := http.Get(url)

	if err != nil {
		fmt.Println(err.Error())
	}

	resp, err := io.ReadAll(response.Body)

	if err != nil{
		fmt.Println(err.Error())
	}

	return resp
}
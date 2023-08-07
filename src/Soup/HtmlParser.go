package soup

import (
	httprequest "myanimelisttop/src/HttpRequest"
	"strconv"
	"sync"

	"github.com/anaskhan96/soup"
)

func GetRequest(limit int, pipe chan []Anime, m *sync.Mutex){
	m.Lock()
	data := httprequest.GetMyAnimeList(limit)
	m.Unlock()
	doc := soup.HTMLParse(string(data))

	tr := doc.FindAll("tr", "class", "ranking-list")
	AnimeList := []Anime{}

	for _, i := range tr {
		var Cell Anime
		rank, _ := strconv.ParseInt(i.Find("span").Text(), 10, 0)
		title := i.Find("h3", "class", "hoverinfo_trigger").Find("a").Text()
		img := i.Find("img").Attrs()["data-src"]
		score, _ := strconv.ParseFloat(i.Find("span", "class", "score-label").Text(), 32)

		Cell.Rank = int(rank)
		Cell.Title = title
		Cell.Img = img
		Cell.Score = float32(score)

		AnimeList = append(AnimeList, Cell)
	}

	pipe <- AnimeList
	
}
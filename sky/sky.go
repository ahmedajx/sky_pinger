package sky

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SingleArticleData struct {
	Link      string `json:"link"`
	Img       string `json:"imgsrc"`
	ShortDesc string `json:"shortdesc"`
	Title     string `json:"title"`
}

func LatestChelseaNews() string {
	response, _ := http.Get("http://skysportsapi.herokuapp.com/sky/football/getteamnews/chelsea/v1.0/")
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
	var links []SingleArticleData
	json.Unmarshal(body, &links)
	return links[0].Title
}

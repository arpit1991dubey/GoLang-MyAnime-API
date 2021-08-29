package main

import (
	"encoding/json" //To encode and decode json
	"fmt"
	"github.com/gorilla/mux" //The name mux stands for "HTTP request multiplexer". Like the standard http.ServeMux, mux.Router matches incoming requests against a list of registered routes
	//"os"

	//"golang.org/x/text/number"
	"log" //To log Errors
	//"math/rand"              // To generate random values
	"net/http" //To create an http server
	//"strconv"                //To convert string into integrs or vice verso
	// and calls a handler for the route that matches the URL or other conditions
	"github.com/gocolly/colly/v2"
	//"github.com/PuerkitoBio/goquery"
)


type statistics struct {
	Rank       string `json:"rank"`
	Polularity string `json:"polularity"`
	Score      string `json:"score"`
	Favorites  int    `json:"favorites"`
	Members    int    `json:"members"`
}

type alternatives struct {
	English  string `json:"English"`
	Synonyms string `json:"Synonyms"`
	Japanese string `json:"Japanese"`
}
type information struct {
	Type      string `json:"type"`
	Episodes  string `json:"episodes"`
	Status    string `json:"status"`
	Aired     string `json:"aired"`
	Premiered string `json:"premiered"`
	Broadcast string `json:"broadcast"`
	Producers string `json:"producers"`
	Licensors string `json:"licensors"`
	Studios   string `json:"studios"`
	Source    string `json:"source"`
	Genres    string `json:"genres"`
	Duration  string `json:"duration"`
	Rating    string `json:"rating"`
} 

type voiceChar struct {
	CharectorName         string `json:"charectorName"`
	Role                  string `json:"role"`
	VoiceActorName        string `json:"voiceActorName"`
	VoiceActorNationality string `json:"voiceActorNationality"`
}
type staff struct {
	Name string `json:"name"`
	Role string `json:"role"`
}
type Anime struct {
	Title      string `json:"title"`
	Statistics statistics  `json:"statistics"`
	Synopsis          string        `json:"synopsis"`
	Background        string        `json:"background"`
	RelatedAnime      []interface{} `json:"relatedAnime"`
	VoiceChars        []voiceChar   `json:"voiceChars"`
	Staff             []staff       `json:"staff"`
	OpeningTheme      string        `json:"opening Theme"`
	EndingTheme       string        `json:"ending Theme"`
	Reviews           []interface{} `json:"reviews"`
	AlternativeTitles alternatives `json:"alternativeTitles"`
	Information information     `json:"information"`
}

 //"anime_detail_related_anime"

/******************Function to get a anime details**************************/

func getAnime(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id:=params["id"]
	url:="https://myanimelist.net/anime/"+id
	
	var responseData Anime
	var isVoiceCharsDone bool =false
	c := colly.NewCollector()
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})
	
	// c.OnHTML("a",func(e *colly.HTMLElement){      INFORMATION
	// 	switch(e.Attr("href")){
	// 	case "https://myanimelist.net/topanime.php?type=tv":
	// 		responseData.Information.Type=e.Text
	// 	case "https://myanimelist.net/anime/season/2021/summer":
	// 		responseData.Information.Premiered=e.Text
	// 	case "/anime/producer/159/":
	// 		responseData.Information.Producers=e.Text
	// 	}
	// })
	c.OnHTML("h1",func(e *colly.HTMLElement){
		switch(e.Attr("class")){
		case "title-name h1_bold_none":
			responseData.Title=e.ChildText("strong")
		}
	})
	c.OnHTML("td",func(e *colly.HTMLElement){
		switch(e.Attr("valign")){
		case "top":
			responseData.Background=e.Text
		}
	})
	c.OnHTML("div",func(e *colly.HTMLElement){
		switch(e.Attr("style")){
		case "width: 225px":
			e.ForEach("div ", func(i int, h *colly.HTMLElement) {
				responseData.Information = information{ 
					Type: h.ChildText("div > a"),
					// Role: h.ChildText("td > div > small"),
					// VoiceActorName: h.ChildText("td > table > tbody > tr > td >a"),
					// VoiceActorNationality: h.ChildText("td > table > tbody > tr > td > small"),
				}
			})
			
		}
	})
	c.OnHTML("div",func(e *colly.HTMLElement){
		switch(e.Attr("class")){
		case "score-label score-8":
			responseData.Statistics.Score=e.Text
		case "detail-characters-list clearfix":
			if !isVoiceCharsDone {
				e.ForEach("div > div > table > tbody > tr", func(i int, h *colly.HTMLElement) {
					responseData.VoiceChars = append(responseData.VoiceChars, voiceChar{ 
						CharectorName: h.ChildText("td > h3 > a"),
						Role: h.ChildText("td > div > small"),
						VoiceActorName: h.ChildText("td > table > tbody > tr > td >a"),
						VoiceActorNationality: h.ChildText("td > table > tbody > tr > td > small"),
					})
				})
				isVoiceCharsDone=true
			}else{
				e.ForEach("div > div > table > tbody > tr", func(i int, h *colly.HTMLElement) {
					responseData.Staff = append(responseData.Staff, staff{ 
						Name: h.ChildText("td > a"),
						Role: h.ChildText("td > div> small"),
					})
				})
			}
		case "margin-top: 15px;":
			responseData.Background=e.Text
		case "theme-songs js-theme-songs opnening":
			responseData.OpeningTheme=e.ChildText("span")
		case "theme-songs js-theme-songs ending":
			responseData.EndingTheme=e.ChildText("span")
		// case "spaceit_pad":
		// 	responseData.AlternativeTitles.English=e.ChildText("span  >")
		
		}
	})
	c.OnHTML("span",func(e *colly.HTMLElement){
		switch(e.Attr("class")){
		case "numbers ranked":
			responseData.Statistics.Rank=e.ChildText("span > strong")
		case "numbers popularity":
			responseData.Statistics.Polularity=e.ChildText("span > strong")
		}
	})
	c.OnHTML("p", func(e *colly.HTMLElement) {
		switch(e.Attr("itemprop")){
		case "description":
			responseData.Synopsis=e.Text
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		json.NewEncoder(w).Encode(responseData)
	})
	c.Visit(url)


}

/******************Main Function to aggregate methods and functions**************************/
func main() {
	// port:=os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/anime/{id}", getAnime).Methods("GET")
	fmt.Println("STARTING SERVER AT PORT 8000\n")
	log.Fatal(http.ListenAndServe(":"+"8000", r))
}

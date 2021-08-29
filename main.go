package main

/* -------------------------------------------------------------------------- */
/*                        Importing Necessary packages                        */
/* -------------------------------------------------------------------------- */

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

/* -------------------------------------------------------------------------- */
/*                         Defining our Data structure                        */
/* -------------------------------------------------------------------------- */

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
	Title             string        `json:"title"`
	Statistics        statistics    `json:"statistics"`
	Synopsis          string        `json:"synopsis"`
	Background        string        `json:"background"`
	RelatedAnime      []interface{} `json:"relatedAnime"`
	VoiceChars        []voiceChar   `json:"voiceChars"`
	Staff             []staff       `json:"staff"`
	OpeningTheme      string        `json:"opening Theme"`
	EndingTheme       string        `json:"ending Theme"`
	Reviews           []interface{} `json:"reviews"`
	AlternativeTitles alternatives  `json:"alternativeTitles"`
	Information       information   `json:"information"`
}

 /* -------------------------------------------------------------------------- */
 /*          Function to get a anime details with the passed anime ID          */
 /* -------------------------------------------------------------------------- */

func getAnime(w http.ResponseWriter, r *http.Request) { // Package http provides HTTP client and server implementations.
	//An http.ResponseWriter value assembles the HTTP server's response; by writing to it, we send data to the HTTP client

	w.Header().Set("Content-Type", "application/json") //setting headers
	params := mux.Vars(r)
	id := params["id"]                           //Getting the id from params
	url := "https://myanimelist.net/anime/" + id // Appending the user ID with URL of myanimelist.net

	var responseData Anime //storing the response data of type Anime in responseData
	var isVoiceCharsDone bool = false
	var infoCount int = 0

	/* ---------- Initialising colly and stating of the scraping phase ---------- */

	c := colly.NewCollector()                      // Instantiate default collector
	c.OnError(func(_ *colly.Response, err error) { // Handling error with colly
		log.Println("Something went wrong:", err)
	})

	/* ----------------------- Finding title of the anime ----------------------- */

	c.OnHTML("h1", func(e *colly.HTMLElement) { // On every a element which has h1 attribute call callback
		switch e.Attr("class") { // Check for the class atribute and match the respective class with the case then insert values
		case "title-name h1_bold_none":
			responseData.Title = e.ChildText("strong")
		}
	})
	/* --------------------- Finding background of the anime -------------------- */

	c.OnHTML("td", func(e *colly.HTMLElement) {
		switch e.Attr("valign") { // Check for the valign atribute and match the respective class with the case then insert values
		case "top":
			responseData.Background = e.Text
		}
	})

	/* -------------------- Finding information of the anime -------------------- */

	c.OnHTML("div", func(e *colly.HTMLElement) {
		switch e.Attr("id") {
		case "content":
			if(infoCount==0){
				e.ForEach("table > tbody > tr > td > div", func(i int, h *colly.HTMLElement) { 
					x := h.Text
				})
			}
		}
	})

	/* ----- Finding Voice Charectors,score,staff,openingtheme,closing theme ----- */

	c.OnHTML("div", func(e *colly.HTMLElement) {
		switch e.Attr("class") { // Check for the class atribute and match the respective class with the case then insert values
		case "score-label score-8":
			responseData.Statistics.Score = e.Text
		case "detail-characters-list clearfix":
			if !isVoiceCharsDone {
				e.ForEach("div > div > table > tbody > tr", func(i int, h *colly.HTMLElement) { //iterating from the parent div to the child div using for each
					responseData.VoiceChars = append(responseData.VoiceChars, voiceChar{ //appending values in the voicechars slice
						CharectorName:         h.ChildText("td > h3 > a"),                          //Re-iterating to the exact position of the element to get the values
						Role:                  h.ChildText("td > div > small"),                     //Re-iterating to the exact position of the element to get the values
						VoiceActorName:        h.ChildText("td > table > tbody > tr > td >a"),      //Re-iterating to the exact position of the element to get the values
						VoiceActorNationality: h.ChildText("td > table > tbody > tr > td > small"), //Re-iterating to the exact position of the element to get the values
					})
				})
				isVoiceCharsDone = true //Since voice charectors and staff have the sme div class and were being implemented in the saeme function
				//we have used a flag variable isVoiceCharsDone to indicate if the Voicechars is already visited
			} else {
				e.ForEach("div > div > table > tbody > tr", func(i int, h *colly.HTMLElement) { //iterating from the parent div to the child div using for each
					responseData.Staff = append(responseData.Staff, staff{ //appending values in the staff slice
						Name: h.ChildText("td > a"),
						Role: h.ChildText("td > div> small"),
					})
				})
			}
		case "margin-top: 15px;": //Similarly matching other classes
			responseData.Background = e.Text
		case "theme-songs js-theme-songs opnening":
			responseData.OpeningTheme = e.ChildText("span")
		case "theme-songs js-theme-songs ending":
			responseData.EndingTheme = e.ChildText("span")
			// case "spaceit_pad":
			// 	responseData.AlternativeTitles.English=e.ChildText("span  >")

		}
	})

	/* ---------------- Finding rank and popularity of the anime ---------------- */

	c.OnHTML("span", func(e *colly.HTMLElement) {
		switch e.Attr("class") {
		case "numbers ranked":
			responseData.Statistics.Rank = e.ChildText("span > strong")
		case "numbers popularity":
			responseData.Statistics.Polularity = e.ChildText("span > strong")
		}
	})

	/* ---------------------- Finding synopsis of the anime --------------------- */

	c.OnHTML("p", func(e *colly.HTMLElement) {
		switch e.Attr("itemprop") {
		case "description":
			responseData.Synopsis = e.Text
		}
	})

	/* ---------------------- Finding synopsis of the anime --------------------- */

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
		json.NewEncoder(w).Encode(responseData)
	})
	c.Visit(url) // Start scraping on the url

}

/* -------------------------------------------------------------------------- */
/*       main function to aggregate methods and functions                     */
/* -------------------------------------------------------------------------- */

func main() {
	// port:=os.Getenv("PORT")
	r := mux.NewRouter()                                 //Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.
	r.HandleFunc("/anime/{id}", getAnime).Methods("GET") // mapping route URL paths to handlers
	fmt.Println("STARTING SERVER AT PORT 8000\n")        // More info on mux here ---> https://github.com/gorilla/mux
	log.Fatal(http.ListenAndServe(":"+"8000", r))
}

# Go MyAnime REST API - Trademarkia Assignmnet
A RESTful API build with Go which scrapes https://myanimelist.net with the anime ID and provides all the datails regarding that particular anime 

This RESTful API with Go is build using using **gorilla/mux** (A nice mux library) and **coly** (Fast and Elegant Scraping Framework for Gophers)

## Todo

- [x] Created data structer according to the website entries.
- [x] Added api endpoint for "anime/animeId".
- [x] Installed and implemented colly for scraping the website.
- [x] Organized the code with packages.
- [x] Scraped all the details from the website for the specific ID or Class tag.
- [x] Build and deployed the api on heroku.
- [x] Wrote test for the endpoint APIs.

## Installation & Run
```
git clone https://git.legalforcelaw.com/18BLC1084/18BLC1084-GoLang-2.git
```
```
# Dependencies installation
go get -u github.com/gorilla/mux
go get -u github.com/gocolly/colly/...
```

```
# Build and Run
cd 18BLC1084-GoLang-2
go build test
go run main.go

```

## API

#### /anime/id
* `GET` : Get all the details associated with that particular anime ID.


## Demo

* Demo Video - https://mighty-basin-31398.herokuapp.com/anime/41488
* Live API Endpoint : https://mighty-basin-31398.herokuapp.com/anime/41488

## Demo with different anime Ids
* https://mighty-basin-31398.herokuapp.com/anime/41488
* https://mighty-basin-31398.herokuapp.com/anime/39247
* https://mighty-basin-31398.herokuapp.com/anime/48580


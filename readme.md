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
git clone github.com/mingrammer/go-todo-rest-api-example
```

```
# Build and Run
cd go-todo-rest-api-example
go build test
go run main.go


# API Endpoint : https://mighty-basin-31398.herokuapp.com/anime/41488
```

## API

#### /anime/id
* `GET` : Get all the details associated with that anime ID.


## Demo

Demo Video - https://mighty-basin-31398.herokuapp.com/anime/41488




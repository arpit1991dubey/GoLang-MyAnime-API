package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router:=mux.NewRouter()
	router.HandleFunc("/anime/{id}",getAnime).Methods("GET")
	return router
}

func TestRouterEndpoint(t *testing.T){
	request,_:=http.NewRequest("GET","/anime/41487",nil)
	response :=httptest.NewRecorder()
	Router().ServeHTTP(response,request)
	assert.Equal(t,200,response.Code,"OK response is expected")
}

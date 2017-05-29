package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/chrisvdg/HTTPSniff/config"
)

// QueryHandler handles query requests
func QueryHandler(w http.ResponseWriter, r *http.Request, sc config.ServerConfig) {
	// print request path
	fmt.Println("-Request:", r.RequestURI)

	// fetch q parameter value
	var (
		q       string
		baseurl string
	)
	baseurl = "https://google.com/search?q="
	r.ParseForm()
	buf := r.Form["q"]
	if len(buf) < 1 {
		fmt.Println("Query not found")
		baseurl = "https://google.com" + r.RequestURI
	} else {
		q = buf[0]
	}

	// get request to google search
	fmt.Println(" - Request to: " + baseurl + q)
	resp, err := http.Get(baseurl + q)
	if err != nil {
		log.Fatal("Something went wrong making the request", err)
	}

	// serve the response to client
	defer resp.Body.Close()
	// headers
	for name, values := range resp.Header {
		w.Header()[name] = values
		if name == "Content-Type" {
			fmt.Println(name, ":", values)
		}
	}
	w.WriteHeader(resp.StatusCode)
	// body
	io.Copy(w, resp.Body)

	// print response from request
	//fmt.Println(" - Response:", resp)
}

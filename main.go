package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/chrisvdg/HTTPSniff/config"
)

func main() {
	// init config
	sc, err := config.NewServerConfig("./config.json")
	if err != nil {
		fmt.Println("Could not get server config: ", err)
		return
	}

	// routes
	http.HandleFunc("/", handler)

	// run server
	fmt.Println("Listening on port", sc.GetPortString())
	http.ListenAndServe(sc.GetPortString(), nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// print request path
	fmt.Println(" - Request:", r.RequestURI)

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
		// probably requesting resource from google.com/... and not google.com/search/...
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
	}
	w.WriteHeader(resp.StatusCode)
	// body
	io.Copy(w, resp.Body)

	// print response from request
	fmt.Println(" - Respons:", resp)

}

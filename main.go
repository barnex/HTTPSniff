package main

import (
	"net/http"

	"fmt"

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
	fmt.Println("Request:", r.RequestURI)

	// fetch q parameter value
	r.ParseForm()
	buf := r.Form["q"]
	if len(buf) < 1 {
		fmt.Println("Query not found")
		fmt.Fprint(w, "Query not found")
		return
	}
	q := buf[0]

	// redirect to Google search
	s := fmt.Sprint("https://google.com/search?q=", q)
	http.Redirect(w, r, s, 303)

	// print redirection url
	fmt.Println("Redirected to:", s)
}

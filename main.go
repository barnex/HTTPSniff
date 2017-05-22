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
	http.HandleFunc("/", sniffer)

	// run server
	fmt.Println("Listening on port", sc.GetPortString())
	http.ListenAndServe(sc.GetPortString(), nil)
}

func sniffer(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.RequestURI)
}

package main

import (
	"net/http"

	"fmt"

	"github.com/chrisvdg/HTTPSniff/config"
	"github.com/chrisvdg/HTTPSniff/controllers"
)

func main() {
	// init config
	sc, err := config.NewServerConfig("./config.json")
	if err != nil {
		fmt.Println("Could not get server config: ", err)
		return
	}

	fmt.Println(sc)

	// routes
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// wrapper to pass along serverconfig to handler
		controllers.QueryHandler(w, r, sc)
	})

	// run server
	fmt.Println("Listening on port", sc.GetPortString())
	http.ListenAndServe(sc.GetPortString(), nil)
}

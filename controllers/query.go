package controllers

import (
	"fmt"
	"net/http"

	"github.com/chrisvdg/HTTPSniff/config"
)

// QueryHandler handles query requests
func QueryHandler(w http.ResponseWriter, r *http.Request, sc config.ServerConfig) {
	// print request path
	fmt.Println("Request:", r.RequestURI)

	// fetch q parameter value
	var s string
	r.ParseForm()
	buf := r.Form["s"]
	if len(buf) < 1 {
		s = sc.DefaultService
	} else {
		s = buf[0]

		// if service not present, assign default service
		if _, pres := sc.Services[s]; !pres {
			s = sc.DefaultService
		}
	}

	// execute corresponding request
	buf = r.Form["q"]
	var q string
	if len(buf) < 1 {
		q = ""
	} else {
		q = buf[0]
	}
	search := fmt.Sprint(sc.Services[s], q)
	http.Redirect(w, r, search, 303)

	// print redirection url
	fmt.Println("Redirected to:", search)
}

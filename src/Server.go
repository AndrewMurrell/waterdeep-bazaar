package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"strings"
	//"io/ioutil"
	"encoding/json"
)

func main() {

	http.HandleFunc("/api/", Api)
	http.Handle("/", http.FileServer(http.Dir("ui")))
	http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
	
}

/* Routers */
func Api(w http.ResponseWriter, r *http.Request) {

	path := html.EscapeString(r.URL.Path)
	switch {
	case strings.HasPrefix(path, "/api/settings/"):
		Handle_Settings(w, r)
	case strings.HasPrefix(path, "/api/options/"):
		Handle_Options(w, r)
	case strings.HasPrefix(path, "/api/results/"):
		Handle_Results(w, r)
	}
	
}

/* The REST Handlers */

func Handle_Settings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Settings: %q", strings.TrimPrefix(html.EscapeString(r.URL.Path), "/api/settings/"))
}

/*[
    {
	"option": "Number of Vendors",
	"value": "8"
    },
    {
	"option": "Kinds of Vendors",
	"value": "all"
    },
    {
	"option": "Max Level of Vendors",
	"value": "20"
    },
    {
	"option": "Min Level of Vendors",
	"value": "1"
    }
]*/


func Handle_Options(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == "GET":
		settings := Options{
			Option{Option: "Number of Vendors"},
			Option{Value: "8"},
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF=8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(settings)
	case r.Method == "POST":
		// TODO: update stuffs
		
	}
}

func Handle_Results(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Results: %q", strings.TrimPrefix(html.EscapeString(r.URL.Path), "/api/results/"))
}


/* JSON Output */

type Option struct {
	Option string `json:"option"`
	Value  string `json:"value"`
}

type Options []Option

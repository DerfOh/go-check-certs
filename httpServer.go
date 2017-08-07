package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// GET string constant
var GET = "GET"

// POST string constant
var POST = "POST"

func serveHTTP() {
	var dir string

	dir = *resultsDir

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/refresh", RefreshHandler)

	// This will serve files under http://localhost:8000/results/<filename>
	r.PathPrefix("/results/").Handler(http.StripPrefix("/results/", http.FileServer(http.Dir(dir))))

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

// HomeHandler for home requests
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	fmt.Println(r.Method)
	switch r.Method {
	case GET:
		refreshResults()
		tmpl.Execute(w, nil)
	case POST:
		r.ParseForm()
		// send addhost the hostname in the form, using index 0 because you should only add one at a time, otherwise go will interpret it as a stringarray
		fmt.Println(r.Form["hostname"][0])
		fmt.Println(r.Form["submit"][0])

		// switch to check if the add or remove button was pressed
		switch r.Form["submit"][0] {
		case "add":
			fmt.Println("adding " + r.Form["hostname"][0])
			addHost(r.Form["hostname"][0])
			refreshResults()
		case "remove":
			fmt.Println("removing " + r.Form["hostname"][0])
			removeHost(r.Form["hostname"][0])
			refreshResults()
		case "refresh":
			refreshResults()
		default:
			refreshResults()
		}
		tmpl.Execute(w, struct{ Success bool }{true})
	default:
		refreshResults()
		tmpl.Execute(w, struct{ Success bool }{true})
	}
}

// RefreshHandler removes results file, then generates a new one, this allows for a cURL cron or some other mechanism to refresh the results in an automated fashion
func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	refreshResults()
	w.Write([]byte("refresh complete"))
}

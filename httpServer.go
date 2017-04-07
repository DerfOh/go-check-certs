package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"

	"html/template"
)

// import (
// 	"html/template"
// 	"net/http"
// )
// type ContactDetails struct {
// 	Email   string
// 	Subject string
// 	Message string
// }

func serveHTTP() {
	var dir string

	dir = "./results"
	//  flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	//  flag.Parse()

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
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	
	refreshResults()
	tmpl.Execute(w, struct{ Success bool }{true})
}

// RefreshHandler removes results file, then generates a new one
func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	refreshResults()
	w.Write([]byte("Refresh page to view results."))
}

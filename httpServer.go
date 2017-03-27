package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func serveHTTP() {
	http.HandleFunc("/results.csv", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(response http.ResponseWriter, request *http.Request) {
	//response.Header().Set("Content-type", "text/html")
	response.Header().Set("Content-type", "text/csv")
	refreshResults()
	webpage, err := ioutil.ReadFile("results.csv")
	if err != nil {
		http.Error(response, fmt.Sprintf("results.csv file error %v", err), 500)
	}
	fmt.Fprint(response, string(webpage))
}

// remove results file, then generate a new one
func refreshResults(){
	err := os.Remove("results.csv")
	createOutPutFile()
	processHosts()
	if err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test title", Desc: "Test description", Content: "Test Content"},
	}
	fmt.Fprint(w, "All articles Endpoint Hit")
	json.NewEncoder(w).Encode(articles)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Homepage Endpoint Hit")
}

func dashboard(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Dashboard Endpoint Hit")
}

func handleReqests() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/app", dashboard)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleReqests()
}

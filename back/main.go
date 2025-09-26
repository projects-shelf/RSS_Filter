package main

import (
	"log"
	"net/http"
)

var blocklist []string

func initBlocklist() {
	var err error
	blocklist, err = LoadBlocklist("blocklist.txt")
	if err != nil {
		log.Fatalf("Failed to load blocklist: %v", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("url")
	if query == "" {
		http.Error(w, "Missing url parameter", http.StatusBadRequest)
		return
	}

	feed, err := FetchFeed(query)
	if err != nil {
		http.Error(w, "Failed to fetch feed", http.StatusInternalServerError)
		return
	}

	filteredItems := FilterFeedItems(feed, blocklist)

	rss, err := BuildRSSFeed(feed, filteredItems)
	if err != nil {
		http.Error(w, "Failed to generate RSS", http.StatusInternalServerError)
		return
	}

	WriteRSSResponse(w, rss)
}

func main() {
	initBlocklist()
	http.HandleFunc("/", handler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

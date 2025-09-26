package main

import (
	"net/http"
	"time"

	"github.com/gorilla/feeds"
	"github.com/mmcdole/gofeed"
)

func BuildRSSFeed(feedData *gofeed.Feed, filteredItems []*gofeed.Item) (string, error) {
	newFeed := &feeds.Feed{
		Title:       feedData.Title,
		Link:        &feeds.Link{Href: feedData.Link},
		Description: feedData.Description,
		Author:      &feeds.Author{Name: "RSS Filter"},
		Created:     time.Now(),
	}

	for _, item := range filteredItems {
		created := time.Now()
		if item.PublishedParsed != nil {
			created = *item.PublishedParsed
		}

		newFeed.Items = append(newFeed.Items, &feeds.Item{
			Title:       item.Title,
			Link:        &feeds.Link{Href: item.Link},
			Description: item.Description,
			Created:     created,
		})
	}

	return newFeed.ToRss()
}

func WriteRSSResponse(w http.ResponseWriter, rss string) {
	w.Header().Set("Content-Type", "application/rss+xml")
	w.Write([]byte(rss))
}

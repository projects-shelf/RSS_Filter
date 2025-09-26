package main

import (
	"strings"

	"github.com/mmcdole/gofeed"
)

func FilterFeedItems(feed *gofeed.Feed, blocklist []string) []*gofeed.Item {
	var filtered []*gofeed.Item
	for _, item := range feed.Items {
		blocked := false

		if item.Link != "" && IsBlocked(item.Link, blocklist) {
			blocked = true
		}

		// Check all links in content (optional)
		if !blocked && item.Content != "" {
			for _, domain := range blocklist {
				if strings.Contains(item.Content, domain) {
					blocked = true
					break
				}
			}
		}

		if !blocked {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

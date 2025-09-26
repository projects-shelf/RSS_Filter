package main

import (
	"bufio"
	"os"
	"strings"
)

func LoadBlocklist(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var blocklist []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			blocklist = append(blocklist, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return blocklist, nil
}

func IsBlocked(url string, blocklist []string) bool {
	for _, domain := range blocklist {
		if strings.Contains(url, domain) {
			return true
		}
	}
	return false
}

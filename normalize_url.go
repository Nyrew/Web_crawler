package main

import (
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", err
	}

	normalizedURL := parsedURL.Host + parsedURL.Path

	normalizedURL = strings.TrimSuffix(normalizedURL, "/")

	return normalizedURL, nil
}

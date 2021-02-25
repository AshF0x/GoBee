package fetcher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Fetch ...
func Fetch(domain string) string {
	resp, err := http.Get(domain)
	if err != nil {
		fmt.Printf("Could not read fetch page")
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Could not read request body")
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	return bodyString
}

// Slicer ...
func Slicer(body string, st string, en string) string {
	start := strings.Split(body, st)
	end := strings.Split(string(start[1]), en)
	currVersion := string(end[0])
	return currVersion
}

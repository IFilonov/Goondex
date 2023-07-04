package main

import (
	"Goondex/part2/pkg/crawler"
	"Goondex/part2/pkg/crawler/spider"
	"flag"
	"fmt"
	"log"
	"strings"
)

func printUrls(docs []crawler.Document, search string) {
	for _, doc := range docs {
		if strings.Contains(doc.URL, search) {
			fmt.Printf("- %s\n", doc.URL)
		}
	}
}

func main() {
	search := flag.String("s", "", "Search word")
	flag.Parse()

	urls := [2]string{"https://go.dev", "https://github.com/"}
	s := spider.New()
	for _, url := range urls {
		data, err := s.Scan(url, 2)
		if err != nil {
			log.Fatal(err)
		}
		printUrls(data, *search)
	}
}

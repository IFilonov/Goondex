package main

import (
	"Goondex/part2/pkg/crawler"
	"Goondex/part2/pkg/crawler/spider"
	index "Goondex/part3/pkg"
	"flag"
	"fmt"
	"sort"
)

func main() {
	search := flag.String("s", "", "Search word")
	flag.Parse()

	urls := [2]string{"https://go.dev", "https://github.com/"}
	s := spider.New()

	var allDocs []crawler.Document
	for _, url := range urls {
		docs, err := s.Scan(url, 2)
		if err != nil {
			fmt.Println("Ошибка сканирования ", url)
			continue
		}
		allDocs = append(allDocs, docs...)
	}

	for i := 0; i < len(allDocs); i++ {
		allDocs[i].ID = i
	}

	sort.Slice(allDocs, func(i, j int) bool {
		return allDocs[i].ID < allDocs[j].ID
	})

	index.IndexDocs(allDocs)
	indexed_docs := index.Words[*search]

	for _, index := range indexed_docs {
		doc := binarySearch(allDocs, index)
		fmt.Println(doc.Title)
	}
}

func binarySearch(docs []crawler.Document, value int) crawler.Document {

	start_index := 0
	end_index := len(docs) - 1

	for start_index <= end_index {

		median := (start_index + end_index) / 2

		if docs[median].ID < value {
			start_index = median + 1
		} else {
			end_index = median - 1
		}

	}
	return docs[start_index]
}

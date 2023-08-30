package main

import (
	"Goondex/part13/pkg/crawler"
	"Goondex/part13/pkg/crawler/spider"
	index "Goondex/part13/pkg/index"
	"Goondex/part13/pkg/webapp"
	"fmt"
)

var urls = []string{"https://go.dev", "https://github.com/"}

const depth = 2

func main() {
	var allDocs []crawler.Document
	allDocs, err := scanDocs()
	if err != nil {
		fmt.Println("Ошибка сканирования")
		return
	}
	setDocsIds(&allDocs)
	index.IndexDocs(allDocs)

	var apiData webapp.ApiData
	apiData.Docs = allDocs

	err = webapp.Start(apiData)
	if err != nil {
		fmt.Println("Ошибка запуска сервера")
		return
	}
}

func scanDocs() ([]crawler.Document, error) {
	var allDocs []crawler.Document
	s := spider.New()
	for _, url := range urls {
		docs, err := s.Scan(url, depth)
		if err != nil {
			return allDocs, err
		}
		allDocs = append(allDocs, docs...)
	}
	return allDocs, nil
}

func setDocsIds(docs *[]crawler.Document) {
	for i := 0; i < len(*docs); i++ {
		(*docs)[i].ID = i
	}
}

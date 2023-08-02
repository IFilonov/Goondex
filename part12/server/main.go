package main

import (
	index "Goondex/part12/pkg"
	"Goondex/part12/pkg/crawler"
	"Goondex/part12/pkg/crawler/spider"
	"Goondex/part12/pkg/webapp"
	"encoding/json"
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
	j, err := json.Marshal(index.Words)
	if err != nil {
		fmt.Println("Ошибка сериализации индекса")
		return
	}
	apiData.IndexDocs = j

	j, err = json.Marshal(allDocs)
	if err != nil {
		fmt.Println("Ошибка сериализации списка документов")
		return
	}
	apiData.Docs = j

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

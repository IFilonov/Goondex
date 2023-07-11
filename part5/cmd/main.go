package main

import (
	"Goondex/part2/pkg/crawler"
	"Goondex/part2/pkg/crawler/spider"
	index "Goondex/part3/pkg"
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
)

var Urls = []string{"https://go.dev", "https://github.com/"}

func main() {
	search := flag.String("s", "", "Search word")
	flag.Parse()

	var allDocs []crawler.Document

	r, err := os.Open("./docs.bin")
	if err == nil {
		allDocs = readDocs(r)
		r.Close()
	}

	was_scan := false
	if len(allDocs) == 0 {
		scanDocs(&allDocs)
		was_scan = true
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

	if was_scan {
		w, err := os.Create("./docs.bin")
		if err != nil {
			fmt.Println("Ошибка записи файла docs.bin")
			return
		}
		defer w.Close()
		writeDocs(allDocs, w)
	}

}

func scanDocs(allDocs *[]crawler.Document) {
	s := spider.New()
	for _, url := range Urls {
		docs, err := s.Scan(url, 2)
		if err != nil {
			fmt.Println("Ошибка сканирования ", url)
			continue
		}
		*allDocs = append(*allDocs, docs...)
	}
}

func readDocs(r io.Reader) []crawler.Document {
	var decodedDocs []crawler.Document
	d := gob.NewDecoder(r)
	err := d.Decode(&decodedDocs)
	if err != nil {
		panic(err)
	}
	return decodedDocs
}

func writeDocs(docs []crawler.Document, w io.Writer) {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	err := e.Encode(docs)
	if err != nil {
		panic(err)
	}
	w.Write(b.Bytes())
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

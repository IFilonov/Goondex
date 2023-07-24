package main

import (
	index "Goondex/part5/pkg"
	"Goondex/part5/pkg/crawler"
	"Goondex/part5/pkg/crawler/spider"
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
)

var Urls = []string{"https://go.dev", "https://github.com/"}

const Filename = "docs.bin"

func main() {
	search := getFlag()
	var allDocs []crawler.Document

	_, err := os.Stat("./" + Filename)
	if err != nil {
		scanDocs(&allDocs)
		setDocsId(&allDocs)
		sortDocs(&allDocs)
		writeFile(allDocs)
	} else {
		allDocs = readFile()
	}
	index.IndexDocs(allDocs)

	if search == nil {
		return
	}

	indexed_docs := index.Words[*search]
	for _, index := range indexed_docs {
		doc := binarySearch(allDocs, index)
		fmt.Println(doc.Title)
	}

}

func getFlag() *string {
	str := flag.String("s", "", "Search word")
	flag.Parse()
	if *str == "" {
		str = nil
	}
	return str
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

func setDocsId(docs *[]crawler.Document) {
	for i := 0; i < len(*docs); i++ {
		(*docs)[i].ID = i
	}
}

func sortDocs(docs *[]crawler.Document) {
	sort.Slice(*docs, func(i, j int) bool {
		return (*docs)[i].ID < (*docs)[j].ID
	})
}

func readFile() []crawler.Document {
	r, err := os.Open("./" + Filename)
	if err != nil {
		fmt.Println("Ошибка открытия файла ", Filename)
		return nil
	}
	docs := readDocs(r)
	r.Close()
	return docs
}

func readDocs(r io.Reader) []crawler.Document {
	var decodedDocs []crawler.Document
	d := gob.NewDecoder(r)
	err := d.Decode(&decodedDocs)
	if err != nil {
		fmt.Println("Ошибка декодирования данных файла ", Filename)
	}
	return decodedDocs
}

func writeFile(docs []crawler.Document) {
	w, err := os.Create("./" + Filename)
	if err != nil {
		fmt.Println("Ошибка создания файла ", Filename)
		return
	}
	defer w.Close()
	writeDocs(docs, w)
}

func writeDocs(docs []crawler.Document, w io.Writer) {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	err := e.Encode(docs)
	if err != nil {
		fmt.Println("Ошибка кодирования файла ", Filename)
		return
	}
	_, err = w.Write(b.Bytes())
	if err != nil {
		fmt.Println("Ошибка записи в файл ", Filename)
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

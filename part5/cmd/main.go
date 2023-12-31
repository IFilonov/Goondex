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
	if err == nil {
		allDocs, err = readFile()
		if err != nil {
			fmt.Printf("Ошибка чтения файла: %v", err)
			return
		}
	}

	if err != nil {
		allDocs = scanDocs()
		setDocsId(&allDocs)
		sortDocs(&allDocs)
		_, err = writeFile(allDocs)
		if err != nil {
			fmt.Printf("Ошибка записи в файл: %v", err)
			return
		}
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

func scanDocs() []crawler.Document {
	var allDocs []crawler.Document
	s := spider.New()
	for _, url := range Urls {
		docs, err := s.Scan(url, 2)
		if err != nil {
			fmt.Println("Ошибка сканирования ", url)
			continue
		}
		allDocs = append(allDocs, docs...)
	}
	return allDocs
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

func readFile() ([]crawler.Document, error) {
	f, err := os.Open("./" + Filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readDocs(f)
}

func readDocs(r io.Reader) ([]crawler.Document, error) {
	var decodedDocs []crawler.Document
	d := gob.NewDecoder(r)
	err := d.Decode(&decodedDocs)
	if err != nil {
		return nil, err
	}
	return decodedDocs, err
}

func writeFile(docs []crawler.Document) (int, error) {
	f, err := os.Create("./" + Filename)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	return writeDocs(docs, f)
}

func writeDocs(docs []crawler.Document, w io.Writer) (int, error) {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	err := e.Encode(docs)
	if err != nil {
		return 0, err
	}
	return w.Write(b.Bytes())
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

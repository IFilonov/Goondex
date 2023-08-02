package index

import (
	"Goondex/part12/pkg/crawler"
	"strings"

	"golang.org/x/exp/slices"
)

var Words = make(map[string][]int)

func IndexDocs(docs []crawler.Document) {
	for _, doc := range docs {
		words := strings.Split(strings.ToLower(doc.Title), " ")
		for _, word := range words {
			indexes := Words[word]
			if indexes == nil {
				Words[word] = []int{doc.ID}
			} else {
				if slices.Contains(indexes, doc.ID) {
					continue
				}
				Words[word] = append(indexes, doc.ID)
			}
		}
	}
}

func BinarySearch(docs []crawler.Document, value int) crawler.Document {

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

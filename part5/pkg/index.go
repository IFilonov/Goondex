package index

import (
	"Goondex/part5/pkg/crawler"
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

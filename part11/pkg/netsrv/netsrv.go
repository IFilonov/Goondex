// Package netsrv реализует telnet сервер для приема поисковых запросов.
package netsrv

import (
	"Goondex/part11/pkg/crawler"
	"Goondex/part11/pkg/crawler/spider"
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

const addr = "0.0.0.0:8000"
const depth = 2

var urls = []string{"https://go.dev", "https://github.com/"}

func StartNetSrv() error {
	var allDocs []crawler.Document
	allDocs, err := scanDocs()
	if err != nil {
		fmt.Println("Ошибка сканирования")
		return err
	}

	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		fmt.Printf("Ошибка открытия порта %v: %v\n", addr, err)
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Ошибка подключения: %v\n", err)
			continue
		}
		go handleConn(conn, allDocs)
	}
}

func handleConn(conn net.Conn, allDocs []crawler.Document) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Second * 600))
	r := bufio.NewReader(conn)
	for {
		word, _, err := r.ReadLine()
		if err != nil {
			fmt.Printf("Ошибка чтения из соединения: %v\n", err)
			return
		}

		fmt.Println("Пришло сообщение: ", string(word))

		err = sendUrls(allDocs, string(word), conn)
		if err != nil {
			fmt.Printf("Ошибка записи : %v\n", err)
			return
		}

		conn.SetDeadline(time.Now().Add(time.Second * 600))
	}
}

func scanDocs() ([]crawler.Document, error) {
	var allDocs []crawler.Document
	s := spider.New()
	for _, url := range urls {
		docs, err := s.Scan(url, 2)
		if err != nil {
			return allDocs, err
		}
		allDocs = append(allDocs, docs...)
	}
	return allDocs, nil
}

func sendUrls(docs []crawler.Document, word string, conn net.Conn) error {
	for _, doc := range docs {
		if strings.Contains(doc.URL, word) || strings.Contains(doc.Title, word) {
			msg := doc.URL + " " + doc.Title + "\n"
			_, err := conn.Write([]byte(msg))
			if err != nil {
				return err
			}
		}
	}
	_, err := conn.Write([]byte("--the end--\n"))
	return err
}

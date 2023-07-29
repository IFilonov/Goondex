// Package netsrv реализует telnet сервер для приема поисковых запросов.
package netsrv

import (
	"Goondex/part11/pkg/crawler/spider"
	"bufio"
	"fmt"
	"net"
	"time"
)

const addr = "0.0.0.0:8000"
const depth = 2

func StartNetSrv() error {
	listener, err := net.Listen("tcp4", addr)
	if err != nil {
		fmt.Printf("Ошибка открытия порта %v: %v\n", addr, err)
		return err
	}

	for {
		// Принимаем подключение.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Ошибка подключения: %v\n", err)
			continue
		}
		// Вызов обработчика подключения.
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	conn.SetDeadline(time.Now().Add(time.Second * 600))
	r := bufio.NewReader(conn)
	for {
		url, _, err := r.ReadLine()
		if err != nil {
			fmt.Printf("Ошибка чтения из соединения: %v\n", err)
			return
		}
		s := spider.New()
		u := string(url)
		fmt.Println("Пришло сообщение: ", u)
		docs, err := s.Scan(u, depth)
		if err != nil {
			fmt.Printf("Ошибка сканирования %v: %v\n", u, err)
			return
		}

		for i := 0; i < len(docs); i++ {
			msg := docs[i].URL + " " + docs[i].Title + "\n"
			_, err = conn.Write([]byte(msg))
			if err != nil {
				fmt.Printf("Ошибка записи : %v\n", err)
				return
			}
			fmt.Println("Передача обратно: ", msg)
		}
		_, err = conn.Write([]byte("--the end--\n"))

		conn.SetDeadline(time.Now().Add(time.Second * 600))
	}
}

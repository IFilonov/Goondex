package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp4", "0.0.0.0:8000")
	if err != nil {
		fmt.Printf("Ошибка открытия соединения: %v", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		msg, _ := reader.ReadString('\n')

		_, err = conn.Write([]byte(msg))
		if err != nil {
			fmt.Printf("Ошибка передачи данных на сервер: %v", err)
			continue
		}

		r := bufio.NewReader(conn)
		for {
			line, _, err := r.ReadLine()
			if err != nil {
				return
			}
			msg := string(line)
			if msg == "--the end--" {
				break
			}
			fmt.Println(msg)
		}
	}
}

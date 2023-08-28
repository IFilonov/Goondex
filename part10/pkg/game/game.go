// Package netsrv реализует telnet сервер для приема поисковых запросов.
package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var pings = map[string]string{"ping": "pong", "pong": "ping", "begin": "ping"}

type Score struct {
	player1Score int
	player2Score int
}

var score Score

func Run() {
	fmt.Println("Type Ctrl-C to exit")
	for {
		ch := make(chan string)
		var wg sync.WaitGroup
		wg.Add(2)

		go play_tennis(ch, 1, &wg)
		go play_tennis(ch, 2, &wg)
		ch <- "begin"

		wg.Wait()
		printScore()
	}
}

func play_tennis(ch chan string, playerNum int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		time.Sleep(time.Second) // Для замедленного отображения
		cmd, ok := <-ch
		fmt.Println("Player" + strconv.Itoa(playerNum) + ": " + cmd)
		if !ok {
			fmt.Println("Error read channel")
			return
		}
		if cmd == "stop" {
			close(ch)
			return
		}
		if playerWin(cmd, ch, playerNum) {
			return
		}
		ch <- pings[cmd]
	}
}

func playerWin(cmd string, ch chan string, playerNum int) bool {
	if (cmd == "ping" || cmd == "pong") && goalScored() {
		fmt.Println("Player" + strconv.Itoa(playerNum) + ": I win!")
		ch <- "stop"
		addScore(playerNum)
		return true
	}
	return false
}

func goalScored() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100) > 80
}

func addScore(playerNum int) {
	if playerNum == 1 {
		score.player1Score += 1
	} else {
		score.player2Score += 1
	}
}

func printScore() {
	fmt.Println("Score: Player1 " + strconv.Itoa(score.player1Score) +
		" points; Player2 " + strconv.Itoa(score.player2Score) + " points;")
}

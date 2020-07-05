package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	var words []string
	var input string
	var message string
	p1 := "Player 1 Wins!"
	fmt.Println("This is Player 2")

	conn, err := net.Dial("tcp", "localhost:5555")

	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		fmt.Println("Please start player1.go first")
		os.Exit(1)
	}

	fmt.Println("Connected!")

	fmt.Println("Game Start!")
	fmt.Println("Waiting for the first word:")
	for {

		message, _ = bufio.NewReader(conn).ReadString('\n')
		fmt.Print(string(message))

		if message == "Player 2 Wins!\n" {
			os.Exit(0)
		}

		if cusInput(&input) {

			msglow := strings.ToLower(message[0 : len(message)-1])
			inlow := strings.ToLower(input)
			words = append(words, msglow)
			_, found := find(words, inlow)

			if inlow[0:2] == msglow[len(msglow)-2:len(msglow)] || !found {
				words = append(words, inlow)
				conn.Write([]byte(input + "\n"))
			} else {
				fmt.Println("Invalid word!")
				fmt.Println(p1)
				conn.Write([]byte(p1 + "\n"))
				os.Exit(0)
			}
		} else {
			fmt.Println(p1)
			conn.Write([]byte(p1 + "\n"))
			os.Exit(0)
		}
	}
}

func cusInput(input *string) bool {

	c1 := make(chan string, 1)

	go func() {
		var str string
		fmt.Scanln(&str)
		c1 <- str
	}()

	select {
	case res := <-c1:
		*input = res
		return true
	case <-time.After(10 * time.Second):
		fmt.Println("Time out!")
		return false
	}

}

func find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

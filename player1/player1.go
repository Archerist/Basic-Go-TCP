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
	var firstentry bool = true
	var input string
	var message string
	var words []string

	p2 := "Player 2 Wins!"

	fmt.Println("This is Player 1")
	lis, err := net.Listen("tcp", "localhost:5555")

	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Waiting for Player 2")
	conn, _ := lis.Accept()

	fmt.Println("Player 2 Connected")

	defer lis.Close()

	for {

		if firstentry {
			fmt.Println("Game Start!")
			fmt.Println("Enter a word: ")
			if cusInput(&input) {
				words = append(words, strings.ToLower(input))
				conn.Write([]byte(input + "\n"))
			} else {
				fmt.Println(p2)
				conn.Write([]byte(p2 + "\n"))
				os.Exit(0)
			}

			firstentry = false
		} else {
			message, _ = bufio.NewReader(conn).ReadString('\n')
			fmt.Print(string(message))
			if message == "Player 1 Wins!\n" {
				os.Exit(0)
			}

			if cusInput(&input) {

				msglow := strings.ToLower(message[0 : len(message)-1])
				inlow := strings.ToLower(input)
				words = append(words, msglow)
				_, found := find(words, inlow)

				if inlow[0:2] == msglow[len(msglow)-3:len(msglow)-1] || !found {
					words = append(words, inlow)

					conn.Write([]byte(input + "\n"))
				} else {
					fmt.Println("Invalid word!")
					fmt.Println(p2)
					conn.Write([]byte(p2 + "\n"))
					os.Exit(0)
				}
			} else {
				fmt.Println(p2)
				conn.Write([]byte(p2 + "\n"))
				os.Exit(0)
			}
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

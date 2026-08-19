package main

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	scanner := bufio.NewScanner(os.Stdin)
	go func() {
		<-sigChan
		fmt.Println("Bye!")
		os.Exit(0)
	}()

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		cmd := scanner.Text()
		fmt.Printf("Got: %s\n", cmd)

		if err := scanner.Err(); err != nil {
			fmt.Fprint(os.Stderr, "Error Reading Stdin: ", err)
		}
	}
}

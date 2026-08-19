package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

func main() {
	rAddr := getAddr("255.255.255.255:5000")
	lAddr := getAddr("0.0.0.0:5000")
	readConn := receiveMsg(lAddr)
	write := connectUdp(rAddr)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		<-sigChan
		fmt.Println("Bye!")
		os.Exit(0)
	}()

	go func() {
		buf := make([]byte, 1024)
		for {
			n, addr, err := readConn.ReadFrom(buf)
			if err != nil {
				fmt.Print("read fail: ", err)
				continue
			}
			if write.LocalAddr().String() == addr.String() {
				continue
			}
			fmt.Printf("[%s]: %s\n", strings.Split(addr.String(), ":")[0], string(buf[:n]))
		}
	}()

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}

		cmd := scanner.Text()
		write.Write([]byte(cmd))

		if err := scanner.Err(); err != nil {
			fmt.Fprint(os.Stderr, "Error Reading Stdin: ", err)
		}
	}
}

func getAddr(addr string) *net.UDPAddr {
	udpAddr, err := net.ResolveUDPAddr("udp4", addr)
	if err != nil {
		panic("Failed to send message.")
	}
	return udpAddr
}

func connectUdp(udpAdrr *net.UDPAddr) *net.UDPConn {
	conn, err := net.DialUDP("udp4", nil, udpAdrr)
	if err != nil {
		panic("Failed to open connection.")
	}
	return conn
}

func receiveMsg(addr *net.UDPAddr) *net.UDPConn {
	conn, err := net.ListenUDP("udp4", addr)
	if err != nil {
		panic("Failed to open connection.")
	}
	return conn
}

package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// to launch
// go run task.go --timeout=100s stackoverflow.com 80 <---(or any other host + port)
//send GET request
/*
==Утилита telnet==

Реализовать простейший telnet-клиент.
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123
Требования:
1. Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
2. Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
3. При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера,
программа должна также завершаться. При подключении к несуществующему сервер, программа должна завершаться через timeout
*/
var (
	timeout time.Duration
)

func init() {
	flag.DurationVar(&timeout, "timeout", 100*time.Second, "connection timeout")
	flag.Parse()
}

// parse provided string to take host and port
func parseString() (string, error) {
	//get unflagged values. in args we store slice of string with parameters which are not marked as flag
	args := flag.Args()

	if len(args) < 2 || args[0] == "" || args[1] == "" {
		return "", errors.New("please enter host and port")
	}
	//host is a first value of unflagged arguments. port is a second one
	host := args[0]
	port := args[1]
	//address for connection
	address := host + ":" + port
	return address, nil
}

func main() {
	//parse provided string and construct address of socket
	address, err := parseString()
	if err != nil {
		fmt.Println(err)
		return
	}

	//establish connection with given timeout
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Println("error connecting to the server ", err)
		return
	}
	defer conn.Close()

	//read data from stdin and send to socket
	scanner := bufio.NewScanner(os.Stdin)
	//for is used as user can send several requests
	for {
		fmt.Print("> ")
		//requests can consist of several rows(lines), so concatenating strings into 'message' until encountering the '^]' character
		var message string
		for scanner.Scan() {
			fmt.Print("> ")
			line := scanner.Text()
			if line == "^]" {
				break
			}
			message += line + "\r\n"
		}
		fmt.Println(message)
		//send message to connection
		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Println("error writing to the server:", err)
			break
		}

		//read data from the socket and sent to stdout
		_, errCopy := io.Copy(os.Stdout, conn)
		if errCopy != nil {
			fmt.Println("Unable to copy from socket to stdout. Press ctrl+c to exit")
			os.Exit(0)
		}
		fmt.Printf("\n")
	}

	//catch ctrl+c signal to gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Program shutdown by signal")
}

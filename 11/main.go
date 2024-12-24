package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	timeoutFlag := flag.String("timeout", "10s", "Timeout for connection (e.g. 10s, 5m, 1h)")
	flag.Parse()

	if len(flag.Args()) != 2 {
		fmt.Println("Usage: go-telnet --timeout=<timeout> <host> <port>")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	timeout, err := time.ParseDuration(*timeoutFlag)
	if err != nil {
		fmt.Printf("Invalid timeout format: %v\n", err)
		os.Exit(1)
	}

	address := fmt.Sprintf("%s:%s", host, port)

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		fmt.Printf("Failed to connect to %s: %v\n", address, err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Printf("Connected to %s\n", address)

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if err != nil {
				fmt.Println("Connection closed by server")
				os.Exit(0)
			}
			fmt.Print(string(buf[:n]))
		}
	}()

	// Горутина для отправки данных в сокет
	buf := make([]byte, 1024)
	for {
		// Чтение данных с stdin (ввод пользователя)
		n, err := os.Stdin.Read(buf)
		if err != nil {
			// Если ошибка при чтении (например, Ctrl+D)
			fmt.Println("Error reading input:", err)
			break
		}
		// Отправка данных на сервер
		_, err = conn.Write(buf[:n])
		if err != nil {
			fmt.Println("Error sending data:", err)
			break
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("start error :", err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("connect error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("client:", conn.RemoteAddr().String())

	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read error:", err)
		return
	}

	message = strings.TrimSpace(message)
	fmt.Println("Получено:", message)

var response string

if message == "/help" {
    response = "Commands: /help, /time\n"
} else if message == "/time" {
    response = time.Now().Format("15:04:05") + "\n"
} else {
    response = "Echo: " + message + "\n"
}

	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("response error ", err)
		return
	}
}

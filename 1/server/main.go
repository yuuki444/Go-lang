package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("start error:", err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("connection error:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Клиент:", conn.RemoteAddr().String())

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read error :", err)
		return
	}

	message := string(buf[:n])
	fmt.Println("Получено:", message)

	var response string
	if message == "/help\n" || message == "/help" {
		response = "Commands: /help, /time\n"
	} else if message == "/time\n" || message == "/time" {
		response = time.Now().Format("15:04:05") + "\n"
	} else {
		response = "Echo: " + message
	}

	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("response error :", err)
		return
	}
}

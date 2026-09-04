package main

import (
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("connection error :", err)
		return
	}
	defer conn.Close()

	var message string
	fmt.Print("Введите сообщение: ")
	_, err = fmt.Scan(&message)
	if err != nil {
		fmt.Println("error ", err)
		return
	}

	_, err = conn.Write([]byte(message + "\n"))
	if err != nil {
		fmt.Println("message error :", err)
		return
	}

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("read error : ", err)
		return
	}

	response := string(buf[:n])
	fmt.Print("Server: ", response)
}

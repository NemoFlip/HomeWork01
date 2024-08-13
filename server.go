package main

import (
	"fmt"
	"syscall"
)

func main() {
	serverSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0) // AF_INET - IPv4 proto,
	if err != nil {
		fmt.Printf("Error creating a socket: %s", err.Error())
	}
	sockAddr := &syscall.SockaddrInet4{
		Port: 8080,
		Addr: [4]byte{127, 0, 0, 1},
	}
	err = syscall.Bind(serverSocket, sockAddr)
	if err != nil {
		fmt.Printf("Error binding the socket to the adress: %s", err.Error())
		return
	}

	err = syscall.Listen(serverSocket, 5)
	if err != nil {
		fmt.Printf("Error listening: %s", err.Error())
		return
	}
	fmt.Println("Server is listening...")

	clientSocket, clientAddress, err := syscall.Accept(serverSocket)
	if err != nil {
		fmt.Printf("Error accepting: %s", err.Error())
		return
	} else {
		fmt.Printf("Установлено соединение с %s \n", clientAddress)
	}

	message := []byte("Ok")
	_, err = syscall.Write(clientSocket, message)
	if err != nil {
		fmt.Printf("Error writing the message: %s", err.Error())
		return
	}

}

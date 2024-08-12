package main

import (
	"fmt"
	"syscall"
)

func main() {
	clientSocket, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		fmt.Printf("Error creating socket: %s", err.Error())
		return
	}
	socketAddress := &syscall.SockaddrInet4{
		Port: 8070,
		Addr: [4]byte{127, 0, 0, 1},
	}
	err = syscall.Connect(clientSocket, socketAddress)
	if err != nil {
		fmt.Printf("Error connecting through socket: %s", err.Error())
		return
	}
	var respMessage []byte
	_, _, err = syscall.Recvfrom(clientSocket, respMessage, 0)
	if err != nil {
		fmt.Printf("Error receiving the message: %s", err.Error())
		return
	}
	fmt.Printf("Response message: %s", string(respMessage))

}

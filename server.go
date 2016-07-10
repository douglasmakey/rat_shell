package main

import "net"
import "fmt"
import "bufio"
import (
	"strings"
	"os/exec"
)

func main() {

	fmt.Println("Launching server...")

	// listen on all interfaces
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()

	// run loop forever
	for {
		// will listen for command to process
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output command received
		fmt.Print("Command Received:", string(message))

		//Test execute command
		parts := strings.Fields(message)
		head := parts[0]
		parts = parts[1:len(parts)]

		out, err := exec.Command(head,parts...).Output()

		if err != nil {
			fmt.Printf("%s", err)
		}

		fmt.Printf("%s", out)

		// send out back to client
		newMessage := strings.ToUpper(message)
		conn.Write([]byte(newMessage + "\n"))

	}
}
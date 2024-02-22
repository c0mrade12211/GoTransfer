package servers

import (
	"fmt"
	"io"
	"net"
	"os"
)

func TCPserver(port int) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer ln.Close()

	fmt.Printf("Server started. Listening on port %d...\n", port)

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Client connected.")

	file, err := os.Create("received_file")
	if err != nil {
		fmt.Println("Error creating file: ", err.Error())
		os.Exit(1)
	}
	defer file.Close()

	_, err = io.Copy(file, conn)
	if err != nil {
		fmt.Println("Error copying file: ", err.Error())
	}
	fmt.Println("File received and saved.")
}

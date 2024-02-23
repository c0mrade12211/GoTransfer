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
		return
	}
	defer ln.Close()
	fmt.Printf("Server started. Listening on port %d...\n", port)

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		fmt.Println("Client connected.")

		file, err := os.Create("received_file")
		if err != nil {
			fmt.Println("Error creating file: ", err.Error())
			conn.Close()
			continue
		}

		_, err = io.Copy(file, conn)
		if err != nil {
			fmt.Println("Error copying file: ", err.Error())
		} else {
			fmt.Println("File received and saved.")
		}

		file.Close()
		conn.Close()
	}
}

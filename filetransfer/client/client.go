package client

import (
	"fmt"
	"io"
	"net"
	"os"
)

func TCPclient_upload(serverIP string, port int, filePath string) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, port))
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		return
	}
	defer conn.Close()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err.Error())
		return
	}
	defer file.Close()
	_, err = conn.Write([]byte(filePath + "\n"))
	if err != nil {
		fmt.Println("Error sending file name:", err.Error())
		return
	}
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error sending file content:", err.Error())
		return
	}

	fmt.Println("File sent successfully.")
}

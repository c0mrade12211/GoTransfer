package main

import (
	client "filetransfer/client"
	"filetransfer/models"
	servers "filetransfer/server"

	"flag"
	"fmt"
)

func main() {
	var FileTransferData models.FileTransferData

	flag.IntVar(&FileTransferData.Port, "port", 8081, "Specify port number. Default is 8081")
	flag.StringVar(&FileTransferData.Protocol, "protoc", "tcp", "Specify protocol . Default is tcp")
	flag.StringVar(&FileTransferData.ClientOrServer, "i", "", "Specify client or server. Default is server")
	flag.StringVar(&FileTransferData.PathToGet, "get", "", "Specify path to get. ")
	flag.StringVar(&FileTransferData.PathToUploadFile, "upload", "", "Specify path to upload file.")
	flag.StringVar(&FileTransferData.ServerIP, "server_ip", "", "Specify server ip. ")
	flag.Parse()

	if FileTransferData.Protocol == "tcp" && FileTransferData.ClientOrServer == "server" {
		fmt.Println("You're set to a TCP server")
		servers.TCPserver(FileTransferData.Port)
	}
	if FileTransferData.ClientOrServer == "client" && FileTransferData.Protocol == "tcp" && FileTransferData.PathToUploadFile != "" && FileTransferData.ServerIP != "" && FileTransferData.Port != 0 {
		client.TCPclient_upload(FileTransferData.ServerIP, FileTransferData.Port, FileTransferData.PathToUploadFile)
	}
	if FileTransferData.ClientOrServer == "server" && FileTransferData.Protocol == "http" {
		servers.HTTPserver(FileTransferData.Port)
	} else {
		flag.PrintDefaults()
	}

}

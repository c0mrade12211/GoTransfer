
---

# File Transfer Application

This is a simple file transfer application written in Go that allows you to transfer files between a client and a server using TCP protocol.

## Usage

### Server
To run the server, use the following command:
go run main.go -i server -protoc tcp -port <port_number>
-  `-i server` : Specifies that the application will run as a server.
-  `-protoc tcp` : Specifies the protocol to be used (default is TCP).
-  `-port <port_number>` : Specifies the port number to listen on.

### Client
To run the client and upload a file to the server, use the following command:
go run main.go -i client -protoc tcp -port <port_number> -upload <file_path> -server_ip <server_ip>
-  `-i client` : Specifies that the application will run as a client.
-  `-protoc tcp` : Specifies the protocol to be used (default is TCP).
-  `-port <port_number>` : Specifies the port number to connect to.
-  `-upload <file_path>` : Specifies the path to the file to upload.
-  `-server_ip <server_ip>` : Specifies the IP address of the server.

If any of the required parameters are missing, the application will display the default settings and usage instructions.

## Dependencies
- Go programming language

## Contributors
- [Your Name](https://github.com/c0mrade12211)

Feel free to contribute to this project by submitting issues or pull requests.

---


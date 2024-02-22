package models

type FileTransferData struct {
	ClientOrServer   string
	PathToGet        string
	PathToUploadFile string
	ServerIP         string
	Port             int
	Protocol         string
}

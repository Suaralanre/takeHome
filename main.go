package main

import (
	"net/http"
	"takeHome/filehandler"
)

func main() {

	// File upload endpoints
	http.HandleFunc("/upload", filehandler.UploadFileHandler)
	http.HandleFunc("/files", filehandler.ListFilesHandler)
	http.HandleFunc("/download", filehandler.DownloadFileHandler)

	// Serve HTML templates
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates"))))

	// Start the server
	http.ListenAndServe(":8080", nil)
}

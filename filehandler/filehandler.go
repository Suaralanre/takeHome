package filehandler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const uploadDir = "files/"

// UploadFileHandler handles file uploads
func UploadFileHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(15 << 20) // 15 MB limit

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error Retrieving the File "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := filepath.Join(uploadDir, handler.Filename)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Error Creating the File ", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error Copying the File", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully Uploaded File: %s\n", handler.Filename)
}

// ListFilesHandler returns a list of uploaded files
func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := getFiles(uploadDir)
	if err != nil {
		http.Error(w, "Error Listing Files", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "List of Uploaded Files:")
	for _, file := range files {
		fmt.Fprintln(w, file)
	}
}

// DownloadFileHandler handles file downloads
func DownloadFileHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("filename")
	if fileName == "" {
		http.Error(w, "Missing 'filename' parameter", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join(uploadDir, fileName)
	http.ServeFile(w, r, filePath)
}

// getFiles returns a list of files in a directory
func getFiles(dirPath string) ([]string, error) {
	var files []string
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, strings.TrimPrefix(path, dirPath))
		}
		return nil
	})
	return files, err
}

package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// UploadFile handles multipart POST requests and saves the file to the local "uploads" directory.
func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form, max 10 MB limit
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Error parsing form or file too large", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file") // The frontend must use form key 'file'
	if err != nil {
		http.Error(w, "Error retrieving file from request. Ensure form field name is 'file'", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Ensure uploads directory exists
	uploadDir := "uploads"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		http.Error(w, "Failed to create upload directory", http.StatusInternalServerError)
		return
	}

	// Create a unique filename to avoid overwrites
	ext := filepath.Ext(handler.Filename)
	baseName := strings.TrimSuffix(handler.Filename, ext)
	newFilename := baseName + "_" + time.Now().Format("20060102150405") + ext
	filePath := filepath.Join(uploadDir, newFilename)

	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to create file on server", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// Copy file contents
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	fileUrl := "/uploads/" + newFilename

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "File uploaded successfully",
		"url":     fileUrl,
	})
}

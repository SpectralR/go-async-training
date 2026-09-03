package job

import (
	"fmt"
	"go-async-training-api/internal/custom_error"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const uploadDir = "/app/uploads"

func Upload(response http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		http.Error(response, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// upload
	parseErr := request.ParseMultipartForm(10 << 20)
	if parseErr != nil {
		http.Error(response, fmt.Sprintf("Error file too large: %s", parseErr), http.StatusRequestEntityTooLarge)
		return
	}

	file, handler, err := request.FormFile("file")
	if err != nil {
		http.Error(response, fmt.Sprintf("Error Retrieving the File: %s", err), http.StatusUnprocessableEntity)
		return
	}

	defer file.Close()

	if !strings.Contains(handler.Header["Content-Type"][0], "text/csv") {
		http.Error(response, "Error Wrong file Format", http.StatusUnprocessableEntity)
		return
	}

	fileError := createFile(handler.Filename, file)
	if fileError != nil {
		http.Error(response, fileError.Error(), http.StatusInternalServerError)
	}

	// queue
}

func createFile(filename string, file multipart.File) error {
	locFile, err := os.Create(filepath.Join(uploadDir, filename))
	if err != nil {
		return &custom_error.FileCreationError{Message: err.Error()}
	}

	defer locFile.Close()

	content, readErr := io.ReadAll(file)

	if readErr != nil {
		return &custom_error.FileReadError{Message: readErr.Error()}
	}

	_, writeErr := locFile.Write(content)
	if writeErr != nil {
		return &custom_error.FileWriteError{Message: writeErr.Error()}
	}

	return nil
}

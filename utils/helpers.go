package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// FileSave fetches the file and saves to disk
func FileSave(r *http.Request) (string, string) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return "", ""
	}
	n := r.Form.Get("email")
	// Retrieve the file from form data
	f, h, err := r.FormFile("file")
	if err != nil {
		return "", ""
	}
	name := strings.Split(h.Filename, ".")

	defer f.Close()
	path := filepath.Join(".", "files")
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := path + "/" + n + filepath.Ext(h.Filename)
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return "", ""
	}
	defer file.Close()
	// Copy the file to the destination path
	_, err = io.Copy(file, f)
	if err != nil {
		return "", ""
	}
	return fullPath, name[1]
}

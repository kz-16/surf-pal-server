package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	http.HandleFunc("/storage/images/", func(w http.ResponseWriter, r *http.Request) {
		// Get the image filename from the URL path
		imageName := filepath.Base(r.URL.Path)

		_, filename, _, ok := runtime.Caller(0)
		if !ok {
			http.Error(w, "err.Error()", http.StatusInternalServerError)
			return
		}
		dir := filepath.Dir(filename)
		absPath, err := filepath.Abs(dir)
		if err != nil {
			http.Error(w, "000000", http.StatusInternalServerError)
			return
		}

		// Construct the full path to the image file
		imagePath := filepath.Join(absPath, "surf_pal_images", imageName)

		// Open the image file
		imgFile, err := os.Open(imagePath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer imgFile.Close()

		// Set the Content-Type header based on the image file extension
		ext := filepath.Ext(imageName)
		switch ext {
		case ".jpg", ".jpeg":
			w.Header().Set("Content-Type", "image/jpeg")
		case ".png":
			w.Header().Set("Content-Type", "image/png")
		default:
			w.Header().Set("Content-Type", "application/octet-stream")
		}

		// Copy the image file to the response writer
		io.Copy(w, imgFile)
	})

	http.ListenAndServe(":8080", nil)
}

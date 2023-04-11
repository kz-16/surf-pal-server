package main

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func main() {
	http.HandleFunc("/storage/images/", func(w http.ResponseWriter, r *http.Request) {
		// Get the image filename from the URL path
		imageName := filepath.Base(r.URL.Path)

		// Construct the full path to the image file
		imagePath := filepath.Join("surf_pal_country_icons", imageName)

		// Open the image file
		imgFile, err := os.Open(imagePath)
		if err != nil {
			http.Error(w, "Image not found", http.StatusInternalServerError)
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

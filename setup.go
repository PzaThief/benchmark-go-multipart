package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// generate file to transfer
	files := map[string]int64{
		"test-100MB": 100 * 1024 * 1024,
		"test-10MB": 10 * 1024 * 1024,
		"test-1MB": 1 * 1024 * 1024,
		"test-100KB": 100 * 1024,
		"test-10KB": 10 * 1024,
		"test-1KB": 1 * 1024,
	}
	for fileName, size := range files {
		f, _ := os.Create(fileName)
		f.Truncate(size)   
	}

	// run http server
	mux := http.NewServeMux()
	mux.HandleFunc("/", uploadHandler)
	http.ListenAndServe(":8080", nil)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)
}

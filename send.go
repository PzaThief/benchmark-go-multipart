package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"acln.ro/zerocopy"
)

func osPipeSend(fileName, url string) {
	r, w, _ := os.Pipe()
	writer := multipart.NewWriter(w)

	file, _ := os.Open(fileName)
	defer file.Close()
	go func() {
		writer.CreateFormFile("file", filepath.Base(file.Name()))
		w.ReadFrom(file)
		writer.Close()
		w.Close()
	}()

	http.Post(url, writer.FormDataContentType(), r)
}

type readCloserPipeWrap struct {
	*zerocopy.Pipe
}

func (p *readCloserPipeWrap) Close() error {
	return p.CloseRead()
}

func zeroCopyLibSend(fileName, url string) {
	p, _ := zerocopy.NewPipe()
	rw := &readCloserPipeWrap{p}
	writer := multipart.NewWriter(rw)

	file, _ := os.Open(fileName)
	defer file.Close()
	go func() {
		writer.CreateFormFile("file", filepath.Base(file.Name()))
		rw.ReadFrom(file)
		writer.Close()
		rw.CloseWrite()
	}()

	http.Post(url, writer.FormDataContentType(), rw)
}

func ioPipeSend(fileName, url string) {
	r, w := io.Pipe()
	writer := multipart.NewWriter(w)

	go func() {
		defer w.Close()
		file, _ := os.Open(fileName)
		defer file.Close()
		part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
		io.Copy(part, file)
		writer.Close()
	}()

	http.Post(url, writer.FormDataContentType(), r)
}

func naiveSend(fileName, url string) {
	file, _ := os.Open(fileName)
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)
	writer.Close()

	http.Post(url, writer.FormDataContentType(), body)
}

package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/weAutomateEverything/pdf2dash/extractPages"
)

func main() {
	err := extractPages.ExtractImages("pdfFile.pdf")
	if err != nil {
		log.Printf("PDF page extraction failed with the following error: %v",err)
	}

	fs := http.FileServer(http.Dir("./staticPages"))

	// Create HTTP server
	http.Handle("/", fs)
	http.HandleFunc("/uploadFile", upload)

	http.Handle("/files/", http.StripPrefix("/files", fs))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

// Upload new pdf file
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		os.Rename(f.Name(), "pdfFile.pdf")

		extractPages.ExtractImages("pdfFile.pdf")
	}
}

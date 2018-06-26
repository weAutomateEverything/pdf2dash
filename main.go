package main

import (
	"log"
	"net/http"
	"pdf2HTML/extractPages"
	"fmt"
	"time"
	"crypto/md5"
	"io"
	"strconv"
	"os"
	"html/template"
)

func main(){
	extractPages.ExtractImages("pdfFile.pdf")

	fs := http.FileServer(http.Dir("./staticPages"))

	http.Handle("/", fs)
	http.HandleFunc("/uploadFile", upload)

	http.Handle("/files/", http.StripPrefix("/files", fs))

	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)
}

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
		os.Rename(f.Name(),"pdfFile.pdf")

		extractPages.ExtractImages("pdfFile.pdf")
	}
}

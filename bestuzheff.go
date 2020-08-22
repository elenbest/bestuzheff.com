package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func indexpage(w http.ResponseWriter, r *http.Request) {

	pageBody, err := ioutil.ReadFile("index.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	fmt.Fprintf(w, string(pageBody))
}

func main() {
	http.HandleFunc("/", indexpage)
	err := http.ListenAndServe(":80", nil) // устанавливаем порт веб-сервера
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

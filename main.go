package main

import (
	"log"
	"net/http"
)

var downloadLink string = "https://ftp.drupal.org/files/projects/drupal-8.1.5.tar.gz"

func httpHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, downloadLink, http.StatusFound)

}

func main() {

	http.HandleFunc("/", httpHandler)
	http.Handle("/static", http.FileServer(http.Dir("./static")))

	log.Fatal(http.ListenAndServe(":8080", nil))

}

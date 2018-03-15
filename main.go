package main

import (
	"fmt"
	"net/http"
)

func deepLink(w http.ResponseWriter, r *http.Request) {
	var html string
	html += "<h1 style=\"text-align:center;\">"
	html += "<a href=\"%s\">Open app</a><br>"
	html += "<a href=\"http://192.168.7.76:8080/promotion?id=c20ad4d76fe97759aa27a0c99bff6710\">My Promotion</a>"
	html += "</h1>"

	fmt.Fprintf(w, html, "myapp://hostname")
}

func main() {
	http.Handle("/.well-known/", http.StripPrefix("/.well-known/", http.FileServer(http.Dir("./public/.well-known/"))))
	http.HandleFunc("/", deepLink)
	http.ListenAndServe(":8080", nil)
}

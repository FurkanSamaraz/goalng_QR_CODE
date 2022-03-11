package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	c := template.Must(template.ParseFiles("index.html"))
	c.Execute(w, req.FormValue("s"))
}

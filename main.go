package main

import (
	"fmt"
	"net/http"
	"text/template"

	"frontendmasters.com/go/femm/api"
	"frontendmasters.com/go/femm/data"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from a Go Program"))
}
func handleTemplate(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("templates/index.templ")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		// The header should be write before sending back response
		// w.WriteHeader(http.StatusInternalServerError)
		return
	}
	html.Execute(w, data.GetAll())
}

func main() {
	server := http.NewServeMux()
	server.HandleFunc("/hello", helloHandler)
	server.HandleFunc("/template", handleTemplate)
	server.HandleFunc("/api/getExhibitions", api.Get)
	server.HandleFunc("/api/exhibitions/new", api.Post)
	fs := http.FileServer(http.Dir("./public"))
	server.Handle("/", fs)
	err := http.ListenAndServe(":3333", server)
	if err == nil {
		fmt.Println("Error while opening the server")
	}
}

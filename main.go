package main

import (
	"html/template"
	"net/http"
)

func main(){
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/home/", mainViewHandler)
	http.ListenAndServe(":8080", nil)
}

func mainViewHandler(responseWriter http.ResponseWriter,request *http.Request){

	t, _ := template.ParseFiles("main.html")
	t.Execute(responseWriter, 0)
}
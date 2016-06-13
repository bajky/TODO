package main

import (
	"html/template"
	"net/http"
	"fmt"
	"strings"
	"strconv"
)


type Todo struct{
	Title string
	Description string
}

type EditPage struct{
	Todos map[int]Todo
	Text string
}

var todos map[int] Todo = make(map[int] Todo)

func main(){
	http.Handle("/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/edit/", editViewHandler)
	http.HandleFunc("/add/", addViewHandler)
	http.HandleFunc("/list/", listViewHandler)
	http.HandleFunc("/save", saveHandler)

	http.ListenAndServe(":8080", nil)
}

func addViewHandler(responseWriter http.ResponseWriter,request *http.Request){

	t, _ := template.ParseFiles("add.html")
	t.Execute(responseWriter, 0)
}

func editViewHandler(responseWriter http.ResponseWriter,request *http.Request){
	editPage := EditPage{}
	editPage.Todos = todos

	//striaaang := string(request.URL.Path[len("/edit"):])
	//fmt.Println(striaaang)

	if strings.Compare(string(request.URL.Path[len("/edit"):]), "/") != 0{
		title := strings.Fields(request.FormValue("title"))
		if len(title) != 0{
			id, err:= strconv.Atoi(title[0])
			if err != nil {

				fmt.Println(err)
				return
			}

			editPage.Text = todos[id].Description
		}

	}else{
		editPage.Text = ""
	}
	t, _ := template.ParseFiles("edit.html")
	t.Execute(responseWriter, editPage)
}

func listViewHandler(responseWriter http.ResponseWriter,request *http.Request){

	fmt.Println(request.Body)
	t, _ := template.ParseFiles("list.html")
	t.Execute(responseWriter, todos)
}



func saveHandler(responseWriter http.ResponseWriter,request *http.Request){

	todo := Todo{
		Title: request.FormValue("title"),
		Description: request.FormValue("description"),
	}
	save(todo)

	t, _ := template.ParseFiles("list.html")
	t.Execute(responseWriter, todos)
}
func save(todo Todo) {
	todos[len(todos)] = todo
}
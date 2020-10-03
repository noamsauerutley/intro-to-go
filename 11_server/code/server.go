package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// PageVariables are variables sent to the browser
type PageVariables struct {
	PageTitle string
	PageTodos []Todo
}

// Todo is a task summary. It must have a Title & Content
type Todo struct {
	Title   string
	Content string
}

var todos []Todo

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home!")
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	pageVariables := PageVariables{
		PageTitle: "Get Todos",
		PageTodos: todos,
	}

	t, err := template.ParseFiles("todos.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error", err)
	}

	err = t.Execute(w, pageVariables)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Requet parsing error", err)
	}

	todo := Todo{
		Title:   r.FormValue("title"),
		Content: r.FormValue("content"),
	}

	todos = append(todos, todo)
	log.Print(todos)
	http.Redirect(w, r, "/todos/", http.StatusSeeOther)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/todos/", getTodos)
	http.HandleFunc("/add-todo/", addTodo)
	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

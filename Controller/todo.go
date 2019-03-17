package todo

import (
	"html/template"
	"net/http"
	"strconv"
	model "todo-app/Model"

	"github.com/julienschmidt/httprouter"
)

// Index is handle
func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var allTodos []model.Todo
	model.Db.Find(&allTodos)
	t, _ := template.ParseFiles("./View/index.html")
	t.Execute(w, allTodos)
}

// Show is handle
func Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./View/show.html")
	todo := model.Todo{}
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Redirect(w, r, "/todo/index", 303)
	}
	model.Db.Where("id = $1", id).First(&todo)
	t.Execute(w, todo)
}

// Edit is handle
func Edit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	todo := model.Todo{}
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Redirect(w, r, "/todo/index", 303)
	}
	model.Db.Model(&todo).Where("ID = ?", id).Update("Content", r.PostFormValue("content"))
	http.Redirect(w, r, "/todo/index/", 303)
}

// Create is handle
func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todo := model.Todo{Content: r.PostFormValue("content")}
	model.Db.Create(&todo)
	http.Redirect(w, r, "/todo/index", 303)
}

// Delete is handle
func Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	todo := model.Todo{}
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.RedirectHandler("/todo/index", 303)
	}
	todo.ID = id
	model.Db.First(&todo)
	model.Db.Delete(&todo)
	http.Redirect(w, r, "/todo/index", 303)
}

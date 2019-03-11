package Todo

import (
	"github.com/julienschmidt/httprouter"
	"html/template"
	"net/http"
	"strconv"
	"todo-app/Model"
)

func Index(w http.ResponseWriter, _ *http.Request, _ httprouter.Params) {
	var allTodos []Model.Todo
	Model.Db.Find(&allTodos)
	t, _ := template.ParseFiles("./View/index.html")
	t.Execute(w, allTodos)
}

func Show(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, _ := template.ParseFiles("./View/show.html")
	todo := Model.Todo{}
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Redirect(w, r, "/todo/index", 303)
	}
	Model.Db.Where("id = $1", id).First(&todo)
	t.Execute(w, todo)
}

func Edit(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	todo := Model.Todo{}
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.Redirect(w, r, "/todo/index", 303)
	}
	Model.Db.Model(&todo).Where("Id = ?", id).Update("Content", r.PostFormValue("content"))
	http.Redirect(w, r, "/todo/show/"+p.ByName("id"), 303)
}

func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	todo := Model.Todo{Content: r.PostFormValue("content")}
	Model.Db.Create(&todo)
	http.Redirect(w, r, "/todo/index", 303)
}

func Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	todo := Model.Todo{}
	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		http.RedirectHandler("/todo/index", 303)
	}
	todo.Id = id
	Model.Db.First(&todo)
	Model.Db.Delete(&todo)
	http.Redirect(w, r, "/todo/index", 303)
}

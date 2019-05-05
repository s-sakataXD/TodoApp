package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	todo "github.com/s-sakataXD/TodoApp/Controller/todo"
)

func routes() http.Handler {
	mux := httprouter.New()
	mux.GET("/todo/index", todo.Index)
	mux.GET("/todo/show/:id", todo.Show)
	mux.POST("/todo/edit/:id", todo.Edit)
	mux.POST("/todo/create/", todo.Create)
	mux.POST("/todo/delete/:id", todo.Delete)
	return mux
}

func main() {
	mux := routes()
	http.ListenAndServe(":9000", mux)
}

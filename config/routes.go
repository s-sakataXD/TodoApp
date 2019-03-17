package main

import (
	"net"
	"net/http"
	"net/http/fcgi"
	todo "todo-app/controller"

	"github.com/julienschmidt/httprouter"
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
	l, _ := net.Listen("tcp", "127.0.0.1:9000")
	mux := routes()
	fcgi.Serve(l, mux)
}

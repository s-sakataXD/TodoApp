package main

import (
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"net/http/fcgi"
	"todo-app/Controller"
)

func routes() http.Handler {
	mux := httprouter.New()
	mux.GET("/todo/index", Todo.Index)
	mux.GET("/todo/show/:id", Todo.Show)
	mux.POST("/todo/edit/:id", Todo.Edit)
	mux.POST("/todo/create/", Todo.Create)
	mux.POST("/todo/delete/:id", Todo.Delete)
	return mux
}

func main() {
	l, err := net.Listen("tcp", "127.0.0.1:9000")
	if err != nil {
		return
	}
	mux := routes()
	fcgi.Serve(l, mux)
}

package routes

import (
	"fmt"
	"html/template"
	"strconv"
	"trash-stack-todo/model"

	"github.com/savsgio/atreugo/v11"
)

func getTodosHandler(rc *atreugo.RequestCtx) error {
	t := new(model.Todo)
	todos, err := t.GetAll()
	if err != nil {
		fmt.Println("Could not get all todos from db", err)
		return err
	}
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	err = tmpl.Execute(rc, todos)
	if err != nil {
		fmt.Println("Could not execute template", err)
		return err
	}
	rc.Response.Header.SetContentType("text/html")
	return nil
}

func createTodoHandler(rc *atreugo.RequestCtx) error {
	todo := rc.PostArgs().Peek("todo")
	done := rc.PostArgs().Peek("done")

	t := &model.Todo{
		Todo: string(todo),
		Done: string(done) == "true",
	}
	err := t.Create()
	if err != nil {
		fmt.Println("Could not create todo", err)
		return rc.TextResponse("500 Internal Server Error", 500)
	}
	return getTodosHandler(rc)
}

func markTodoHandler(rc *atreugo.RequestCtx) error {
	idArg := rc.UserValue("id").(string)
	id, err := strconv.ParseUint(idArg, 10, 64)
	if err != nil {
		fmt.Println("Could not parse id", err)
	}
	t := &model.Todo{
		Id: id,
	}
	err = t.MarkDone()
	if err != nil {
		fmt.Println("Could not update todo", err)
		return rc.TextResponse("500 Internal Server Error", 500)
	}
	return getTodosHandler(rc)
}

func SetupAndRun() {
	config := atreugo.Config{
		Addr: "0.0.0.0:8000",
	}
	server := atreugo.New(config)
	server.GET("/", getTodosHandler)
	server.PUT("/todo/{id}", markTodoHandler)
	//server.DELETE("/todo/{id}", DeleteTodoHandler)
	server.POST("/create", createTodoHandler)
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

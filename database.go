package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	DB_USER    = "fionwan"
	DB_NAME    = "todosDB"
	TABLE_NAME = "todos"
)

var db *sql.DB

func InitDB() {
	var err error

	dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable", DB_USER, DB_NAME)

	db, err = sql.Open("postgres", dbinfo)
	checkErr(err)
	fmt.Println("DB connection should be live")
}

func GetTodoList() (todos Todos) {
	var temp Todo

	rows, err := db.Query("select * from todos;")
	checkErr(err)

	for rows.Next() {
		rows.Scan(&temp.Id, &temp.Name, &temp.Completed)
		todos = append(todos, temp)
	}
	return todos
}

func GetTodoItem(id string) (todo Todo) {
	err := db.QueryRow("select * from todos where id = $1;", id).Scan(&todo.Id, &todo.Name, &todo.Completed)
	if err != nil {
		fmt.Println(err)
		todo.Id = -1
	}
	return todo
}

func InsertTodoItem(todo Todo) (added Todo) {
	if db == nil {
		err := errors.New("attempted to insert with no DB connection")
		LogError(err)

		todo.Id = -1
		return todo
	} else {
		err := db.QueryRow("insert into todos (name, completed) values ($1, $2) returning id;",
			todo.Name, todo.Completed).Scan(&todo.Id)
		if err != nil {
			todo.Id = -1
			LogError(err)
		}
		fmt.Println("last inserted id: ", todo.Id)
		return todo
	}
}

func MarkDone(id string) (todo Todo) {
	err := db.QueryRow("update todos set complete=true where id = $1 returning *;", id).Scan(
		&todo.Id, &todo.Name, &todo.Completed)
	if err != nil {
		todo.Id = -1
		LogError(err)
	}
	return todo
}

func DeleteTodoItem(id string) (deleted int) {
	err := db.QueryRow("delete from todos where id = $1 returning id;", id).Scan(&deleted)
	if err != nil {
		LogError(err)
		return -1
	}

	return deleted
}

func checkErr(err error) {
	if err != nil {
		LogError(err)
		panic(err)
	}
}

func LogError(err error) {
	fmt.Println(err.Error())
}

// func TodoDelete(w http.ResponseWriter, r *http.Request) {
// 	vars:= mux.Vars(r)
// 	todoId := vars[“todoId”]
// 	res := DeleteTodo(todoId)
// 	if res == -1 {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	} else {
// 		w.WriteHeader(http.StatusOK)
// 	}
// }
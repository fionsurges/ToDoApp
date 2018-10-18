package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	TABLE_NAME = "todos"
)

type DbConnection struct {
	DbUser string
	DbName string
	DbPort string
	DbHost string
	DbPassword string
	Conn *sql.DB
}

func (db *DbConnection) InitDB() {
	var err error

	if db.DbUser == "" {
		db.DbUser = os.Getenv("DB_USER")
	}
	if db.DbName == "" {
		db.DbName = os.Getenv("DB_NAME")
	}
	if db.DbPort == "" {
		db.DbPort = os.Getenv("DB_PORT")
	}
	if db.DbHost == "" {
		db.DbHost = os.Getenv("DB_HOST")
	}
	if db.DbPassword == "" {
		db.DbPassword = os.Getenv("DB_PASSWORD")
	}

	dbinfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", db.DbHost, db.DbPort, db.DbUser, db.DbName, db.DbPassword)

	db.Conn, err = sql.Open("postgres", dbinfo)
	checkErr(err)
	fmt.Println("DB connection should be live")
}

func (db *DbConnection) GetTodoList() (todos Todos) {
	var temp Todo

	rows, err := db.Conn.Query("select * from todos;")
	checkErr(err)

	for rows.Next() {
		rows.Scan(&temp.Id, &temp.Name, &temp.Completed)
		todos = append(todos, temp)
	}
	return todos
}

func (db *DbConnection) GetTodoItem(id string) (todo Todo) {
	err := db.Conn.QueryRow("select * from todos where id = $1;", id).Scan(&todo.Id, &todo.Name, &todo.Completed)
	if err != nil {
		fmt.Println(err)
		todo.Id = -1
	}
	return todo
}

func (db *DbConnection) InsertTodoItem(todo Todo) (added Todo) {
	if db.Conn == nil {
		err := errors.New("attempted to insert with no DB connection")
		LogError(err)

		todo.Id = -1
		return todo
	} else {
		err := db.Conn.QueryRow("insert into todos (name, completed) values ($1, $2) returning id;",
			todo.Name, todo.Completed).Scan(&todo.Id)
		if err != nil {
			todo.Id = -1
			LogError(err)
		}
		fmt.Println("last inserted id: ", todo.Id)
		return todo
	}
}

func (db *DbConnection) MarkDone(id string) (todo Todo) {
	err := db.Conn.QueryRow("update todos set complete=true where id = $1 returning *;", id).Scan(
		&todo.Id, &todo.Name, &todo.Completed)
	if err != nil {
		todo.Id = -1
		LogError(err)
	}
	return todo
}

func (db *DbConnection) DeleteTodoItem(id string) (deleted int) {
	err := db.Conn.QueryRow("delete from todos where id = $1 returning id;", id).Scan(&deleted)
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

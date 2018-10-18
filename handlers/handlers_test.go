package handlers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/fionwan/todoApp/database"
)

func Test_TodoShow(t *testing.T) {
	request := httptest.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()

	testdb := database.DbConnection{
		DbUser: "fionwan",
		DbName: "todosDB",
		DbPort: "5432",
		DbHost: "localhost",
	}

	dbinfo := fmt.Sprintf("sslmode=disable host=%s port=%s user=%s dbname=%s password=''", testdb.DbHost, testdb.DbPort, testdb.DbUser, testdb.DbName)
	testdb.Conn, _ = sql.Open("postgres", dbinfo)

	testTodo := database.Todo{
		Name:      "test",
		Completed: false,
	}

	testdb.InsertTodoItem(testTodo)

	testapp := App{
		db: &testdb,
	}

	testapp.TodoShow(w, request)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	testdb.Conn.QueryRow("delete * from todos;")
}

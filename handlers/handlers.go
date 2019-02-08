package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/fionwan/todoApp/database"
	"github.com/gorilla/mux"
)

type App struct {
	db     *database.DbConnection
	Router *mux.Router
}

func NewApp(db *database.DbConnection) *App {
	a := App{
		db: db,
	}

	routes := Routes{
		Route{
			"Index",
			"GET",
			"/",
			Index,
		},
		Route{
			"TodoIndex",
			"GET",
			"/todos",
			a.TodoShow,
		},
		Route{
			"TodoShow",
			"GET",
			"/todos/{todoId}",
			a.TodoIndex,
		},
		Route{
			"TodoCreate",
			"POST",
			"/todos",
			a.TodoCreate,
		},
		Route{
			"TodoDelete",
			"DELETE",
			"/todos/{todoId}",
			a.TodoDelete,
		},
		Route{
			"TodoEdit",
			"PATCH",
			"/todos/{todoId}",
			a.TodoEdit,
		},
		Route{
			"TodoMarkDone",
			"PATCH",
			"/todos/{todoId}/done",
			a.TodoMarkDone,
		},
	}

	a.Router = NewRouter(routes)
	return &a
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func (a *App) TodoShow(w http.ResponseWriter, r *http.Request) {

	todos := a.db.GetTodoList()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		fmt.Println(err)
	}
}

func (a *App) TodoIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	res := a.db.GetTodoItem(todoId)
	if res.Id == -1 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(err)
		}
	}
}

func (a *App) TodoMarkDone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	res := a.db.MarkDone(todoId)
	if res.Id == -1 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(err)
		}

	}
}

func (a *App) TodoEdit(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	var todo database.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		database.LogError(err)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Println(err)
		}

		return
	}

	res := a.db.ChangeTodo(todo.Name, todoId)

	if res.Id == -1 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(res); err != nil {
			fmt.Println(err)
		}

	}
}

func (a *App) TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	res := a.db.DeleteTodoItem(todoId)
	if res == -1 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (a *App) TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo database.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		database.LogError(err)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			fmt.Println(err)
		}

		return
	}

	t := a.db.InsertTodoItem(todo)
	if t.Id == -1 {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		fmt.Println(err)
	}
}

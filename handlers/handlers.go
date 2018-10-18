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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
	todos := database.GetTodoList()

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	res := database.GetTodoItem(todoId)
	if res.Id == -1 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}
	}
}

func TodoMarkDone(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	res := database.MarkDone(todoId)
	if res.Id == -1 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(res); err != nil {
			panic(err)
		}

	}
}

func TodoDelete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]

	res := database.DeleteTodoItem(todoId)
	if res == -1 {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func TodoCreate(w http.ResponseWriter, r *http.Request) {
	var todo database.Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422)
		database.LogError(err)

		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}

		return
	}

	t := database.InsertTodoItem(todo)
	if t.Id == -1 {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

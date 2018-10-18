package handlers

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
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
		TodoShow,
	},
	Route{
		"TodoShow",
		"GET",
		"/todos/{todoId}",
		TodoIndex,
	},
	Route{
		"TodoCreate",
		"POST",
		"/todos",
		TodoCreate,
	},
	Route{
		"TodoDelete",
		"DELETE",
		"/todos/{todoId}",
		TodoDelete,
	},
	Route{
		"TodoMarkDone",
		"PUT",
		"/todos/{todoId}",
		TodoMarkDone,
	},
}

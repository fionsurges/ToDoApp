package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	// "github.com/fionwan/todoApp/database"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", TodoShow).Methods("GET")

	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

// func Router() *mux.Router {
// 	router := mux.NewRouter()
// 	router.HandleFunc("/", TodoShow).Methods("GET")
// 	return router
// }

func TestTodoShow(t *testing.T) {
	request, _ := http.NewRequest("GET", "/todos", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
}

package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing" 

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/todos", TodoShow).Methods("GET")
    return router
}


func TestTodoShow(t *testing.T) {

	request, _ := http.NewRequest("GET", "/todos", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "OK response is expected")
	// req, err := http.NewRequest("GET", "/todos", nil)
	// if err != nil {
	//     t.Fatal(err)
	// }
	// rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(TodoShow)
	// handler.ServeHTTP(rr, req)
	// if status := rr.Code; status != http.StatusOK {
	//     t.Errorf("handler returned wrong status code: got %v want %v",
	//         status, http.StatusOK)
	// }
}

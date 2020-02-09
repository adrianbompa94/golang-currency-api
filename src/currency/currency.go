package currency

import (
	"net/http"
	"io/ioutil"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Todo struct {
	Slug  string `json:"slug"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/todo", GetATodo)
	return router
}

func GetATodo(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://api.exchangeratesapi.io/latest")
	if err != nil {}
	data, _ := ioutil.ReadAll(response.Body)
	render.JSON(w, r, string(data))
}
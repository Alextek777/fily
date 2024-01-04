package app

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func main() {

	templates = template.Must(template.ParseGlob("resources/html/*.html"))

	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)

}

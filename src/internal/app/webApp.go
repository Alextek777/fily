package app

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Alextek777/fily/src/internal/config"
	. "github.com/Alextek777/fily/src/internal/storage"
)

type webServer struct {
	cfg   *config.Config
	store Storage
}

func newWebServer(cfg *config.Config, store Storage) *webServer {
	return &webServer{
		cfg:   cfg,
		store: store,
	}
}

var templates *template.Template

func indexHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func (server *webServer) run() {

	templates = template.Must(template.ParseGlob("src/resources/html/*.html"))
	listenAddr := fmt.Sprintf("%s:%s",
		server.cfg.Web.Ip,
		server.cfg.Web.Port,
	)

	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler).Methods("GET")
	http.Handle("/", router)
	http.ListenAndServe(listenAddr, nil)
}

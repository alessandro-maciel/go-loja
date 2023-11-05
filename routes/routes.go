package routes

import (
	"net/http"

	"github.com/alessandro-maciel/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.Create)
	http.HandleFunc("/insert", controllers.Store)
	http.HandleFunc("/delete", controllers.Delete)
}

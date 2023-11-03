package routes

import (
	"net/http"

	"github.com/alessandro-maciel/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}

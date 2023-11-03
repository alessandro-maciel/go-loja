package main

import (
	"net/http"

	"github.com/alessandro-maciel/routes"
	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}

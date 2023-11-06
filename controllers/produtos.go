package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/alessandro-maciel/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosOsProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)
}

func Create(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Store(w http.ResponseWriter, r *http.Request) {
	updateOrCreate(w, r)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	produto := models.GetProduto(id)

	templates.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	updateOrCreate(w, r)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	models.DeleteProduto(id)

	http.Redirect(w, r, "/", 301)
}

func updateOrCreate(w http.ResponseWriter, r *http.Request) {
	produto := models.Produto{}

	var idConv int
	id := r.FormValue("id")
	if id != "" {
		id, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do id", err)
		}
		idConv = id
	}

	produto.Id = idConv
	produto.Nome = r.FormValue("nome")
	produto.Descricao = r.FormValue("descricao")

	preco, err := strconv.ParseFloat(r.FormValue("preco"), 64)
	if err != nil {
		log.Println("Erro na conversão do preço", err)
	}
	produto.Preco = preco

	quantidade, err := strconv.Atoi(r.FormValue("quantidade"))
	if err != nil {
		log.Println("Erro na conversão da quantidade", err)
	}
	produto.Quantidade = quantidade

	if produto.Id == 0 {
		produto.CriarNovoProduto()
	}

	if produto.Id != 0 {
		produto.AtualizarProduto()
	}

	http.Redirect(w, r, "/", 301)
}

package models

import "github.com/alessandro-maciel/db"

type Produto struct {
	Id              int
	Nome, Descricao string
	Preco           float64
	Quantidade      int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	var produtos []Produto

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

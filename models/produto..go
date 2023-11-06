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

	selectProdutos, err := db.Query("select * from produtos order by id desc")
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

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}

	defer db.Close()

	return produtos
}

func (p *Produto) CriarNovoProduto() {
	db := db.ConectaComBancoDeDados()

	sql, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	sql.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade)

	defer db.Close()
}

func GetProduto(id string) Produto {
	db := db.ConectaComBancoDeDados()

	result, err := db.Query("select * from produtos where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	produto := Produto{}

	for result.Next() {
		err = result.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)
		if err != nil {
			panic(err.Error())
		}
	}

	defer db.Close()

	return produto
}

func (p *Produto) AtualizarProduto() {
	db := db.ConectaComBancoDeDados()

	sql, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}

	sql.Exec(p.Nome, p.Descricao, p.Preco, p.Quantidade, p.Id)

	defer db.Close()
}

func DeleteProduto(id string) {
	db := db.ConectaComBancoDeDados()

	sql, err := db.Prepare("delete from produtos where id= $1")
	if err != nil {
		panic(err.Error())
	}

	sql.Exec(id)

	defer db.Close()
}

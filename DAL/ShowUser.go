package DAL

import (
	"TesteGoRicardo/DataBase"
	"TesteGoRicardo/Model"
)

func ShowUser(nId int) Model.Names {
	db := DataBase.DbConn()
	// Usa o ID para fazer a consulta e tratar erros
	selDB, err := db.Query("SELECT * FROM names WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	// Monta a strcut para ser utilizada no template
	n := Model.Names{}
	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variaveis
		var id int
		var name, email string

		// Faz o Scan do SELECT
		err = selDB.Scan(&id, &name, &email)
		if err != nil {
			panic(err.Error())
		}
		// Envia os resultados para a struct
		n.Id = id
		n.Name = name
		n.Email = email
	}
	// Fecha a conexão
	defer db.Close()

	return n
}

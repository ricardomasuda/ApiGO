package DAL

import (
	"TesteGoRicardo/DataBase"
	"TesteGoRicardo/Model"
)

func ListUser() []Model.Names {
	db := DataBase.DbConn()

	selDB, err := db.Query("SELECT * FROM `names` ORDER BY `names`.`name` ASC")
	if err != nil {
		panic(err.Error())
	}

	// Monta a struct para ser utilizada no template
	n := Model.Names{}

	// Monta um array para guardar os valores da struct
	res := []Model.Names{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variáveis
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

		// Junta a Struct com Array de Struct
		res = append(res, n)
	}
	return res
}

package DAL

import (
	"TesteGoRicardo/DataBase"
	"log"
)

func UpdateUser(name string, email string, id int) bool {
	// Abre a conexão com o banco de dados usando a função: dbConn()
	db := DataBase.DbConn()

	// Verifica o METHOD do formulário passado

	// Prepara a SQL e verifica errors
	insForm, err := db.Prepare("UPDATE names SET name=?, email=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	// Insere valores do formulário com a SQL tratada e verifica erros
	_,err=insForm.Exec(name, email, id)
	if err != nil {
		return false
	}

	// Exibe um log com os valores digitados no formulario
	log.Println("UPDATE: Name: " + name + " |E-mail: " + email)

	// Encerra a conexão do dbConn()
	defer db.Close()
	return true
}

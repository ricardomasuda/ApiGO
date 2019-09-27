package DAL

import (
	"TesteGoRicardo/DataBase"
	"log"
)

func DeletarUser(nId  int) bool{
	// Abre conexão com banco de dados usando a função: dbConn()
	db := DataBase.DbConn()

	// Prepara a SQL e verifica errors
	delForm, err := db.Prepare("DELETE FROM names WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	// Insere valores do form com a SQL tratada e verifica errors
	_, err = delForm.Exec(nId)
	if err != nil {
		return false
	}
	// Exibe um log com os valores digitados no form
	log.Println("DELETE")

	// Encerra a conexão do dbConn()
	defer db.Close()
	return true
}

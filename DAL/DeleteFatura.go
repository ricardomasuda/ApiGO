package DAL

import (
	"TesteGoRicardo/DataBase"
	"log"
	"strconv"
)

func DeleteFatura(nId int) (bool){
	// Abre conexão com banco de dados usando a função: dbConn()
	db := DataBase.DbConn()

	// Prepara a SQL e verifica errors
	delForm, err := db.Prepare("DELETE FROM fatura WHERE id_fatura=?")
	if err != nil {
		panic(err.Error())
	}

	// Insere valores do form com a SQL tratada e verifica errors
	_,err= delForm.Exec(nId)
	if err != nil {
		return false
	}

	// Exibe um log com os valores digitados no form
	log.Println("DELETE :"+strconv.Itoa(nId))

	// Encerra a conexão do dbConn()
	defer db.Close()
	return true
}

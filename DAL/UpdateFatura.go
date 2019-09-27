package DAL

import (
	"TesteGoRicardo/DataBase"
	"log"
	"strconv"
)

func UpdateFatura(valor int, categoria string, id int , status string) bool {

	// Abre a conexão com o banco de dados usando a função: dbConn()
	db := DataBase.DbConn()

	// Prepara a SQL e verifica errors
	insForm, err := db.Prepare("UPDATE fatura SET valor=?, categoria=?,status=? WHERE id_fatura=?")
	if err != nil {
		panic(err.Error())
	}

	// Insere valores do formulário com a SQL tratada e verifica erros
	_,err= insForm.Exec(valor, categoria,status ,id)
	if err != nil {
		return false
	}

	// Exibe um log com os valores digitados no formulario
	log.Println("UPDATE: categoria: " + categoria + " |faura: " + strconv.Itoa(valor) +" |status: "+status)

	// Encerra a conexão do dbConn()
	defer db.Close()

	return true
}

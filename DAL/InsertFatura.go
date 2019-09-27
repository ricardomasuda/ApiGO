package DAL

import (
	"TesteGoRicardo/DataBase"
	"TesteGoRicardo/Model"
	"log"
	"strconv"
)

func InsertFatura(fatura *Model.Fatura)  {
	db := DataBase.DbConn()

	// Prepara a SQL e verifica errors
	insForm, err := db.Prepare("INSERT INTO `fatura`( `idusuario`, `nome_empresa`, `valor`, `data_vencimento`, `pagou`) VALUES (?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}

	// Insere valores do formulario com a SQL tratada e verifica errors
	id,err :=insForm.Exec(fatura.Idusuario,fatura.Nome_empresa,fatura.Valor,fatura.Data_vencimento,fatura.Pagou)
	if err != nil {
		return
	}

	LastInsertId, err := id.LastInsertId()
	fatura.IdFatura= int(LastInsertId)
	// Exibe um log com os valores digitados no formulário
	log.Println("INSERT: ID " + strconv.FormatInt(LastInsertId, 10) + " Idusuario: " + strconv.Itoa(fatura.Idusuario) + " | Nome_empresa: " + fatura.Nome_empresa)
	// Encerra a conexão do dbConn()
	defer db.Close()


}
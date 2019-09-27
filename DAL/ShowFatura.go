package DAL

import (
	"TesteGoRicardo/DataBase"
	"TesteGoRicardo/Model"
)

func ShowFatura(nId int) Model.Fatura {

	//abre conexão
	db := DataBase.DbConn()


	// Usa o ID para fazer a consulta e tratar erros
	selDB, err := db.Query("SELECT * FROM `fatura` WHERE `idfatura`=?", nId)
	if err != nil {
		panic(err.Error())
	}

	// Monta a strcut para ser utilizada no template
	n := Model.Fatura{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variaveis
		var idfatura,idusuario,valor,pagou int
		var nome_empresa,data_vencimento string

		// Faz o Scan do SELECT
		err = selDB.Scan(&idfatura, &idusuario, &nome_empresa,&valor,&data_vencimento,&pagou)
		if err != nil {
			panic(err.Error())
		}

		n.Idusuario = idusuario
		n.IdFatura = idfatura
		n.Pagou = pagou
		n.Valor = valor
		n.Nome_empresa = nome_empresa
		n.Data_vencimento = data_vencimento
	}
	// Fecha a conexão
	defer db.Close()

	return n
}

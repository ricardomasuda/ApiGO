package DAL

import (
	"TesteGoRicardo/DataBase"
	"TesteGoRicardo/Model"
)

func ListFatura()  []Model.Fatura {

	db := DataBase.DbConn()

	selDB, err := db.Query("SELECT `idfatura`, `idusuario`, `nome_empresa`, `valor`, `data_vencimento`, `pagou` FROM `fatura` WHERE 1")
	if err != nil {
		panic("[ListFatura]Erro ao buscar informações no banco"+err.Error())
	}

	// Monta a struct para ser utilizada no template
	n := Model.Fatura{}

	// Monta um array para guardar os valores da struct
	res := []Model.Fatura{}

	// Realiza a estrutura de repetição pegando todos os valores do banco
	for selDB.Next() {
		// Armazena os valores em variáveis
		var idfatura,idusuario,valor,pagou int
		var nome_empresa,data_vencimento string


		// Faz o Scan do SELECT
		err = selDB.Scan(&idfatura, &idusuario, &nome_empresa,&valor,&data_vencimento,&pagou)
		if err != nil {
			panic("[ListFatura]Erro ao fazer o Scan do SELECT"+ err.Error())
		}

		// Envia os resultados para a struct
		n.Idusuario = idusuario
		n.IdFatura = idfatura
		n.Pagou = pagou
		n.Valor = valor
		n.Nome_empresa = nome_empresa
		n.Data_vencimento = data_vencimento


		// Junta a Struct com Array de Struct Fatura
		res = append(res, n)
	}
	return res
}

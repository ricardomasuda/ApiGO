package Controllers

import (
	"TesteGoRicardo/DAL"
	"TesteGoRicardo/Model"
)

func Deleteteste() bool{
	var check bool
	//1 setup exemplo : Criar e salvar  usuario fake no banco
	name := "RICARDO"
	email:= "mricardo1611@gmail.com"
	id:=DAL.InsertUser(name , email)

	//2 exercise exemplo : Chamar o controller
	DAL.DeletarUser(id)

	//3 verify : checa se o retorno é o esperado
	user:=Model.Names{}
	user=DAL.ShowUser(id)
	if  user.Id ==id  {
		check = false
	}

	//4 teardown : limpa os dados fakes
	//nao precisa
	return  check
}

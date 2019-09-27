package Controllers

import (
	"TesteGoRicardo/DAL"
	"TesteGoRicardo/Model"
	"testing"

)

func testeUpdate(t *testing.T) bool{

	var check bool
	name := "RICARDO"
	email:= "mricardo1611@gmail.com"
	//1 setup exemplo : Criar e salvar  usuario fake no banco
	id:=DAL.InsertUser(name , email)


	//2 exercise exemplo : Chamar o controller
	name = "RICARDO MASUDA"
	email= "m-ricardo1611@hotmail.com"
	_=DAL.UpdateUser(name,email, id);

	//3 verify : checa se o retorno é o esperado
	user:=Model.Names{}
	user = DAL.ShowUser(id)
	if user.Name == name && user.Email == email {
		check = true
	} else {
		t.Errorf("Não foi possivel inserir")
		check = false
	}
	//4 teardown : limpa os dados fakes
	DAL.DeletarUser(id)

	//retorna o dado do teste
	return  check
}
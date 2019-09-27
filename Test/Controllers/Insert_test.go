package Controllers

import (
	"TesteGoRicardo/DAL"
	"TesteGoRicardo/Model"
	"testing"
)

func testeInsert(t *testing.T) bool{

	var check bool
	name := "RICARDO"
	email:= "mricardo1611@gmail.com"
	n := Model.Names{}
	//1 setup exemplo : Criar e salvar  usuario fake no banco
	//para o insert não é preciso

	//2 exercise exemplo : Chamar o controller
	id:=DAL.InsertUser(name , email)

	//3 verify : checa se o retorno é o esperado
	n = DAL.ShowUser(id)
	if n.Email == email && n.Name == name {
		check = true
	} else {
		t.Errorf("Não foi possivel inserir")
	}

	//4 teardown : limpa os dados fakes


	//retorna o dado do teste
	return  check
}

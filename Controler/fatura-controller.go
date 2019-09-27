package Controler

import (
	"TesteGoRicardo/DAL"
	"TesteGoRicardo/Model"
	"encoding/json"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"strconv"
)



func EditFatura(w http.ResponseWriter, r *http.Request) {
	// Pega o ID do parametro da URL e  converte pra int
	nId,_ := strconv.Atoi(r.URL.Query().Get("id"))
	// Monta a struct para ser utilizada no template
	n := Model.Fatura{}
	//Busca a fatura no banco
	n = DAL.ShowFatura(nId)
	t,_ := template.ParseFiles("tmpl/EditFatura.html","tmpl/Menu.html","tmpl/header.html","tmpl/headlink.html")
	_ = t.Execute(w, n)

}

func ShowFatura(w http.ResponseWriter, r *http.Request) {
	// Pega o ID do parametro da URL
	nId,_ := mux.Vars(r)["id"]
	// Monta a strcut para ser utilizada no template
	fatura := Model.Fatura{}

	//converte em int
	Id,_ := strconv.Atoi(nId)

	fatura=DAL.ShowFatura(Id)
	resfatura, err := json.Marshal(fatura)
	if err != nil{
		panic(err)
	}


	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resfatura)
}

// Função New apenas exibe o formulário para inserir novos dados
func NewFatura(w http.ResponseWriter, r *http.Request) {
	t,_ := template.ParseFiles("tmpl/NewFatura.html","tmpl/Menu.html","tmpl/header.html","tmpl/headlink.html")
	_ = t.Execute(w, nil)
}

func InsertFatura(w http.ResponseWriter, r *http.Request) {
	//struct
	var fatura Model.Fatura

	_ = json.NewDecoder(r.Body).Decode(&fatura)

	log.Println(fatura.Idusuario)


	if r.Method == "POST" {
		//Inseri a Fatura no banco
		DAL.InsertFatura(&fatura)
	}

	resfatura, err := json.Marshal(fatura)
	if err != nil{
		panic(err)
	}
	//Retorna a fatura
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resfatura)
}

func ListFatura(w http.ResponseWriter, r *http.Request) {

	// Monta um array para guardar os valores da struct
	res := []Model.Fatura{}
	//Busca no banco todas as faturas
	res=DAL.ListFatura()
	resfatura, err := json.Marshal(res)
	if err != nil{
		panic(err)
	}


	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(resfatura)


}

func UpdateFatura(w http.ResponseWriter, r *http.Request) {

	var retorno bool
	// Verifica o METHOD do formulário passado
	if r.Method == "POST" {

		// Pega os campos do formulário
		valor , _:= strconv.Atoi(r.FormValue("valor"))
		categoria := r.FormValue("categoria")
		status := r.FormValue("status")
		id,_ := strconv.Atoi(r.FormValue("uid"))

		//insere no banco as informação de fatura
		retorno = DAL.UpdateFatura(valor , categoria , id, status)

		// Exibe um log com os valores digitados no formulario
		log.Println("UPDATE: categoria: " + categoria + " |faura: " + strconv.Itoa(valor))
	}
	// Retorna a HOME
	if retorno == true {
		http.Redirect(w, r, "/listaFatura", 301)
	}
}

func DeleteFatura(w http.ResponseWriter, r *http.Request) {
	//receber valor do id
	nId,_ := strconv.Atoi(r.URL.Query().Get("id"))
	//retorno do DAL de deletar fatura
	retorno :=DAL.DeleteFatura(nId)
	// Exibe um log com os valores digitados no form
	log.Println("DELETE")
	if retorno == true{
		// Retorna a Home Fatura
		http.Redirect(w, r, "/listaFatura", 301)
	}

}

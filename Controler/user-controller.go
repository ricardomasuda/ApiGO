package Controler

import (
	"TesteGoRicardo/DAL"
	"TesteGoRicardo/Model"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

//var tmpl = template.Must(template.ParseGlob("tmpl/*"))

// Função Edit, edita os dados
func Edit(w http.ResponseWriter, r *http.Request) {
	// Pega o ID do parametro da URL e converte pra int
	nId,_ := strconv.Atoi(r.URL.Query().Get("id"))
	// Monta a struct para ser utilizada no template
	n := Model.Names{}
	//Busca o user no banco
	n=DAL.ShowUser(nId)
	// Mostra o template com formulário preenchido para edição
	t,_ := template.ParseFiles("tmpl/EditUser.html","tmpl/Menu.html","tmpl/header.html","tmpl/headlink.html")
	_ = t.Execute(w, n)
}

func Show(w http.ResponseWriter, r *http.Request) {
	// Pega o ID do parametro da URL
	nId, _ := strconv.Atoi(r.URL.Query().Get("id"))
	// Monta a strcut para ser utilizada no template
	n := Model.Names{}
	n = DAL.ShowUser(nId)
	// Mostra o template
	t,_ := template.ParseFiles("tmpl/ShowUser.html","tmpl/Menu.html","tmpl/header.html","tmpl/headlink.html")
	_ = t.Execute(w, n)

}


// Função New apenas exibe o formulário para inserir novos dados
func New(w http.ResponseWriter, r *http.Request) {
	t,_ := template.ParseFiles("tmpl/NewUser.html","tmpl/Menu.html","tmpl/header.html","tmpl/headlink.html")
	_ = t.Execute(w, nil)
}


func Insert(w http.ResponseWriter, r *http.Request) {
	// Verifica o METHOD do fomrulário passado
	if r.Method == "POST" {
		// Pega os campos do formulário
		name := r.FormValue("name")
		email := r.FormValue("email")
		DAL.InsertUser(name,email)
	}
	//Retorna a HOME
	http.Redirect(w, r, "/", 301)
}

func Index(w http.ResponseWriter, r *http.Request) {
	// Monta um array para guardar os valores da struct
	res := []Model.Names{}
	//Busca no banco todos os Users
	res= DAL.ListUser()
	t,_ := template.ParseFiles("tmpl/index.html","tmpl/Menu.html","tmpl/header.html","tmpl/headlink.html")
	_ = t.Execute(w, res)
}


func DeleteUser(w http.ResponseWriter, r *http.Request) {
	//pega o valor via GET
	nId,_ := strconv.Atoi(r.URL.Query().Get("id"))
	//retorno do DAL de delete User
	retorno := DAL.DeletarUser(nId)

	// Exibe um log com os valores digitados no form
	log.Println("DELETE")

	if retorno == true{
		// Retorna a HOME
		http.Redirect(w, r, "/", 301)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Verifica o METHOD do formulário passado
	if r.Method == "POST" {
		// Pega os campos do formulário
		name := r.FormValue("name")
		email := r.FormValue("email")
		id ,_ := strconv.Atoi(r.FormValue("uid"))
		//insert User
		DAL.UpdateUser(name, email , id);
		// Exibe um log com os valores digitados no formulario
		log.Println("UPDATE: Name: " + name + " |E-mail: " + email)
	}
	// Retorna a HOME
	t,_ := template.ParseFiles("tmpl/index.html","tmpl/Menu.html","tmpl/header.html","tmpl/headlink.html")
	_ = t.Execute(w, nil)
}

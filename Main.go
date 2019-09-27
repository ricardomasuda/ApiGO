package main

import (
	"TesteGoRicardo/Controler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	log.Println("Server started on: http://localhost:9001")
	rtr := mux.NewRouter()
	// Gerencia as URLs
	//http.HandleFunc("/", Controler.Index)
	rtr.HandleFunc("/", Controler.Index)
	rtr.HandleFunc("/Fatura/{id}", Controler.ShowFatura).Methods("GET")
	rtr.HandleFunc("/newFatura", Controler.NewFatura).Methods("POST")
	rtr.HandleFunc("/editFatura", Controler.EditFatura)
	rtr.HandleFunc("/Fatura", Controler.ListFatura).Methods("GET")

	//// Ações
	rtr.HandleFunc("/Fatura", Controler.InsertFatura).Methods("POST")
	rtr.HandleFunc("/updateFatura", Controler.UpdateFatura)
	rtr.HandleFunc("/deleteFatura", Controler.DeleteFatura)

	//bootstrap
	rtr.PathPrefix("/layout/").Handler(http.StripPrefix("/layout/", http.FileServer(http.Dir("./layout"))))

	http.Handle("/",rtr)
	// Inicia o servidor na porta 9000/ link localhost:9000
	http.ListenAndServe(":9001", nil)
}

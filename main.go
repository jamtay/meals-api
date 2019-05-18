package main

import (
	"github.com/gorilla/mux"
	"log"
	"meals-api/resources"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	setUpMealApi(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func setUpMealApi(router *mux.Router) {
	router.HandleFunc("/meals", resources.GetMeals).Methods("GET")
	router.HandleFunc("/meals", resources.CreateMeal).Methods("POST")
	//router.HandleFunc("/meals/{id}", blah).Methods("PUT")
	//router.HandleFunc("/meals/{id}", blah).Methods("DELETE")
	//router.HandleFunc("/meals/{id}", blah).Methods("GET")
}

//func setUpItemApi(router *mux.Router) {
// //Set up in different resources folder
//	router.HandleFunc("/meals/{id}/items", resources.GetMealItems).Methods("GET")
//}
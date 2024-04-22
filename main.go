package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Item struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
}

var items []Item

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/item/{id}", getItem).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/item/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/item/{id}", deleteItem).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8080", router))
}

func getItems(w http.ResponseWriter, r *http.Request){
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request){
	path := mux.Vars(r)

	id, err := strconv.Atoi(path["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	for _ , item := range items{
		if item.ID == id{
			json.NewEncoder(w).Encode(item)
		}
	}

	http.NotFound(w, r)
}

func createItem(w http.ResponseWriter, r *http.Request)  {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	item.ID = len(items) + 1
	items = append(items, item)

	json.NewEncoder(w).Encode(item)

}

func updateItem(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)

	id ,err := strconv.Atoi(params["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	var updateItem Item
	err = json.NewDecoder(r.Body).Decode(&updateItem)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

	}

	for i, item := range items{
		if item.ID == id{
			items[i] = updateItem
			json.NewEncoder(w).Encode(updateItem)
			return
		}
	}

	http.NotFound(w,r)
}

func deleteItem(w http.ResponseWriter, r *http.Request)  {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, item := range items{
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	
	http.NotFound(w,r)
}
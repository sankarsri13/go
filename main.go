package main

import (
	"encoding/json"
	"log"
	"net/http"
	
	"github.com/gorilla/mux"
)
type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Name string `json:"name"`
	Author string `json:"author"` 
}
var books []Book
func getAllBooks(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
	
}
func addBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books,book)
	json.NewEncoder(w).Encode(book)

}
func getBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for _,item:=range books{
		if(item.ID==params["id"]){
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
	
}
func deleteBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	params:=mux.Vars(r)
	for index,item:=range books{
		if(item.ID==params["id"]){
			books=append(books[:index],books[index+1:]...)
		}
	}
json.NewEncoder(w).Encode(books)
	
}
func main()  {
	r:= mux.NewRouter()
	books=append(books,Book{ID:"1",Isbn:"1234",Name:"Wings of fire",Author:"Abdul Kalam"})
	r.HandleFunc("/api/books",getAllBooks).Methods("GET")
	r.HandleFunc("/api/book/{id}",getBook).Methods("GET")
	r.HandleFunc("/api/book/{id}",deleteBook).Methods("DELETE")
	r.HandleFunc("/api/books",addBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082",r))
}
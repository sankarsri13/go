package main

import (
	"encoding/json"
	"log"
	"net/http"
	"fmt"
	
	// "github.com/gorilla/mux"
)
type Book struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Name string `json:"name"`
	Author string `json:"author"` 
}
var books []Book

func addBook(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type","application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books,book)
	json.NewEncoder(w).Encode(book)

}


func BookHandler(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "GET":{
		fmt.Print("Get book called")
	w.Header().Set("Content-Type","application/json")
	// params:=mux.Vars(r)
	keys, _ := r.URL.Query()["id"]
	fmt.Println(keys)
	for _,item:=range books{
		if(item.ID==keys[0]){
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
	}
	case "DELETE":{
			w.Header().Set("Content-Type","application/json")
			keys, _ := r.URL.Query()["id"]
	for index,item:=range books{
		if(item.ID==keys[0]){
			books=append(books[:index],books[index+1:]...)
		}
	}
json.NewEncoder(w).Encode(books)
	}
		
	}
}
func allBooksHandler(w http.ResponseWriter, r *http.Request)  {
	switch r.Method {
	case "GET":{
		w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(books)
	}
	case "POST":{
		w.Header().Set("Content-Type","application/json")
		var book Book
		_ = json.NewDecoder(r.Body).Decode(&book)
		books = append(books,book)
		json.NewEncoder(w).Encode(books)
	}
		
	}
}
func main()  {
	
	books=append(books,Book{ID:"1",Isbn:"1234",Name:"Wings of fire",Author:"Abdul Kalam"})
	http.HandleFunc("/api/books",allBooksHandler)
	http.HandleFunc("/api/book/",BookHandler)
	// http.HandleFunc("/api/book/{id}",deleteBook)
	// http.HandleFunc("/api/books",addBook)
	log.Fatal(http.ListenAndServe(":8090",nil))
}
package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
    
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/sankarsri13/go_rest_mongo/helper"
	"github.com/sankarsri13/go_rest_mongo/models"
)
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")


	var books []models.Book

	
	collection := helper.ConnectDB()

	
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		
		var book models.Book
		
		err := cur.Decode(&book) 
		if err != nil {
			log.Fatal(err)
		}

		
		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(books) 
}
func getBook(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

	var book models.Book
	
	var params = mux.Vars(r)

	
	id, _ := primitive.ObjectIDFromHex(params["id"])

	collection := helper.ConnectDB()

	
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(book)
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var book models.Book

	
	_ = json.NewDecoder(r.Body).Decode(&book)

	
	collection := helper.ConnectDB()

	
	result, err := collection.InsertOne(context.TODO(), book)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json")

	
	var params = mux.Vars(r)

	
	id, err := primitive.ObjectIDFromHex(params["id"])

	collection := helper.ConnectDB()

	
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

func main() {
	
	r := mux.NewRouter()

  	
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")


	log.Fatal(http.ListenAndServe(":8008", r))

}


  
package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectDB() *mongo.Collection {

	
	
	ctx, cancel := context.WithTimeout(context.Background(), 10)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
	   "mongodb+srv://vijaysri13:vijaysri13@cluster0-82yqf.mongodb.net/test?retryWrites=true&w=majority",
	))
	if err != nil { log.Fatal(err) }
	
	
	
	

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("go_rest_api").Collection("books")

	return collection
}

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}


func GetError(err error, w http.ResponseWriter) {

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
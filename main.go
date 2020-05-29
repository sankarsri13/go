package main

import (
	
	"log"
	"encoding/json"
	"database/sql"
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	_ "github.com/lib/pq"
)
const (
	host     = "localhost"
	port     = 5435
	user     = "postgres"
	password = "vijaysri13"
	dbname   = "react-go"
  )
type City struct {
	ID int `json:"id"`
	Name string `json:"name"`
	URL string `json:"url"`
}
type Hotel struct {
	ID int 
	Name string 
	City_id int
	Hotel_url string
	Parking bool
	Swimming bool
	Televison bool
	Gym bool
	Price int
	Star int
	Rating int
}
var db *sql.DB
var err error
func connectDatabase() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err = sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  
  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
  fmt.Printf("%T",err)
 return db
}
func getAllCities(w http.ResponseWriter, r *http.Request) {
	

  //getting all cities name and thumbnail
  sqlStatement := `SELECT * FROM city;`
  rows,err2:=db.Query(sqlStatement)
  if err2!=nil{
	  panic(err2)
  }


	var city []City

	defer rows.Close()
  		for rows.Next() {
    

		var temp City
    	err := rows.Scan(&temp.ID, &temp.Name, &temp.URL)
    	if err != nil {
      
      	panic(err)
    	}
    	city=append(city,temp)
  }

  err:= rows.Err()
  if err != nil {
    panic(err)
  }
  w.Header().Set("Content-Type","application/json")
  json.NewEncoder(w).Encode(city)
  
}
func getHotelsInCity(w http.ResponseWriter, r *http.Request)  {
	params:=mux.Vars(r)
	// fmt.Printf("%s",params["id"])
	//getting hotels in city
	sqlStatement := `SELECT * FROM hotel where city_id=$1;`
  rows,err2:=db.Query(sqlStatement,params["id"])
  if err2!=nil{
	  panic(err2)
  }


	var hotel []Hotel

	defer rows.Close()
  		for rows.Next() {
    

		var temp Hotel
    	err := rows.Scan(&temp.ID,&temp.Name,&temp.City_id,&temp.Hotel_url,&temp.Parking,&temp.Swimming,&temp.Televison,&temp.Gym,&temp.Price,&temp.Star,&temp.Rating)
    	if err != nil {
      
      	panic(err)
    	}
    	hotel=append(hotel,temp)
  }

  err:= rows.Err()
  if err != nil {
    panic(err)
  }
  w.Header().Set("Content-Type","application/json")
  json.NewEncoder(w).Encode(hotel)
  
}
func getHotelDetails(w http.ResponseWriter, r *http.Request)  {
	params:=mux.Vars(r)
	//getting single hotel-specific detail
	sqlStatement:=`SELECT * FROM hotel where hotel_id=$1`
	row:=db.QueryRow(sqlStatement,params["id"])
	var hotel Hotel
	err3:=row.Scan(&hotel.ID,&hotel.Name,&hotel.City_id,&hotel.Hotel_url,&hotel.Parking,&hotel.Swimming,&hotel.Televison,&hotel.Gym,&hotel.Price,&hotel.Star,&hotel.Rating)
	if err3!=nil{
		panic(err3)
	}
	w.Header().Set("Content-Type","application/json")
  	json.NewEncoder(w).Encode(hotel)
}
func main()  {
	db=connectDatabase()
	r:= mux.NewRouter()
	r.HandleFunc("/api/cities/",getAllCities).Methods("GET")
	r.HandleFunc("/api/cities/{id}",getHotelsInCity).Methods("GET")
	r.HandleFunc("/api/hotel/{id}",getHotelDetails).Methods("GET")
	log.Fatal(http.ListenAndServe(":8098",r))
}

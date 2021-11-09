package main

//run this go file as this is the final code for rest api with db gorm mux and cors is coming soon
import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Customer struct {
	gorm.Model
	Name    string `json:"name"`
	OrderId int    `json:"order_id"`
}
type Order struct {
	gorm.Model
	Number   uint     `json:"number"`
	Customer Customer `gorm:"ForeignKey:OrderId"`
}

var db *gorm.DB
var e error

func main() {
	db, e = gorm.Open("postgres", "user=postgres password=tatva972000 dbname=postgres sslmode=disable")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Connection Established")
	}
	defer db.Close()
	db.SingularTable(true)
	db.AutoMigrate(&Customer{}, &Order{})
	r := mux.NewRouter()
	r.HandleFunc("/", getnames).Methods("GET")
	r.HandleFunc("/{id}", getbyname).Methods("GET")
	r.HandleFunc("/", addcustomer).Methods("POST")
	r.HandleFunc("/{id}", deletecustomer).Methods("DELETE")
	r.HandleFunc("/{id}", updatecustomer).Methods("PUT")

	http.ListenAndServe(":1991", r)
}
func getnames(w http.ResponseWriter, r *http.Request) {
	// var customers []Customer
	var customers []Customer
	if e := db.Find(&customers).Error; e != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
		json.NewEncoder(w).Encode(&customers)
	}
}
func getbyname(w http.ResponseWriter, r *http.Request) {
	var customers []Customer
	param := mux.Vars(r)
	if e := db.Where("ID = ?", param["id"]).Find(&customers).Error; e != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
		json.NewEncoder(w).Encode(&customers)
	}
}

func addcustomer(w http.ResponseWriter, r *http.Request) {
	var order Order
	var _ = json.NewDecoder(r.Body).Decode(&order)
	db.Create(&order)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(&order)
}
func deletecustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	param := mux.Vars(r)
	if e := db.Where("ID = ?", param["id"]).First(&customer).Error; e != nil {
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
	} else {
		db.Where("ID=?", param["id"]).Delete(&customer)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
	}
}

func updatecustomer(w http.ResponseWriter, r *http.Request) {
	var customer Customer
	param := mux.Vars(r)
	if e := db.Where("ID = ?", param["id"]).First(&customer).Error; e != nil {
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
	} else {
		_ = json.NewDecoder(r.Body).Decode(&customer)
		db.Save(&customer)
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
		json.NewEncoder(w).Encode(&customer)
	}
}

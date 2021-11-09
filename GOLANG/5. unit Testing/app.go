package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	a.Router.HandleFunc("/user", a.createUser).Methods("POST")
	// a.Router.HandleFunc("/user/{id:[0-9]+}", a.getUser).Methods("GET")
	a.Router.HandleFunc("/user/{id:[0-9]+}", a.updateUser).Methods("PUT")
	a.Router.HandleFunc("/user/{id:[0-9]+}", a.deleteUser).Methods("DELETE")
}

type user struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *user) updateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var user user //why i initiated Department variable here?

	param := mux.Vars(r)
	if e := db.Where("id = ?", param["id"]).First(&user).Error; e != nil {
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))

	} else {
		_ = json.NewDecoder(r.Body).Decode(&user)
		db.Save(&user)
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
		json.NewEncoder(w).Encode(&user)

	}
}
func (u *user) deleteUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var users []user
	param := mux.Vars(r)
	if e := db.Where("id = ?", param["id"]).First(&users).Error; e != nil {
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
	} else {
		db.Where("id=?", param["id"]).Delete(&users)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "00")

	}
}
func (u *user) createUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var user user //why i initiated Department variable here?

	var _ = json.NewDecoder(r.Body).Decode(&user)

	db.Create(&user)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(user)
}
func getUsers(db *gorm.DB, start, count int, w http.ResponseWriter, r *http.Request) { //return error in all{[]user
	// var customers []Customer
	type Result struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var result []Result

	if e := db.Find(&result).Error; e != nil { //there has to be some changes regarding the joins bcs only find is not getting all the data only the data that i am asking for which is dep
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
		// log.Error("No Data Found in DB")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
		json.NewEncoder(w).Encode(result)
		// logging()
		// log.Info("Some info. Earth is not flat.")
		// log.Warning("This is a warning")
		// log.Error("Not fatal. An error. Won't stop execution")
		// log.Info("View Function Called")

		// log.Fatal("MAYDAY MAYDAY MAYDAY. Execution will be stopped here")
		// log.Panic("Do not panic")

	}
}

// app.go

package main

import (
	//"database/gorm"
	//"database/gorm"

	"fmt"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	//_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	//_ "github.com/go-gorm-driver/mygorm"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := "host=localhost user=postgres password=tatva972000 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	//connectionString := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai", user, password, dbname)
	// , e = gorm.Open("postgres", "user=postgres password=tatva972000 dbname=postgres sslmode=disable")
	var err error
	a.DB, err = gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Now it should connect")
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/users", a.getUsers).Methods("GET")
	// a.Router.HandleFunc("/user", a.createUser).Methods("POST")
	//a.Router.HandleFunc("/user/{id:[0-9]+}", a.getUser).Methods("GET")
	// a.Router.HandleFunc("/user/{id:[0-9]+}", a.updateUser).Methods("PUT")
	// a.Router.HandleFunc("/user/{id:[0-9]+}", a.deleteUser).Methods("DELETE")
}

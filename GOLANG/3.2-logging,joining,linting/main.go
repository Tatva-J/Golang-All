package main //gvm use go1.17

//run this go file as this is the final code for rest api with db gorm mux and cors is coming soon
import (
	"encoding/json" //for encoding and decoding json files
	"fmt"           //showing ouput to the console
	"net/http"      //for making http server and for making http request and response

	"github.com/gorilla/mux"                     //for Making routes of the server
	"github.com/jinzhu/gorm"                     //for making struct fields table in Underlying Database
	_ "github.com/jinzhu/gorm/dialects/postgres" //the database connection driver
	log "github.com/sirupsen/logrus"             //for doing logs
)

// type Customer struct {
// 	gorm.Model
// 	Name    string `json:"name"`
// 	OrderId int    `json:"order_id"`
// }
// type Order struct {
// 	gorm.Model
// 	Number   uint     `json:"number"`
// 	Customer Customer `gorm:"ForeignKey:OrderId"`
// }
type Department struct {
	gorm.Model            //this will indicate to make this struct db model when called Automigrate in main
	Depno      uint       `json:"dep_no"`
	DepName    string     `json:"dep_name"`
	Employee   []Employee `gorm:"ForeignKey:DepartmentId"` //one to many relationship
	Project    []Project  `gorm:"ForeignKey:DepartmentId"` //one to many relationship
}

type Employee struct {
	gorm.Model
	Name         string `json:"name"`
	Idno         int    `json:"idno"`
	Age          int    `json:"age"`
	DepartmentId int    `json:"dep_id"` //foreign key from table department
}
type Project struct {
	gorm.Model
	Name         string `json:"name"`
	DepartmentId int    `json:"dep_id"` //foreign key from table department
}

var db *gorm.DB //Global db variable for accesing database
var e error     //error variable

func main() { //connectig to the Postgresql DB
	db, e = gorm.Open("postgres", "user=postgres password=tatva972000 dbname=postgres sslmode=disable")
	if e != nil {
		fmt.Println(e) //if any error print error
	} else {
		fmt.Println("Connection Established")
	}
	defer db.Close()                                       //defer will be executed at the end of the main function
	db.SingularTable(true)                                 //Gorm has a in-built method for that that will be set in global level so all tables will be singular.
	db.AutoMigrate(&Employee{}, &Department{}, &Project{}) //create tables
	r := mux.NewRouter()                                   //all routes with associated function
	r.HandleFunc("/", getnames).Methods("GET")
	r.HandleFunc("/", addcustomer).Methods("POST")
	r.HandleFunc("/{id}", deleteCustomer).Methods("DELETE")
	r.HandleFunc("/{id}", updateCustomer).Methods("PUT")
	log.Fatal(http.ListenAndServe(":1991", r))
}

// func Handler() {

// 	//var d Department
// 	//fmt.Println("%T", d)
// 	log.Fatal(http.ListenAndServe(":1991", r))
// 	Handler()
// }

func logging() { //customized log code
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)
}

// func Db_conn() {

// }

//this is join query
//select dep_name,e.name,e.age,p.name from employee e inner join department d on
//e.department_id=d.id inner join project p on p.department_id=d.id;

func getnames(w http.ResponseWriter, r *http.Request) {
	// var customers []Customer
	type Result struct {
		Depno    uint   `json:"dep_no"`
		Dep_Name string `json:"dep_name"`
		// Project_Name    string `json:"project_name"`
		Age  int    `json:"age"`
		Name string `json:"name"`
	}
	var result []Result
	//making join query and will return three fields that we are storing in above created Result struct.
	if e := db.Table("employee").Select("department.depno,department.dep_name,employee.age,project.name").Joins("JOIN department on department.id = employee.department_id").Joins("JOIN project on project.department_id=department.id").Find(&result).Error; e != nil { //there has to be some changes regarding the joins bcs only find is not getting all the data only the data that i am asking for which is dep
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
		log.Error("No Data Found in DB")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
		json.NewEncoder(w).Encode(result)
		logging()
		log.Info("Some info. Earth is not flat.")
		log.Warning("This is a warning")
		log.Error("Not fatal. An error. Won't stop execution")
		log.Info("View Function Called")

		// log.Fatal("MAYDAY MAYDAY MAYDAY. Execution will be stopped here")
		// log.Panic("Do not panic")

	}
	// db.Table("employee").Select("dep_name,employee.name,employee.age").Joins("full join department on employee.department_id = department.id").Scan(&result)

	// db.Table("employee").Select("department.depno,department.dep_name,employee.age,project.name").Joins("JOIN department on department.id = employee.department_id").Joins("JOIN project on project.department_id=department.id").Find(&result)
	// json.NewEncoder(w).Encode(result)
}

// db.Model(&Employee{}).Select("dep_name, employees.name,employees.age,p.name").Joins("inner join project p on p.department_id=d.id").Scan(&result{})
// db.Joins("JOIN employee ON employee.department_id = department.id").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Find(&result)``
func addcustomer(w http.ResponseWriter, r *http.Request) {
	var department Department //by creating this ,it can have another slice inside it of employee and Project struct.

	var _ = json.NewDecoder(r.Body).Decode(&department)

	db.Create(&department)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(department)
	log.Info("EMployee,Project and Department INFORAMTINO ADDED SUCCESFULY!!!")

}

func updateCustomer(w http.ResponseWriter, r *http.Request) {
	var department Department

	param := mux.Vars(r)
	if e := db.Where("id = ?", param["id"]).First(&department).Error; e != nil {
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))

	} else {
		_ = json.NewDecoder(r.Body).Decode(&department)
		db.Save(&department)
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
		json.NewEncoder(w).Encode(&department)
		log.Info("Updated Succesfully!!")

	}
}

// Delete customer
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	var customer []Department
	param := mux.Vars(r)
	if e := db.Where("id = ?", param["id"]).First(&customer).Error; e != nil {
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
	} else {
		db.Where("id=?", param["id"]).Delete(&customer)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "00")
		log.Info("Deleted successfully")

	}
}

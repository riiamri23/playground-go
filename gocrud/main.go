package main

import (
	"database/sql"
	"encoding/json"
	"example/gocrud/utils"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Employee struct {
	ID                           string  `json:"id"`
	Firstname                    *string `json:"firstname"`
	Lastname                     *string `json:"lastname"`
	ContactNo                    *string `json:"contact_no"`
	OfficialEmail                *string `json:"official_email"`
	PersonalEmail                *string `json:"personal_email"`
	IdentityNo                   *string `json:"identity_no"`
	DateOfBirth                  *string `json:"date_of_birth"`
	Gender                       *string `json:"gender"`
	EmergencyContactRelationship *string `json:"emergency_contact_relationship"`
	EmergencyContact             *string `json:"emergency_contact"`
	EmergencyContactAddress      *string `json:"emergency_contact_address"`
	CurrentAddress               *string `json:"current_address"`
	PermanentAddress             *string `json:"permanent_address"`
	City                         *string `json:"city"`
	Designation                  *string `json:"designation"`
	TypeEmp                      *string `json:"type"`
	Status                       *string `json:"status"`
	EmploymentStatus             *string `json:"employment_status"`
	Picture                      *string `json:"picture"`
	JoiningDate                  *string `json:"joining_date"`
	ExitDate                     *string `json:"exit_date"`
	GrossSalary                  *string `json:"gross_salary"`
	Bonus                        *string `json:"Bonus"`
	BranchId                     *string `json:"branch_id"`
	DepartmentId                 *string `json:"department_id"`
}

type Message struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:amri@tcp(127.0.0.1:3306)/db_hris")

	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/employees", getEmployees).Methods("GET")
	router.HandleFunc("/employees", addEmployee).Methods("POST")
	router.HandleFunc("/employees/{id}", getEmployee).Methods("GET")
	router.HandleFunc("/employees/{id}", updEmployee).Methods("PUT")
	router.HandleFunc("/employees/{id}", delEmployee).Methods("DELETE")

	http.ListenAndServe(":2323", router)

}

/******** Example End point
localhost:{port}/employees || Action : GET

localhost:2323/employees */
func getEmployees(w http.ResponseWriter, _ *http.Request) {

	var employees []Employee

	result, err := db.Query("SELECT id, firstname, lastname, contact_no, official_email, personal_email, identity_no, date_of_birth, gender, emergency_contact_relationship, emergency_contact, emergency_contact_address, current_address, permanent_address, city, designation, `type`, status, employment_status, picture, joining_date, exit_date, gross_salary, bonus, branch_id, department_id from employees")
	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	for result.Next() {

		var employee Employee
		err := result.Scan(&employee.ID, &employee.Firstname, &employee.Lastname, &employee.ContactNo, &employee.OfficialEmail, &employee.PersonalEmail, &employee.IdentityNo, &employee.DateOfBirth, &employee.Gender, &employee.EmergencyContactRelationship, &employee.EmergencyContact, &employee.EmergencyContactAddress, &employee.CurrentAddress, &employee.PermanentAddress, &employee.City, &employee.Designation, &employee.TypeEmp, &employee.Status, &employee.EmploymentStatus, &employee.Picture, &employee.JoiningDate, &employee.ExitDate, &employee.GrossSalary, &employee.Bonus, &employee.BranchId, &employee.DepartmentId)

		if err != nil {
			panic(err.Error())
		}

		employees = append(employees, employee)

	}

	spew.Dump(employees)
	res, _ := json.Marshal(employees)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

/******** Example End point
localhost:{port}/employees || Action : POST

localhost:2323/employees

payload :
{
    "firstname": "testing",
    "lastname": "testing",
    "contact_no": "082222222222",
    "official_email": "testing@gmail.com",
    "personal_email": "testing@gmail.com",
    "identity_no": "12333333333",
    "date_of_birth": "1997-12-12",
    "gender": 0,
    "emergency_contact_relationship": "0822222222222",
    "emergency_contact": "082222222222",
    "emergency_contact_address": "Puri permai",
    "password": "string",
    "current_address": "Puri Permai",
    "permanent_address": "Puri Permai",
    "city": "Tangsel",
    "designation": "Heaven",
    "type": "",
    "status": "",
    "employment_status": "permanent",
    "picture": "123.jpg",
    "joining_date": "2023-12-12",
    "exit_date": "",
    "gross_salary": 100000,
    "bonus": 1000,
    "branch_id": 0,
    "department_id": 0,
    "remember_token": 0
} */
func addEmployee(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare("INSERT INTO employees(firstname, lastname, contact_no, official_email, personal_email, identity_no, date_of_birth, gender, emergency_contact_relationship, emergency_contact, emergency_contact_address, password, current_address, permanent_address, city, designation, `type`, status, employment_status, picture, joining_date, exit_date, gross_salary, bonus, branch_id, department_id, deleted_at, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, NULL, NOW(), NOW());")

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	spew.Dump(keyVal)

	firstname := keyVal["firstname"]
	lastname := keyVal["lastname"]
	contact_no := keyVal["contact_no"]
	official_email := keyVal["official_email"]
	personal_email := keyVal["personal_email"]
	identity_no := keyVal["identity_no"]
	date_of_birth := keyVal["date_of_birth"]
	gender, err := strconv.Atoi(keyVal["gender"])
	emergency_contact_relationship := keyVal["emergency_contact_relationship"]
	emergency_contact := keyVal["emergency_contact"]
	emergency_contact_address := keyVal["emergency_contact_address"]
	password := keyVal["password"]
	current_address := keyVal["current_address"]
	permanent_address := keyVal["permanent_address"]
	city := keyVal["city"]
	designation := keyVal["designation"]
	_type := keyVal["type"]
	status, err := strconv.Atoi(keyVal["status"])
	employment_status := keyVal["employment_status"]
	picture := keyVal["picture"]
	joining_date := keyVal["joining_date"]
	exit_date := utils.NewNullString(keyVal["exit_date"])
	gross_salary, err := strconv.Atoi(keyVal["gross_salary"])
	bonus, err := strconv.Atoi(keyVal["bonus"])
	branch_id := keyVal["branch_id"]
	department_id := keyVal["department_id"]

	_, err = stmt.Exec(firstname, lastname, contact_no, official_email, personal_email, identity_no, date_of_birth, gender, emergency_contact_relationship, emergency_contact, emergency_contact_address, password, current_address, permanent_address, city, designation, _type, status, employment_status, picture, joining_date, exit_date, gross_salary, bonus, branch_id, department_id)

	if err != nil {
		panic(err.Error())
	}

	spew.Dump("New Employee was created")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res, _ := json.Marshal(Message{Message: "New Employee was created", Status: http.StatusOK})

	w.Write(res)
	return
}

/******** Example End point
localhost:{port}/employees/{id} || Action : PUT

localhost:2323/employees/13

payload :
{
    "firstname": "testing5",
    "lastname": "testing5",
    "contact_no": "082222222222",
    "official_email": "testing5@gmail.com",
    "personal_email": "testing5@gmail.com",
    "identity_no": "12333333333",
    "date_of_birth": "1997-12-12",
    "gender": "0",
    "emergency_contact_relationship": "0822222222222",
    "emergency_contact": "082222222222",
    "emergency_contact_address": "Puri permai",
    "password": "string",
    "current_address": "Puri Permai",
    "permanent_address": "Puri Permai",
    "city": "Tangsel",
    "designation": "Heaven",
    "type": "",
    "status": "1",
    "employment_status": "permanent",
    "picture": "123.jpg",
    "joining_date": "2023-12-13",
    "exit_date": "",
    "gross_salary": "2000000",
    "bonus": "1000",
    "branch_id": "0",
    "department_id": "0",
    "remember_token": "0"
} */
func updEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE employees SET firstname=?, lastname=?, contact_no=?, official_email=?, personal_email=?, identity_no=?, date_of_birth=?, gender=?,  emergency_contact_relationship=?, emergency_contact=?, emergency_contact_address=?, password=?, current_address=?, permanent_address=?, city=?, designation=?, `type`=?, status=?, employment_status=?, picture=?, joining_date=?, exit_date=?, gross_salary=?, bonus=?, branch_id=?, department_id=?, updated_at=NOW() WHERE id=?;")
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)

	firstname := keyVal["firstname"]
	lastname := keyVal["lastname"]
	contact_no := keyVal["contact_no"]
	official_email := keyVal["official_email"]
	personal_email := keyVal["personal_email"]
	identity_no := keyVal["identity_no"]
	date_of_birth := keyVal["date_of_birth"]
	gender, err := strconv.Atoi(keyVal["gender"])
	emergency_contact_relationship := keyVal["emergency_contact_relationship"]
	emergency_contact := keyVal["emergency_contact"]
	emergency_contact_address := keyVal["emergency_contact_address"]
	password := keyVal["password"]
	current_address := keyVal["current_address"]
	permanent_address := keyVal["permanent_address"]
	city := keyVal["city"]
	designation := keyVal["designation"]
	_type := keyVal["type"]
	status, err := strconv.Atoi(keyVal["status"])
	employment_status := keyVal["employment_status"]
	picture := keyVal["picture"]
	joining_date := keyVal["joining_date"]
	exit_date := utils.NewNullString(keyVal["exit_date"])
	gross_salary, err := strconv.Atoi(keyVal["gross_salary"])
	bonus, err := strconv.Atoi(keyVal["bonus"])
	branch_id := keyVal["branch_id"]
	department_id := keyVal["department_id"]

	_, err = stmt.Exec(firstname, lastname, contact_no, official_email, personal_email, identity_no, date_of_birth, gender, emergency_contact_relationship, emergency_contact, emergency_contact_address, password, current_address, permanent_address, city, designation, _type, status, employment_status, picture, joining_date, exit_date, gross_salary, bonus, branch_id, department_id, params["id"])
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res, _ := json.Marshal(Message{Message: "Succesfully Update Employee", Status: http.StatusOK})

	w.Write(res)
	return
}

/******** Example End point
localhost:{port}/employees/{id} || Action : DELETE

localhost:2323/employees/12 */
func delEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := db.Prepare("DELETE FROM employees WHERE id=?")

	if err != nil {
		panic(err.Error())
	}

	_, err = stmt.Exec(params["id"])

	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res, _ := json.Marshal(Message{Message: "Successfully delete employee", Status: http.StatusOK})

	w.Write(res)
	return
}

/******** Example End point
localhost:{port}/employees/{id} || Action : GET

localhost:2323/employees/13 */
func getEmployee(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Query("SELECT id, firstname, lastname, contact_no, official_email, personal_email, identity_no, date_of_birth, gender, emergency_contact_relationship, emergency_contact, emergency_contact_address, current_address, permanent_address, city, designation, `type`, status, employment_status, picture, joining_date, exit_date, gross_salary, bonus, branch_id, department_id from employees where id=?", params["id"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var employee Employee
	totalRows := 0
	for result.Next() {
		totalRows++
		err := result.Scan(&employee.ID, &employee.Firstname, &employee.Lastname, &employee.ContactNo, &employee.OfficialEmail, &employee.PersonalEmail, &employee.IdentityNo, &employee.DateOfBirth, &employee.Gender, &employee.EmergencyContactRelationship, &employee.EmergencyContact, &employee.EmergencyContactAddress, &employee.CurrentAddress, &employee.PermanentAddress, &employee.City, &employee.Designation, &employee.TypeEmp, &employee.Status, &employee.EmploymentStatus, &employee.Picture, &employee.JoiningDate, &employee.ExitDate, &employee.GrossSalary, &employee.Bonus, &employee.BranchId, &employee.DepartmentId)

		if err != nil {
			panic(err.Error())
		}
	}

	spew.Dump(totalRows)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if totalRows == 0 {
		res, _ := json.Marshal(Message{Message: "Employee is not found", Status: http.StatusNoContent})
		w.Write(res)
		return
	}

	res, _ := json.Marshal(employee)
	w.Write(res)
	return
}

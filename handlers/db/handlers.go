package db

import (
	"encoding/json"
	"github.com/Eydzhpee08/university/handlers/models"
	"github.com/Eydzhpee08/university/utils"
	"github.com/gorilla/mux"
	"net/http"
)

//func CreateEmployeeDocx(w http.ResponseWriter, r *http.Request) {
//	json.NewEncoder(w).Encode(resp)
//}

// CreateEmployee
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee
	//var files Files
	//json.NewDecoder(r.Body).Decode(&emp)
	emp.FullName = r.PostFormValue("fullName")
	emp.Email = r.PostFormValue("email")
	emp.NumberPhone = r.PostFormValue("numberPhone")
	emp.Position = r.PostFormValue("position")
	emp.Job = r.PostFormValue("job")
	emp.Ranks = r.PostFormValue("ranks")
	emp.Report = r.PostFormValue("report")
	//Database.Create(&emp)
	path, types := utils.FileSave(r)
	emp.FileName = path
	emp.TypeFile = types
	Database.Create(&emp)

	json.NewEncoder(w).Encode(path)
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []models.Employee
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee models.Employee
	Database.First(&employee, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode(employee)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee models.Employee
	Database.First(&employee, mux.Vars(r)["eid"])
	json.NewDecoder(r.Body).Decode(&employee)
	Database.Save(&employee)
	json.NewEncoder(w).Encode(employee)

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp models.Employee
	Database.Delete(&emp, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode("Employee is deleted ||")
}

func GetAllDocx(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []models.Docx
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)
}

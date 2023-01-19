package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rakesh1602/GoLang/connection"
	"github.com/rakesh1602/GoLang/model"
	"log"
	"net/http"
)

func AddEmployeeDetails(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	db := connection.Connection()
	var employee model.Employee
	json.NewDecoder(request.Body).Decode(&employee)
	err := db.Create(&employee)
	if err == nil {
		log.Fatalln("Employee details should not be empty.")
	}
	if db != nil {
		log.Println("Employee details saved successfully !!")
	}

	json.NewEncoder(response).Encode(employee.ID)
	log.Println("Your employee id {} is: ", employee.ID)
	return
}

func GetEmployeeDetailsById(response http.ResponseWriter, request *http.Request) {
	log.Println("Retrieving employee details")

	response.Header().Set("Content-Type", "application/json")

	db := connection.Connection()

	var employee model.Employee
	param := mux.Vars(request)

	err := db.Model(model.Employee{}).Preload("Department").First(&employee, param["employeeId"])

	if err == nil {
		log.Fatalln("Employee details with the id {} not found: ", employee.ID)
	} else {
		json.NewEncoder(response).Encode(employee)
		return
	}

}

func GetAllEmployees(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	db := connection.Connection()

	var employees []model.Employee

	log.Println("Retrieving all employees details.")

	err := db.Model(model.Employee{}).Preload("Department").Find(&employees)
	if err == nil {
		log.Fatalln("Employees details not found !!")
	} else {
		json.NewEncoder(response).Encode(&employees)
	}
}

func UpdateEmployee(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	db := connection.Connection()
	var param = mux.Vars(request)
	var employee model.Employee
	db.Model(model.Employee{}).Preload("Department").First(&employee, param["employeeId"])
	json.NewEncoder(response).Encode(employee)
	db.Save(&employee)
	return
}

func DeleteEmployeeById(response http.ResponseWriter, request *http.Request) {
	log.Println("Retrieving employee details")

	response.Header().Set("Content-Type", "application/json")

	db := connection.Connection()

	var employee model.Employee
	param := mux.Vars(request)

	err := db.Model(model.Employee{}).Preload("Department").First(&employee, param["employeeId"])
	db.Delete(&employee)

	if err == nil {
		log.Fatalln("Employee details with the id {} not found: ", employee.ID)
	} else {
		json.NewEncoder(response).Encode("Deleted successfully")
		return
	}
}

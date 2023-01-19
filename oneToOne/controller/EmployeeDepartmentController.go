package controller

import (
	"github.com/gorilla/mux"
	"github.com/rakesh1602/GoLang/service"
	"log"
	"net/http"
)

func Routers() {
	router := mux.NewRouter()
	router.HandleFunc("/addemployees", service.AddEmployeeDetails).Methods("POST")
	router.HandleFunc("/getemployees/{employeeId}", service.GetEmployeeDetailsById).Methods("GET")
	router.HandleFunc("/getemployees", service.GetAllEmployees).Methods("GET")
	router.HandleFunc("/updateemployee/{employeeId}", service.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/deleteemployee/{employeeId}", service.DeleteEmployeeById).Methods("DELETE")

	err := http.ListenAndServe(":4300", router)
	if err != nil {
		log.Fatalln(err)
		return
	}

}

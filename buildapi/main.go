package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Course Model for courses
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice string  `json:"coursePrice"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	LastName string `json:"lastName"`
}

var courses []Course

func main() {
	log.Println("Calling main method")
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: string(299), Author: &Author{FullName: "Rakesh Chavan"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: string(199), Author: &Author{FullName: "Rajesh chavan"}})

	r.HandleFunc("/home", serveHome).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", getAllCourse).Methods("GET")
	r.HandleFunc("/addcourse", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	//listener
	log.Fatal(http.ListenAndServe(":4300", r))

}

//Fake DB

// IsEmpty MiddleWare Helper file
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API using GOLang </h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting all users")
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	log.Print("Getting one user")
	w.Header().Set("content-Type", "application/json")

	//grab id from request
	param := mux.Vars(r)

	//loop through the course
	for _, course := range courses {
		if course.CourseId == param["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating one resource")
	w.Header().Set("content-Type", "application/json")

	//if body is nill

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	log.Println("Updating the course")
	w.Header().Set("Content-Type", "application/json")

	//grab the id
	param := mux.Vars(r)

	//loop id, remove, add with  ID

	for index, course := range courses {
		if course.CourseId == param["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = param["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	log.Println("Delete one course")
	w.Header().Set("Content-Type", "applicatioan/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}

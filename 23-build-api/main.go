package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for a course - file
type Course struct {
	CourseID    string  `json:"course_id"`
	CourseName  string  `json:"course_name"`
	Description string  `json:"description"`
	CoursePrice float32 `json:"course_price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"full_name"`
	Website  string `json:"website"`
}

// fake database

var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty(s string) bool {
	return c.CourseName == ""
}

func main() {

	fmt.Println("Hello, playground")

	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{
		CourseID:    "1",
		CourseName:  "Course 1",
		Description: "Course 1 Description",
		CoursePrice: 100.00,
		Author: &Author{
			FullName: "Author 1",
			Website:  "https://author1.com",
		},
	})

	courses = append(courses, Course{
		CourseID:    "2",
		CourseName:  "Course 2",
		Description: "Course 2 Description",
		CoursePrice: 200.00,
		Author: &Author{
			FullName: "Author 2",
			Website:  "https://author2.com",
		},
	})

	// routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", GetAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", GetCourse).Methods("GET")
	r.HandleFunc("/course", CreateCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", UpdateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", DeleteCourse).Methods("DELETE")

	// listen and serve
	log.Fatal(http.ListenAndServe(":8080", r))

}

// contollers - file

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, playground"))
}

func GetAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetCourses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func GetCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetCourse")
	w.Header().Set("Content-Type", "application/json")

	// get id from req params
	params := mux.Vars(r)

	// loop through courses and find with id
	for _, course := range courses {
		if course.CourseID == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// return empty course
	json.NewEncoder(w).Encode("Course not found")
}

func CreateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateCourse")
	w.Header().Set("Content-Type", "application/json")

	// if course is empty, return
	if r.Body == nil {
		json.NewEncoder(w).Encode("Course is empty")
		return
	}

	// if course is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty(course.CourseName) {
		json.NewEncoder(w).Encode("No Course in json body")
		return
	}

	// check only if course name is duplicate
	for _, c := range courses {
		if c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course name already exists")
			return
		}
	}

	// randomize id and append to courses
	rand.Seed(time.Now().UnixNano())
	course.CourseID = strconv.Itoa(rand.Intn(100000000))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)

}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateCourse")
	w.Header().Set("Content-Type", "application/json")

	// get id from req params
	params := mux.Vars(r)

	// loop through courses and find with id
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseID = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	// return empty course
	json.NewEncoder(w).Encode("Course not found")
}

func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteCourse")
	w.Header().Set("Content-Type", "application/json")

	// get id from req params
	params := mux.Vars(r)

	// loop through courses and find with id
	for index, course := range courses {
		if course.CourseID == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course deleted")
			return
		}
	}

	// return empty course
	json.NewEncoder(w).Encode("Course not found")
}

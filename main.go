package main

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

// the Post is a struct  that  represent  a single  post

type Post struct {
	Title string `json:"title"`
	Body  string `json:"content"`
}

// our fake DB
var data = make(map[string]Post) // our fake db dict[integer] => Object

func gen_uuid() string {
	a := uuid.New()
	return a.String()
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/json")

	var newPost Post = Post{}
	// get data
	json.NewDecoder(r.Body).Decode(&newPost)

	// store data

	data[gen_uuid()] = newPost

	// print data
	json.NewEncoder(w).Encode(newPost)

}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	v, exist := data[id]

	// no such object in map
	if !exist {
		w.WriteHeader(400)
		w.Write([]byte("{\"error\": \"No such id\"} "))
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(v)

}
func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(data)
}
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	_, exist := data[id]

	// no such object in map
	if !exist {
		w.WriteHeader(400)
		w.Write([]byte("No such id"))
		return
	}
	// get the updated object
	var UpdatedPost Post
	json.NewDecoder(r.Body).Decode(&UpdatedPost)

	// store it

	data[id] = UpdatedPost

	// return the new object

	json.NewEncoder(w).Encode(data[id])
}
func DeletePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	id := vars["id"]
	_, exist := data[id]

	// no such object in map
	if !exist {
		w.WriteHeader(400)
		w.Write([]byte("No such id"))
		return
	}

	// delete it
	delete(data, id)
	// return 200 status code
	w.WriteHeader(200)
	w.Write([]byte("Successfully deleted"))
}

func main() {

	router := mux.NewRouter()

	// example
	data["f0668f75-b99e-447f-8495-fbe660bb1892"] = Post{"exampleTitle", "Example Content"}

	router.HandleFunc("/", GetPosts).Methods("GET")
	router.HandleFunc("/post/{id}", GetPost).Methods("GET")
	router.HandleFunc("/post/update/{id}", UpdatePost).Methods("PUT")
	router.HandleFunc("/post/delete/{id}", DeletePost).Methods("DELETE")
	router.HandleFunc("/post/create", CreatePost).Methods("POST")

	http.ListenAndServe(":8080", router)
}

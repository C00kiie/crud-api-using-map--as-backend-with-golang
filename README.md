

# crud api in Golang

this is a simple demonstration of using gorilla/mux router with plain golang
`http` library

# routes
```go
	router.HandleFunc("/", GetPosts).Methods("GET")
	router.HandleFunc("/post/{id}", GetPost).Methods("GET")
	router.HandleFunc("/post/update/{id}", UpdatePost).Methods("PUT")
	router.HandleFunc("/post/delete/{id}", DeletePost).Methods("DELETE")
	router.HandleFunc("/post/create", CreatePost).Methods("POST")
```

# tests
inside `tests/` each of the routes can be tested indivdually 


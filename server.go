package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	ID    int
	Name  string
	Score int
}

var Users = []User{
	User{ID: 0, Name: "Ahmed", Score: 10},
	User{ID: 1, Name: "hady", Score: 9},
	User{ID: 2, Name: "sara", Score: 6},
	User{ID: 3, Name: "Jack", Score: 12},
}

func mainServer() {

	http.HandleFunc("/users", handleUsers)

	http.HandleFunc("/", index)

	fmt.Println("starting server")
	http.ListenAndServe(":8090", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling /req")
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message":"hello World :)"}`)
}

func handleUsers(w http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		get(w, req)
	} else if req.Method == "POST" {
		post(w, req)
	} else {
		fmt.Println("handling invalid/ users req")
		errorHandler(w, req, http.StatusMethodNotAllowed, fmt.Errorf("Invalid Method"))
	}
}

func get(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling get req")
	// http://localhost:8090/users
	// http://localhost:8090/users?id=0
	query := req.URL.Query()
	id := query.Get("id")
	var result []byte
	var err error
	// if no id
	if id == "" {
		result, err = json.Marshal(Users)
	} else {
		// conver it from string to int
		idInt, err := strconv.Atoi(id)
		if err == nil {
			result, err = json.Marshal(Users[idInt])
		}
	}

	// if we had any err return status 500
	if err != nil {
		errorHandler(w, req, http.StatusInternalServerError, err)
		return
	}
	// set header return data
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(result))
}

func post(w http.ResponseWriter, req *http.Request) {
	fmt.Println("handling post req")
	var u User
	defer req.Body.Close()
	// read req body
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		errorHandler(w, req, http.StatusInternalServerError, err)
		return
	}
	// decode the body
	err = json.Unmarshal(b, &u)
	if err != nil {
		errorHandler(w, req, http.StatusInternalServerError, err)
		return
	}

	// db user save
	Users = append(Users, u)
	w.WriteHeader(http.StatusCreated)
}

func errorHandler(w http.ResponseWriter, req *http.Request, status int, err error) {
	w.WriteHeader(status)
	w.Header().Add("Content-type", "application/json")
	fmt.Fprintf(w, `{error:%v}`, err.Error())
}

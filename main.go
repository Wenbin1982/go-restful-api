package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"

)

// message struct
type MESSAGE struct {
	TIMESTAMP string `json:"timestamp"`
	USER string `json:"user"`
	TEXT string `json:"text"`
}

// return value struct
type REVAL struct {
	 OK bool `json:"ok"`
}

// Init message var as a slice message struct
var messages []MESSAGE
var reval []REVAL

// GET /messages
func getMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if len(messages) < 100 {
		json.NewEncoder(w).Encode(messages[:len(messages)])
	} else {
		json.NewEncoder(w).Encode(messages[len(messages)-100:])
	}
}

// POST /messages
func updateMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	var message MESSAGE
	messages = append(messages)
	_ = json.NewDecoder(r.Body).Decode(&message)
	message.TIMESTAMP = fmt.Sprintf("%d", time.Now().UTC().Unix())
	messages = append(messages, message)
	json.NewEncoder(w).Encode(reval[1])
	return
}

//GET /users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := make([]string, 0)
	for _, message := range messages {
		users = append(users, message.USER)
	}
	json.NewEncoder(w).Encode(users)}


func main() {
	// Init router
	r := mux.NewRouter()
	// Hardcoded data
	messages = append(messages, MESSAGE{TIMESTAMP: "1491345710.18",USER:"superman",TEXT: "hello"})
	messages = append(messages, MESSAGE{TIMESTAMP: "1491345713.18",USER:"batman",TEXT: "hello"})
	reval = append(reval,REVAL{OK: false})
	reval = append(reval,REVAL{OK: true})
	http.HandleFunc("/status",
		func(c http.ResponseWriter, req *http.Request) {
			c.Write([]byte("alive"))
		})
	r.HandleFunc("/messages", getMessages).Methods("GET")
	r.HandleFunc("/messages", updateMessages).Methods("POST")
	r.HandleFunc("/users", getUsers).Methods("GET")

	// Start server
	err := http.ListenAndServe(":8081", r)
	panic(err)
}
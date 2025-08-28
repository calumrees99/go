package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type User struct {
	Name string `json:"name"`
}

var userCache = make(map[int]User)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /users", getUsers)
	mux.HandleFunc("GET /users/{id}", getUser)
	mux.HandleFunc("POST /users", createUser)
	mux.HandleFunc("DELETE /users/{id}", deleteUser)
	mux.HandleFunc("PATCH /users/{id}", updateUser)

	// Start server
	fmt.Println("Server is listening on :8080")
	http.ListenAndServe(":8080", mux)
}
func getUsers(
	w http.ResponseWriter,
	r *http.Request) {

	users := userCache

	w.Header().Set("Content-Type", "application.json")
	j, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func getUser(
	w http.ResponseWriter,
	r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, ok := userCache[id]
	if !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application.json")
	j, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func createUser(
	w http.ResponseWriter,
	r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	userCache[len(userCache)+1] = user
	w.WriteHeader(http.StatusNoContent)
}

func deleteUser(
	w http.ResponseWriter,
	r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := userCache[id]; !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	delete(userCache, id)
	w.WriteHeader(http.StatusNoContent)
}

func updateUser(
	w http.ResponseWriter,
	r *http.Request) {

	id, error := strconv.Atoi(r.PathValue("id"))
	if error != nil {
		http.Error(w, error.Error(), http.StatusBadRequest)
		return
	}

	if _, ok := userCache[id]; !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	var user User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if user.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	userCache[id] = user

	w.WriteHeader(http.StatusNoContent)
}

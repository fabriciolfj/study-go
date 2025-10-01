package example_http_person

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

var persons = []Person{}

var auth = map[string]string{"admin": "admin"}

func login(username, password string) bool {
	if value, ok := auth[username]; ok && value == password {
		return true
	}
	return false
}

func salvePerson(w http.ResponseWriter, r *http.Request) {
	username, password, _ := r.BasicAuth()
	login(username, password)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var person Person
	err = json.Unmarshal(body, &person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for i := range persons {
		if persons[i].Id == person.Id {
			http.Error(w, "person already exists", http.StatusBadRequest)
			return
		}
	}

	persons = append(persons, person)
	w.WriteHeader(http.StatusCreated)
}

func getPersons(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}

	num, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range persons {
		if persons[i].Id == int64(num) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(persons[i])
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(
		map[string]string{"error": "person not found"},
	)

}

func Execute() {
	http.HandleFunc("POST /api/v1/persons", salvePerson)
	http.HandleFunc("GET /api/v1/persons", getPersons)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}

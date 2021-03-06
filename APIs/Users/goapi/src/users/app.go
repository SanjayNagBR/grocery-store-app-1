package main

import (
	"encoding/json"
	"fmt"
	"log"
	http "net/http"

	. "users/dao"
	. "users/models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var dao = UsersDAO{}

func PingEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Shayona Grocery Store API is alive!")
}

// POST /users: create a new user
func CreateUserEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	user.ID = bson.NewObjectId()
	if err := dao.Insert(user); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, user)
}

// GET /users - get all user
func GetAllUsersEndPoint(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	users, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, users)
}

// GET /users/{email}
func GetUserEndPoint(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	user, err := dao.FindByEmail(params["email"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid User Email")
		return
	}
	respondWithJson(w, http.StatusOK, user)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	dao.Server = "mongodb://shivam:msmpsm1@ds225294.mlab.com:25294/shayona-store"
	dao.Database = "shayona-store"
	dao.Connect()
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingEndPoint).Methods("GET")
	r.HandleFunc("/users", CreateUserEndPoint).Methods("POST")
	r.HandleFunc("/users", GetAllUsersEndPoint).Methods("GET")
	r.HandleFunc("/users/{email}", GetUserEndPoint).Methods("GET")
	if err := http.ListenAndServe(":3001", r); err != nil {
		log.Fatal(err)
	}
}

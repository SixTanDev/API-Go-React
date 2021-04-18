// Implementation to handle API entry points i.e Get, Post, Put DELETE methods.

package main

import (
	"encoding/json" // To be able to handle the json data
	"fmt"
	"io/ioutil" // We use the ReadALL function to be able to read the data that comes to us from the server request and that data is of type slices bytes.
	"net/http"  // To use some important data types.
	"strings"   // Some important functions are used to be able to handle and manipulate the strings, more specifically in the function that controls the PUT and DELETE methods.

	"github.com/gorilla/mux"
)

// Function that controls the routes or entry points of the required methods: PUT, DELETE, POST and GET.

func setupRouterForEmpleyee(router *mux.Router) {

	router.HandleFunc("/", getRegister).Methods("GET")

	// It's OPTIONS arguments is to be able to handle the cros, get doesn't need it.

	router.HandleFunc("/", createRegister).Methods("POST", "OPTIONS")

	// The string {data.id} is a parameter to be able to pass values to the routes for further processing.

	router.HandleFunc("/{data.id}", updateRegister).Methods("PUT", "OPTIONS")

	router.HandleFunc("/{data.id}", deleteRegister).Methods("DELETE", "OPTIONS")
}

// Entry point handled by the DELETE method, this function deletes the data that is saved from the database.
// w = The variable w, is to be able to write data to the server that sent the request.
// r = The content that the server request arrives with

func deleteRegister(w http.ResponseWriter, r *http.Request) {

	// We handle CORS first.

	enableCORS(&w, r)

	if (*r).Method == "OPTIONS" {
		return
	}
	var idEmployee string
	vars := mux.Vars(r)

	// Processing the parameter that are passed from the path to remove the characters '{' and '}'
	// The data in this hash table structure is stored in this way.
	// map[key1:"{data.id}"], but  we need as data.id where it is a type string

	idEmployee = strings.TrimLeft(vars["data.id"], "{")
	idEmployee = strings.TrimRight(idEmployee, "}")

	// Then we connect to the database and finally we execute the command to be able to delete the record of a certain employee.
	db, err := getDB()
	if err != nil {

		panic(err.Error())
	}

	// Security: question mark in mysql command is to avoid Mysql injection.

	_, err = db.Exec("DELETE FROM Employees WHERE ID = ?", idEmployee)

	if err != nil {

		panic(err.Error())
	}

	defer db.Close()

	fmt.Fprintf(w, "Se elimino un rgistro")
}

// Function to update a record in the database,

func updateRegister(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w, r)

	if (*r).Method == "OPTIONS" {
		return
	}

	// This header gives information to the browser so that it knows what data we are sending and tmaibén helps us to debug the code.

	w.Header().Set("Content-type", "application/json")

	var idEmployee string
	vars := mux.Vars(r)

	var updateEmployee employee

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Fprintf(w, "Inserte a valid Register\n")
	}

	// This function helps us deserialize our JSON array to get useful data.
	json.Unmarshal(reqBody, &updateEmployee)

	idEmployee = strings.TrimLeft(vars["data.id"], "{")
	idEmployee = strings.TrimRight(idEmployee, "}")

	// Function to update the data found in the dataBaseController.go file
	updateData(&updateEmployee, idEmployee)
}

// Function to obtain a certain number of records, in our case 10 per page.

func getRegister(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w, r)

	w.Header().Set("Content-type", "application/json")

	// Function that reloads the data on each call and is found in the dataBaseController.go file.
	reloadData()

	json.NewEncoder(w).Encode(Register)
}

// This function helps us create the records that come to us through HTTP requests.

func createRegister(w http.ResponseWriter, r *http.Request) {

	enableCORS(&w, r)

	if (*r).Method == "OPTIONS" {
		return
	}

	w.Header().Set("Content-type", "application/json")

	var newRegister employee

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {

		fmt.Fprintf(w, "Inserte a valid Register\n")
	}

	json.Unmarshal(reqBody, &newRegister)

	// This function connects to the database and creates a new record in the dataBaseController.go file.
	enterData(&newRegister)

	json.NewEncoder(w).Encode(newRegister)
}

// Here are the functions that perform mysql requests.

package main

import "fmt"

// Function that is executed by the function that handles the GET method found
// in the routes.go file, this function reloads the data to be able to send it to the requesting server.

func getData() (int, error) {

	var i int

	db, err := getDB()

	if err != nil {

		return i, err
	}

	rows, err := db.Query("SELECT * FROM Employees")

	for i = 0; rows.Next(); i++ {
		if err != nil {
			return i, err
		}
	}

	// We close the database to free up the resources we use, since they don't believe in trees :3

	defer db.Close()

	return i, err
}

// This function creates a new record in the Mysql database.

func enterData(newEmployee *employee) {

	db, err := getDB()

	if err != nil {

		fmt.Printf("Eror en la base de datos\n")
	}

	fmt.Println(newEmployee)

	// Modifica el EXEC
	_, err = db.Exec("INSERT INTO Employees "+Insert+"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?) ",
		(*newEmployee).PrimerNombre, (*newEmployee).SegundoNombre,
		(*newEmployee).PrimerApellido, (*newEmployee).SegundoApellido,
		(*newEmployee).PaisDelEmpleo, (*newEmployee).TipoIdentificacion,
		(*newEmployee).NumeroIdentifiacion, (*newEmployee).CorreoElectronico,
		(*newEmployee).Area, (*newEmployee).Estado)

	if err != nil {

		fmt.Printf("Error Ingresar datos: Erro %v\n", err.Error())
	}

	// We close the database to free up the resources we use, since they don't believe in trees :3
	defer db.Close()

}

// This function updates a piece of information in the database by means of its ID that is unique for each user.

func updateData(newEmployee *employee, idEmployee string) {

	db, err := getDB()

	if err != nil {

		fmt.Printf("Eror en la base de datos\n")
	}

	fmt.Println(idEmployee)

	_, err = db.Exec("UPDATE Employees SET "+Update+" WHERE ID = ?",
		(*newEmployee).PrimerNombre, (*newEmployee).SegundoNombre,
		(*newEmployee).PrimerApellido, (*newEmployee).SegundoApellido,
		(*newEmployee).PaisDelEmpleo, (*newEmployee).TipoIdentificacion,
		(*newEmployee).NumeroIdentifiacion, (*newEmployee).Area, idEmployee)

	if err != nil {

		fmt.Printf("Error Ingresar datos: Erro %v\n", err.Error())
	}

	// We close the database to free up the resources we use, since they don't believe in trees :3

	defer db.Close()
}

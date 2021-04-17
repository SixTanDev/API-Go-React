// Implementation for:
// 1. We create the schema to save the requested data in the database.
// 2. We create some variables and constants that will be used for requests to the database.
// 3. We create the Global Registry variable to save the current data from the
//    database. This is important for when we send Get requests

package main

import "time"

type employee struct {
	ID                  string     `json:ID`
	PrimerNombre        string     `json:PrimerNombre`
	SegundoNombre       string     `json:SegundoNombre`
	PrimerApellido      string     `json:PrimerApellido`
	SegundoApellido     string     `json:SegundoApellido`
	PaisDelEmpleo       string     `json:PaisDelEmpleo`
	TipoIdentificacion  string     `json:TipoIdentificacion`
	NumeroIdentifiacion string     `json:NumeroIdentifiacion`
	CorreoElectronico   string     `json:CorreoElectronico`
	Area                string     `json:Area`
	Estado              bool       `json:Estado`
	FechaIngreso        *time.Time `json:FechaIngreso`
	FechaRegistro       *time.Time `json:FechaRegistro`
}

// Creamos una estructua de tipo employee para guardar los datos.

var (
	Register   = []employee{}
	indexFinal int
	indexInit  int
)

const (
	Insert = "(PrimerNombre, SegundoNombre, PrimerApellido, SegundoApellido, PaisDelEmpleo, TipoIdentificacion, NumeroIdentifiacion, CorreoElectronico, Area, Estado)"

	Update = "PrimerNombre=?, SegundoNombre=?, PrimerApellido=?, SegundoApellido=?, PaisDelEmpleo=?, TipoIdentificacion=?, NumeroIdentifiacion=?, Area=?"

	Values = "(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
)

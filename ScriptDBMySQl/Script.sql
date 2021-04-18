CREATE DATABASE IF NOT EXISTS CIDENET;
USE CIDENET;

CREATE TABLE IF NOT EXISTS Employees(
          ID int AUTO_INCREMENT,
   	  PrimerNombre varchar(20) not null,
     	  PrimerApellido varchar(20) not null,
     	  SegundoNombre varchar(50),
   	  SegundoApellido varchar(20),
   	  PaisDelEmpleo varchar (50) not null,
   	  TipoIdentificacion varchar(50) not null,
   	  NumeroIdentifiacion  varchar (20) unique not null,
   	  CorreoElectronico varchar(300) unique not null,
   	  Area varchar(50) not null,
   	  Estado bool not null,
     	  FechaIngreso date not null,
   	  FechaRegistro datetime not null,
   	  PRIMARY KEY (ID)
)

package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func init() {
	database, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
	defer database.Close()

	database.Exec(`BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS eventi (
  idevento	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
  codcliente	INTEGER NOT NULL,
  tipoevento	TEXT,
  timestamp	TEXT,
  tipoprodotto	TEXT,
  nomeprodotto	TEXT,
  pagato	NUMERIC,
  note BLOB

);
CREATE TABLE IF NOT EXISTS clienti (
  idcliente	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE,
  codcliente	INTEGER NOT NULL UNIQUE,
  titolo	TEXT,
  cognome	TEXT,
  nome	TEXT,
  mail	TEXT,
  telefono1	TEXT,
  telefono2	TEXT,
  telefono3	TEXT,
  note BLOB
);
COMMIT;`)
}

const (
	dbfile = "crm.db"
)

type clienti struct {
	codcliente int
	titolo     string
	nome       string
	cognome    string
	mail       string
	telefono1  string
	telefono2  string
	telefono3  string
}

func main() {

	database, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
	defer database.Close()

	cliente := new(clienti)
	codcliente := 300

	nomecliente, _ := database.Prepare("select nome from clienti where codcliente = ?")
	cognomecliente, _ := database.Prepare("select cognome from clienti where codcliente = ?")

	cognomecliente.QueryRow(codcliente).Scan(&cliente.cognome)

	rows, err := nomecliente.Query(codcliente)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}

	// var provincia string
	for rows.Next() {
		// 	rows.Scan(&idcliente, &titolo, &cognome, &nome, &mail, &telefono, &loc, &provincia)
		// 	fmt.Println(strconv.Itoa(idcliente), titolo, cognome, nome, telefono, loc, provincia)

		rows.Scan(&cliente.nome)
		fmt.Println(cliente.nome, cliente.cognome)
	}
}

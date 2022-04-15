package Sql

import "database/sql"

//TODO: find a way to define this in a file like the appsettings in .net or even better, secrets
const (
	username = "gouser"
	password = "gopass"
	hostname = "127.0.0.1:3307"
	dbname   = "testgo"
)

func GetDb() (db *sql.DB) {
	db, err := sql.Open("mysql", username+":"+password+"@("+hostname+")/"+dbname)
	if err != nil{
		panic(err.Error())
	}

	return db
}

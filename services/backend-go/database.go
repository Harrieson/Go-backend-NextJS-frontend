package main

import (
	"database/sql"
	"fmt"
)
func database(){
	db, err:=sql.Open("mysql", "root:@tcp(localhost:3306/users")
	if err != nil{
		panic(err.Error())
}
defer db.Close()
fmt.Println("Successful Connection")
insert, err := db.Query("INSERT INTO users_table VALUES('Harrie')")
if err != nil {
	panic(err.Error())
}
defer insert.Close()
fmt.Println("Successful!, Values Added.")
}
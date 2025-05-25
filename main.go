package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/// okay

var (
	db sql.DB
	Cash int = 1000
	RP int = 50
	availableMachines map[availMachine]any
)


func init() {
	db, err = sql.Open("mysql", "osaka:osaka@tcp(127.0.0.1:3306)/go")
	if err != nil {
		log.Fatal("uhh: ", err)
	}
}

func main() {
	for {

	}
}

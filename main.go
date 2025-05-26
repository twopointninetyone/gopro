package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

/// okay

var (
	db *sql.DB
	Cash int = 1000
	RP int = 50
	availableMachineTiers map[*MachineList]any
)


func init() {
	var err error

	db, err = sql.Open("mysql", "osaka:osaka@tcp(127.0.0.1:3306)/go")
	if err != nil {
		log.Fatal("uhh: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("ping err", err)
	}
}

func main() {
	for {

	}
}

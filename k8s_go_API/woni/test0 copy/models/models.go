package models

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func CheckErr(err error) {
	if err != nil {
		fmt.Println("here error : ", err)
		log.Fatal(err)

		panic(err)
	}
}

func Setup() {
	var err error
	db, err = sql.Open("mysql", "user_gin0:dudaji@/db_gin0")
	CheckErr(err)
	//defer db.Close() : main문에서는 defer로 지연시킬수 있지만 함수로짤때는 X
}

func Close() {
	defer db.Close()
}

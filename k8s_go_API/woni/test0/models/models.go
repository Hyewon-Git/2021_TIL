package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func CheckErr(err error) {
	if err != nil {
		fmt.Println("here error : ", err)
		log.Fatal(err)

		panic(err)
	}
}

func Setup() {
	var err error
	/*
		db, err = sql.Open("mysql", "user_gin0:dudaji@/db_gin0")
		//defer db.Close() : main문에서는 defer로 지연시킬수 있지만 함수로짤때는 X
	*/

	dsn := "user_gin0:dudaji@/db_gin0"
	db, err = gorm.Open(mysql.Open(dsn))
	CheckErr(err)

}

func Close() {
	fmt.Println("close")
	//defer db.Close()
}

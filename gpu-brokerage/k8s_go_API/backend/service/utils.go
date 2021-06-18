package service

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func Checkerror(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
}

func ReturnError(err error) (statusCode int, errorMessage string) {
	// 400 -client요청오류 / 500-server 내부오류
	fmt.Println("return error")
	if strings.Contains(err.Error(), "not found") || strings.Contains(err.Error(), "already exists") {
		statusCode = 404
		log.Print(err.Error())
	} else {
		statusCode = 500
		log.Fatal(err)
	}
	errorMessage = err.Error()
	return
}

func GetJSONStr(v interface{}) string {
	jsonObj, err := json.Marshal(v)
	if err != nil {
		panic(err.Error())
	}
	return string(jsonObj)
}

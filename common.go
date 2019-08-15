package main

import (
	"log"
)

//CheckError function
func CheckError(err error) {
	if err != nil {
		log.Println("Error: ", err)
	}
}

//FailOnError - fail on error
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalln("Fatal Error", msg, err)
	}
}

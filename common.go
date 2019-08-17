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
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalln("Fatal Error", msg, err)
	}
}

func logMessage(msg string) {
	log.Println(msg)
}

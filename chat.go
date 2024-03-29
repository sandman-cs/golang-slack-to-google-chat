package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func sendChatMessage(url string, input chan string) {

	for {
		msg := <-input

		data := []byte("{'text':'" + msg + "'}")

		req, err := http.NewRequest("POST", url+"&threadKey="+threadKey, bytes.NewBuffer(data))
		if err != nil {
			log.Fatal("Error reading request. ", err)
		}

		// Set headers
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		// Set client timeout
		client := &http.Client{Timeout: time.Second * 10}

		// Validate cookie and headers are attached
		fmt.Println(req.Header)

		// Send request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Error reading response. ", err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)

		//Need to remove this for PROD, only log for errors!!!!!!
		fmt.Println("response Headers:", resp.Header)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("Error reading body. ", err)
		}
		fmt.Printf("%s\n", body)
	}
}

func sendChatCard(url string, input chan imagePost) {

	part1 := "{\"cards\": [{\"sections\": [{\"widgets\": [{\"image\": {\"imageUrl\": \""
	part2 := "\",\"onClick\": {\"openLink\": {\"url\": \""
	part3 := "\"}}}}]}]}]}"

	for {
		msg := <-input

		data := []byte(part1 + msg.imageLink + part2 + msg.imageURL + part3)

		//For Debug Only
		log.Println(data)

		req, err := http.NewRequest("POST", url+"&threadKey="+threadKey, bytes.NewBuffer(data))
		if err != nil {
			log.Fatal("Error reading request. ", err)
		}

		// Set headers
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		// Set client timeout
		client := &http.Client{Timeout: time.Second * 10}

		// Validate cookie and headers are attached
		fmt.Println(req.Header)

		// Send request
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("Error reading response. ", err)
		}
		defer resp.Body.Close()

		fmt.Println("response Status:", resp.Status)

	}
}

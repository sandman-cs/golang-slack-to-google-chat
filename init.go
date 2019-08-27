package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	slack "github.com/nlopes/slack"
)

var (
	token         string
	api           *slack.Client
	rtm           *slack.RTM
	conf          configuration
	messages      [255]chan string
	imageMessages [255]chan imagePost
	threadKey     string
)

func init() {

	conf.AppName = "Slack-to-Chat-Bridge"
	conf.AppVer = "1.0.0"
	conf.ServerName, _ = os.Hostname()

	//Load Configuration Data
	dat, _ := ioutil.ReadFile("conf.json")
	err := json.Unmarshal(dat, &conf)
	CheckError(err)

	token = conf.SlackToken

	api = slack.New(token)
	rtm = api.NewRTM()
	go rtm.ManageConnection()

	//New Launch Code Here....

	if len(conf.Channels) > 0 {
		//Keep this part, spawn all the cool new stuff...............................
		for index, element := range conf.Channels {
			// Create Channel and launch publish threads.......
			log.Println("Creating Channel #", index)
			messages[index] = make(chan string, 128)
			imageMessages[index] = make(chan imagePost, 128)
			//Spawn Sending threads
			go sendChatMessage(element.ChatURL, messages[index])
			go sendChatCard(element.ChatURL, imageMessages[index])
		}
	} else {
		log.Fatalln("No Channels Configured...")
	}

}

package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	slack "github.com/nlopes/slack"
)

var (
	token string
	api   *slack.Client
	rtm   *slack.RTM
	url   string
	conf  configuration
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

	//Google Variables
	url = conf.Channels[0].ChatURL

}

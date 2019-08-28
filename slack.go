package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	slack "github.com/nlopes/slack"
)

func respond(rtm *slack.RTM, msg *slack.MessageEvent, prefix string) {
	var response string
	text := msg.Text
	text = strings.TrimPrefix(text, prefix)
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)

	//user := msg.Username
	fmt.Println("msg: ", msg)

	fmt.Println("Message Text: ", text)

	acceptedGreetings := map[string]bool{
		"what's up?": true,
		"hey!":       true,
		"yo":         true,
	}
	acceptedHowAreYou := map[string]bool{
		"how's it going?": true,
		"how are ya?":     true,
		"feeling okay?":   true,
	}

	if acceptedGreetings[text] {
		response = "What's up buddy!?!?!"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	} else if acceptedHowAreYou[text] {
		response = "Good. How are you?"
		rtm.SendMessage(rtm.NewOutgoingMessage(response, msg.Channel))
	}
}

func replaceUserIDWithName(rtm *slack.RTM, msg string) (result string) {

	if strings.Contains(msg, "<@") {
		result = msg
		for {
			if strings.Contains(result, "<@") {

				//log.Println(result)
				szTemp := result[strings.Index(result, "<@")+2 : strings.Index(result, "<@")+11]
				//log.Println("User ID: ", szTemp)
				szTemp1 := getUserNameFromID(rtm, szTemp)
				//log.Println("User Name: ", szTemp1)
				//fmt.Sprintf("'@%s'",
				result = strings.Replace(result, "<@"+szTemp+">", "*@"+szTemp1+"*", -1)
				//log.Println("Current Result: ", result)
			} else {
				return result
			}
		}
	} else {
		return msg
	}
}

func getUserFromMessage(msg string) (result string) {

	//var temp string

	if strings.Contains(msg, "<@") {

		result = msg[strings.Index(msg, "<@")+2 : strings.Index(msg, "<@")+11]

	} else {
		return "No user defined"
	}
	return result

	//return result
}

func getImageFromMessage(rtm *slack.RTM, element slack.File) string {

	var url string

	if len(element.Thumb1024) > 10 {
		url = element.Thumb1024
	} else if len(element.Thumb480) > 10 {
		url = element.Thumb480
	} else if len(element.Thumb360) > 10 {
		url = element.Thumb360
	} else {
		url = element.Thumb160
	}

	szFileName := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	szFile := conf.ImagePath + szFileName + ".png"
	f, err := os.Create(szFile)
	if err != nil {
		return szFileName
	}

	defer f.Close()
	rtm.GetFile(url, f)
	return szFileName + ".png"

}

func getSendingUserFromMessage(msg string) string {

	log.Println("Message to get user from: ", msg)

	temp := msg[10:19]
	log.Println("UserID: ", temp)
	return temp
}

func getUserNameFromID(rtm *slack.RTM, usr string) (result string) {

	user, err := rtm.GetUserInfo(usr)
	if err != nil {
		log.Printf("[WARN]  could not grab user information: %s", usr)
		result = "unknown"
	} else {
		result = user.Profile.RealName
	}
	return
}

func getChannelNameFromID(rtm *slack.RTM, szChannel string) (result string) {

	channel, err := rtm.GetChannelInfo(szChannel)
	if err != nil {
		log.Printf("[WARN]  could not grab user information: %s", szChannel)
		result = "unknown"
	} else {
		result = channel.Name
	}
	return
}

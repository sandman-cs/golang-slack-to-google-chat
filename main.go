package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	slack "github.com/nlopes/slack"
)

func main() {

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				res2B, _ := json.Marshal(ev)
				logMessage(fmt.Sprintf("[DEBUG] containing: %s\n", string(res2B)))
				//fmt.Printf("Message: %v\n", ev)

				user := getUserNameFromID(rtm, ev.User)
				channel := getChannelNameFromID(rtm, ev.Channel)
				posted := false

				//getImageFromMessage(rtm, "https://files.slack.com/files-pri/T034AG6NF-FMKC3QMU7/download/image.png")

				if user != "unknown" {

					for index, element := range conf.Channels {
						if ev.Channel == element.SlackChannelID || channel == element.SlackChannelName {

							buf := bytes.Buffer{}

							buf.WriteString(replaceUserIDWithName(rtm, ev.Msg.Text))
							fileList := ev.Msg.Files
							for _, element := range fileList {
								buf.WriteString("\n" + element.Thumb480)
							}
							messages[index] <- fmt.Sprintf("*%s:*\n%s", user, buf.String())
							posted = true
						}
					}
				}

				if !posted {
					logMessage("Message did not match a config entry:")
					logMessage(fmt.Sprintf(" %s (%s) on %s (%s)\n", user, ev.User, channel, ev.Channel))
					//res2B, _ := json.Marshal(ev)
					//logMessage(fmt.Sprintf("[DEBUG] containing: %s\n", string(res2B)))
					//logMessage(fmt.Sprintf("[DEBUG] containing: (%+v)\n", ev))
				}

			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				break Loop

			default:
				//Take no action
			}
		}
	}
}

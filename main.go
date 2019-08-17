package main

import (
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
				fmt.Printf("Message: %v\n", ev)

				user := getUserNameFromID(rtm, ev.User)
				channel := getChannelNameFromID(rtm, ev.Channel)

				posted := false

				if user != "unknown" {

					for index, element := range conf.Channels {
						if ev.Channel == element.SlackChannelID || channel == element.SlackChannelName {
							messages[index] <- fmt.Sprintf("*%s:*\n%s", user, replaceUserIDWithName(rtm, ev.Msg.Text))
							posted = true
						}
					}
				}

				if !posted {
					logMessage("Message did not match a config entry:")
					logMessage(fmt.Sprintf(" %s (%s) on %s (%s)\n", user, ev.User, channel, ev.Channel))
					logMessage(fmt.Sprintf("[DEBUG] containing: (%s)\n", ev.Msg.Text))
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

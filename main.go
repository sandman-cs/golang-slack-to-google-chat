package main

import (
	"fmt"
	"log"

	slack "github.com/nlopes/slack"
)

func main() {

Loop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			//fmt.Print("\nEvent Received: ")
			switch ev := msg.Data.(type) {
			case *slack.ConnectedEvent:
				fmt.Println("Connection counter:", ev.ConnectionCount)

			case *slack.MessageEvent:
				fmt.Printf("Message: %v\n", ev)
				//info := rtm.GetInfo()
				//prefix := fmt.Sprintf("<@%s> ", info.User.ID)

				user := getUserNameFromID(rtm, ev.User)
				channel := getChannelNameFromID(rtm, ev.Channel)

				log.Printf("[DEBUG] received message from %s (%s) on (%s)\n", user, ev.User, channel)
				log.Printf("[DEBUG] containing: (%s)\n", ev.Msg.Text)
				sendToChat(replaceUserIDWithName(rtm, ev.Msg.Text))

				//if strings.Contains(ev.Msg.Text, prefix) {
				//	log.Println("This message was sent to me :-)")
				//}

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

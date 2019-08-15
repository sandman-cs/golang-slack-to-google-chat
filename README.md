# golang-slack-to-google-chat
Slack bot that listens for messages on Slack and posts them to a Google Chat Room.

Sample Config:
`
{
    "Channels": [{
        "SlackChannelName":"test-channel1",
        "SlackChannelID":"optional",
        "ChatURL":"https://chat.googleapis.com/v1/spaces/aaa1xxxx"
    },
    {
        "SlackChannelName":"test-channel2",
        "SlackChannelID":"optional2",
        "ChatURL":"https://chat.googleapis.com/v1/spaces/????"
    "SlackToken":"Put your slack bot user token here"
}
`


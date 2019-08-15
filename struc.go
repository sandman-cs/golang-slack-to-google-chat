package main

// Configuration File Opjects
type configuration struct {
	Channels   []sourceDest
	AppName    string
	AppVer     string
	ServerName string
	LocalEcho  bool
	SlackToken string
}

type sourceDest struct {
	SlackChannelName string
	SlackChannelID   string
	ChatURL          string
}

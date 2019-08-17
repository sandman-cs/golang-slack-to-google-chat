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

/*

//Example add message, need to create admin workflow to update config.


Message: &{{message CMF6NC8VC UM6GHFTRS <@UM6GHFTRS> has joined the channel 1565830764.000900  false [] [] <nil>  false 0 channel_join false  1565830764.000900   <nil> U0B68H1L3     [] 0 []  [] false <nil>  0 T034AG6NF []  false false {[]}} <nil> <nil>}
*/

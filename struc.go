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

type imagePost struct {
	imageLink string
	imageURL  string
}

//Delete later if not needed.
type chatImage struct {
	Cards []struct {
		Sections []struct {
			Widgets []struct {
				Image struct {
					ImageURL string `json:"imageUrl"`
					OnClick  struct {
						OpenLink struct {
							URL string `json:"url"`
						} `json:"openLink"`
					} `json:"onClick"`
				} `json:"image"`
			} `json:"widgets"`
		} `json:"sections"`
	} `json:"cards"`
}

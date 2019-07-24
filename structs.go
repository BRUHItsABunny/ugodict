package ugodict

import gokhttp "github.com/BRUHItsABunny/gOkHttp"

type UrbanClient struct {
	Client  *gokhttp.HttpClient
	BaseURL string
}

type UrbanResponse struct {
	List []*UrbanDefinition `json:"list"`
}

type UrbanDefinition struct {
	Author      string      `json:"author"`
	CurrentVote string      `json:"current_vote"`
	DefId       int         `json:"defid"`
	Definition  string      `json:"definition"`
	Example     string      `json:"example"`
	PermaLink   string      `json:"permalink"`
	SoundUrls   interface{} `json:"sound_urls"`
	ThumbsDown  int         `json:"thumbs_down"`
	ThumbsUp    int         `json:"thumbs_up"`
	Word        string      `json:"word"`
	WrittenOn   string      `json:"written_on"`
}

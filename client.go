package ugodict

import (
	"encoding/json"
	"log"
	"net/http"
)

type UrbanClient struct {
	Client  http.Client
	BaseURL string
}

type UrbanTermResponse struct {
	List []map[string]interface{} `json:"list"`
}

type UrbanResult struct {
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

func GetClient() UrbanClient {
	return UrbanClient{
		Client:  http.Client{},
		BaseURL: "https://api.urbandictionary.com/v0/",
	}
}

func (client UrbanClient) DefineByTerm(word string) ([]UrbanResult, error, error) {
	req, _ := http.NewRequest("GET", client.BaseURL+"define", nil)
	query := req.URL.Query()
	query.Add("term", word)
	req.URL.RawQuery = query.Encode()
	resp, err := client.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	var result UrbanTermResponse
	var resultUrban []UrbanResult
	err2 := json.NewDecoder(resp.Body).Decode(&result)
	if err2 != nil {
		log.Fatalln(err)
	}
	for _, urbanDef := range result.List {
		test := UrbanResult{
			Author:      urbanDef["author"].(string),
			CurrentVote: urbanDef["current_vote"].(string),
			DefId:       int(urbanDef["defid"].(float64)),
			Definition:  urbanDef["definition"].(string),
			Example:     urbanDef["example"].(string),
			PermaLink:   urbanDef["permalink"].(string),
			SoundUrls:   urbanDef["sound_urls"],
			ThumbsDown:  int(urbanDef["thumbs_down"].(float64)),
			ThumbsUp:    int(urbanDef["thumbs_up"].(float64)),
			Word:        urbanDef["word"].(string),
			WrittenOn:   urbanDef["written_on"].(string),
		}
		resultUrban = append(resultUrban, test)
	}
	return resultUrban, err, err2
}

func (client UrbanClient) DefineById(definitionId string) (UrbanResult, error, error) {
	req, _ := http.NewRequest("GET", client.BaseURL+"define", nil)
	query := req.URL.Query()
	query.Add("defid", definitionId)
	req.URL.RawQuery = query.Encode()
	resp, err := client.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	var result UrbanTermResponse
	err2 := json.NewDecoder(resp.Body).Decode(&result)
	if err2 != nil {
		log.Fatalln(err)
	}
	urbanDef := result.List[0]
	resultUrban := UrbanResult{
		Author:      urbanDef["author"].(string),
		CurrentVote: urbanDef["current_vote"].(string),
		DefId:       int(urbanDef["defid"].(float64)),
		Definition:  urbanDef["definition"].(string),
		Example:     urbanDef["example"].(string),
		PermaLink:   urbanDef["permalink"].(string),
		SoundUrls:   urbanDef["sound_urls"],
		ThumbsDown:  int(urbanDef["thumbs_down"].(float64)),
		ThumbsUp:    int(urbanDef["thumbs_up"].(float64)),
		Word:        urbanDef["word"].(string),
		WrittenOn:   urbanDef["written_on"].(string),
	}
	return resultUrban, err, err2
}

func (client UrbanClient) DefineRandom(word string) (UrbanResult, error, error) {
	req, _ := http.NewRequest("GET", client.BaseURL+"random", nil)
	resp, err := client.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	var result UrbanTermResponse
	err2 := json.NewDecoder(resp.Body).Decode(&result)
	if err2 != nil {
		log.Fatalln(err)
	}
	urbanDef := result.List[0]
	resultUrban := UrbanResult{
		Author:      urbanDef["author"].(string),
		CurrentVote: urbanDef["current_vote"].(string),
		DefId:       int(urbanDef["defid"].(float64)),
		Definition:  urbanDef["definition"].(string),
		Example:     urbanDef["example"].(string),
		PermaLink:   urbanDef["permalink"].(string),
		SoundUrls:   urbanDef["sound_urls"],
		ThumbsDown:  int(urbanDef["thumbs_down"].(float64)),
		ThumbsUp:    int(urbanDef["thumbs_up"].(float64)),
		Word:        urbanDef["word"].(string),
		WrittenOn:   urbanDef["written_on"].(string),
	}
	return resultUrban, err, err2
}

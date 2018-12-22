package ugodict

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type UrbanClient struct {
	Client  *http.Client
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

func GetClient() UrbanClient {
	return UrbanClient{
		Client: &http.Client{
			Timeout: time.Second * 10,
			Transport: &http.Transport{
				TLSHandshakeTimeout: 5 * time.Second,
				DisableCompression:  false,
				DisableKeepAlives:   false,
			}},
		BaseURL: "https://api.urbandictionary.com/v0/",
	}
}

func (client UrbanClient) DefineByTerm(word string) ([]*UrbanDefinition, error, error) {
	req, _ := http.NewRequest("GET", client.BaseURL+"define", nil)
	query := req.URL.Query()
	query.Add("term", word)
	req.URL.RawQuery = query.Encode()
	resp, err := client.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err, nil
	}
	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatalln(err2)
		return nil, err, err2
	}
	defer resp.Body.Close()
	result := new(UrbanResponse)
	_ = json.Unmarshal(bodyBytes, result)
	if len(result.List) > 0 {
		return result.List, err, err2
	} else {
		return nil, err, err2
	}
}

func (client UrbanClient) DefineById(definitionId string) (*UrbanDefinition, error, error) {
	req, _ := http.NewRequest("GET", client.BaseURL+"define", nil)
	query := req.URL.Query()
	query.Add("defid", definitionId)
	req.URL.RawQuery = query.Encode()
	resp, err := client.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err, nil
	}
	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatalln(err2)
		return nil, err, err2
	}
	defer resp.Body.Close()
	result := new(UrbanResponse)
	_ = json.Unmarshal(bodyBytes, result)
	if len(result.List) > 0 {
		return result.List[0], err, err2
	} else {
		return nil, err, err2
	}
}

func (client UrbanClient) DefineRandom(word string) (*UrbanDefinition, error, error) {
	req, _ := http.NewRequest("GET", client.BaseURL+"random", nil)
	resp, err := client.Client.Do(req)
	if err != nil {
		log.Fatalln(err)
		return nil, err, nil
	}
	bodyBytes, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		log.Fatalln(err2)
		return nil, err, err2
	}
	defer resp.Body.Close()
	result := new(UrbanResponse)
	_ = json.Unmarshal(bodyBytes, result)
	if len(result.List) > 0 {
		return result.List[0], err, err2
	} else {
		return nil, err, err2
	}
}

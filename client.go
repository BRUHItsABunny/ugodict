package ugodict

import (
	"errors"
	gokhttp "github.com/BRUHItsABunny/gOkHttp"
	"net/http"
	"net/url"
)

func GetClient() UrbanClient {
	client := gokhttp.GetHTTPClient(nil)
	return UrbanClient{
		Client:  &client,
		BaseURL: "https://api.urbandictionary.com/v0/",
	}
}

func (client UrbanClient) DefineByTerm(word string) ([]*UrbanDefinition, error) {

	var err error
	var req *http.Request
	var resp *gokhttp.HttpResponse
	var result UrbanResponse

	req, err = client.Client.MakeGETRequest(client.BaseURL+"define", url.Values{"term": []string{word}}, map[string]string{})
	if err == nil {
		resp, err = client.Client.Do(req)
		if err == nil {
			err = resp.Object(&result)
			if err == nil {
				if len(result.List) > 0 {
					return result.List, nil
				}
				err = errors.New("no results found")
			}
		}
	}
	return nil, err
}

func (client UrbanClient) DefineById(definitionId string) (*UrbanDefinition, error) {

	var err error
	var req *http.Request
	var resp *gokhttp.HttpResponse
	var result UrbanResponse

	req, err = client.Client.MakeGETRequest(client.BaseURL+"define", url.Values{"defid": []string{definitionId}}, map[string]string{})
	if err == nil {
		resp, err = client.Client.Do(req)
		if err == nil {
			err = resp.Object(&result)
			if err == nil {
				if len(result.List) > 0 {
					return result.List[0], nil
				}
				err = errors.New("no results found")
			}
		}
	}
	return nil, err
}

func (client UrbanClient) DefineRandom() (*UrbanDefinition, error) {

	var err error
	var req *http.Request
	var resp *gokhttp.HttpResponse
	var result UrbanResponse

	req, err = client.Client.MakeGETRequest(client.BaseURL+"random", url.Values{}, map[string]string{})
	if err == nil {
		resp, err = client.Client.Do(req)
		if err == nil {
			err = resp.Object(&result)
			if err == nil {
				if len(result.List) > 0 {
					return result.List[0], nil
				}
				err = errors.New("no results found")
			}
		}
	}
	return nil, err
}

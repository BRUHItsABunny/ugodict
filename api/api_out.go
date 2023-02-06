package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/BRUHItsABunny/gOkHttp/requests"
	"github.com/BRUHItsABunny/ugodict/constants"
	"net/http"
	"net/url"
	"strconv"
)

func DefineByTermRequest(ctx context.Context, term string) (*http.Request, error) {
	parameters := url.Values{
		"term": {term},
	}

	req, err := requests.MakeGETRequest(ctx, constants.EndpointDefine,
		requests.NewHeaderOption(DefaultHeaders()),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func DefineByIdRequest(ctx context.Context, definitionId int) (*http.Request, error) {
	parameters := url.Values{
		"defid": {strconv.Itoa(definitionId)},
	}

	req, err := requests.MakeGETRequest(ctx, constants.EndpointDefine,
		requests.NewHeaderOption(DefaultHeaders()),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func RandomDefinitionRequest(ctx context.Context) (*http.Request, error) {
	req, err := requests.MakeGETRequest(ctx, constants.EndpointRandom,
		requests.NewHeaderOption(DefaultHeaders()),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

func VoteRequest(ctx context.Context, voteBody *VoteRequestBody) (*http.Request, error) {
	bodyBytes, err := json.Marshal(voteBody)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal: %w", err)
	}

	headers := DefaultHeaders()
	delete(headers, "content-type")
	req, err := requests.MakePOSTRequest(ctx, constants.EndpointVote,
		requests.NewHeaderOption(headers),
		requests.NewPOSTJSONOption(bodyBytes, false),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakePOSTRequest: %w", err)
	}
	return req, nil
}

func AutoCompleteRequest(ctx context.Context, term string) (*http.Request, error) {
	parameters := url.Values{
		"term": {term},
	}

	req, err := requests.MakeGETRequest(ctx, constants.EndpointAutoComplete,
		requests.NewHeaderOption(DefaultHeaders()),
		requests.NewURLParamOption(parameters),
	)
	if err != nil {
		return nil, fmt.Errorf("requests.MakeGETRequest: %w", err)
	}
	return req, nil
}

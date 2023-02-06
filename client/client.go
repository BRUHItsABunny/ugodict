package client

import (
	"context"
	"fmt"
	"github.com/BRUHItsABunny/ugodict/api"
	"net/http"
)

type UrbanClient struct {
	Client *http.Client
}

func NewUrbanClient(hClient *http.Client) *UrbanClient {
	if hClient == nil {
		hClient = http.DefaultClient
	}
	return &UrbanClient{Client: hClient}
}

func (c *UrbanClient) DefineByTerm(ctx context.Context, term string) ([]*api.Definition, error) {
	req, err := api.DefineByTermRequest(ctx, term)
	if err != nil {
		return nil, fmt.Errorf("api.DefineByTermRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.DefineResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result.List, nil
}

func (c *UrbanClient) DefineById(ctx context.Context, definitionId int) (*api.Definition, error) {
	req, err := api.DefineByIdRequest(ctx, definitionId)
	if err != nil {
		return nil, fmt.Errorf("api.DefineByIdRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.DefineResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	if len(result.List) > 0 {
		return result.List[0], nil
	}
	return nil, nil
}

func (c *UrbanClient) RandomDefinition(ctx context.Context) (*api.Definition, error) {
	req, err := api.RandomDefinitionRequest(ctx)
	if err != nil {
		return nil, fmt.Errorf("api.DefineByTermRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.DefineResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	if len(result.List) > 0 {
		return result.List[0], nil
	}
	return nil, nil
}

func (c *UrbanClient) Vote(ctx context.Context, body *api.VoteRequestBody) (*api.VoteResponse, error) {
	req, err := api.VoteRequest(ctx, body)
	if err != nil {
		return nil, fmt.Errorf("api.DefineByTermRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.VoteResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result, nil
}

func (c *UrbanClient) AutoComplete(ctx context.Context, term string) ([]*api.Suggestion, error) {
	req, err := api.AutoCompleteRequest(ctx, term)
	if err != nil {
		return nil, fmt.Errorf("api.AutoCompleteRequest: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("c.Client.Do: %w", err)
	}

	result, err := api.ObjectParser(api.AutoCompleteResponse{}, resp)
	if err != nil {
		return nil, fmt.Errorf("api.ObjectParser: %w", err)
	}
	return result.Results, nil
}

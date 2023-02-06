package ugodict

import (
	"github.com/BRUHItsABunny/ugodict/api"
	"github.com/BRUHItsABunny/ugodict/client"
	"net/http"
)

func NewUrbanClient(hClient *http.Client) *client.UrbanClient {
	return client.NewUrbanClient(hClient)
}

func NewVoteRequestBody(definitionId int, up bool) *api.VoteRequestBody {
	return api.NewVoteRequestBody(definitionId, up)
}

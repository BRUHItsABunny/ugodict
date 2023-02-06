package api

type VoteRequestBody struct {
	DefinitionId int    `json:"defid"`
	Direction    string `json:"direction"`
}

func NewVoteRequestBody(definitionId int, up bool) *VoteRequestBody {
	result := &VoteRequestBody{
		DefinitionId: definitionId,
		Direction:    "up",
	}
	if !up {
		result.Direction = "down"
	}
	return result
}

type ErrorResponse struct {
	ErrorMessage string `json:"error"`
}

func (err *ErrorResponse) Error() string {
	return err.ErrorMessage
}

type VoteResponse struct {
	Up     int64  `json:"up"`
	Down   int64  `json:"down"`
	Status string `json:"status"`
}

type AutoCompleteResponse struct {
	Results []*Suggestion `json:"results"`
}

type Suggestion struct {
	Preview string `json:"preview"`
	Term    string `json:"term"`
}

type DefineResponse struct {
	List []*Definition `json:"list"`
}

type Definition struct {
	Definition   string `json:"definition"`
	Permalink    string `json:"permalink"`
	ThumbsUp     int    `json:"thumbs_up"`
	Author       string `json:"author"`
	Word         string `json:"word"`
	DefinitionID int    `json:"defid"`
	CurrentVote  string `json:"current_vote"`
	WrittenOn    string `json:"written_on"`
	Example      string `json:"example"`
	ThumbsDown   int    `json:"thumbs_down"`
}

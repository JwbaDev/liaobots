package liaobots


type ModelReq struct {
	Key string
}


type Model struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MaxLength  int    `json:"maxLength"`
	TokenLimit int    `json:"tokenLimit"`
}


type ModelsResponse struct {
	Data []Model `json:"data"`
}
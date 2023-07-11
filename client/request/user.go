package request



type UserReq struct {
	Authcode string `json:"authcode"`
}

type Referral struct {
	Amount  float64 `json:"amount"`
	Balance float64 `json:"balance"`
}

type UserResponse struct {
	Amount                   float64  `json:"amount"`
	AuthCode                 string   `json:"authCode"`
	Balance                  float64  `json:"balance"`
	Gpt4FreeRemain           float64  `json:"gpt4FreeRemain"`
	Link                     string   `json:"link"`
	Referral                 Referral `json:"referral"`
	UnlimitedAdvancedEndTime int      `json:"unlimitedAdvancedEndTime"`
	UnlimitedEndTime         int      `json:"unlimitedEndTime"`
}

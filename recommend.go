package liaobots

type RecommendReq struct {
	Messages []Message `json:"messages"`
	AuthCode string    `json:"authcode"`
}
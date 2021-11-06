package indicina

type CustomStatement struct {
	Statement []CustomTransaction `json:"statement"`
}

func (ms *CustomStatement) IsStatement() bool {
	return true
}

type CustomTransaction struct {
	Id        string `json:"id"`
	Type      string `json:"type"`
	Amount    int    `json:"amount"`
	Narration string `json:"narration"`
	Date      string `json:"date"`
	Balance   int    `json:"balance"`
}

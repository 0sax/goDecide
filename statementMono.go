package indicina

import "time"

type MonoStatement struct {
	Paging struct {
		Total    int    `json:"total"`
		Page     int    `json:"page"`
		Previous string `json:"previous"`
		Next     string `json:"next"`
	} `json:"paging"`
	Data []MonoTransaction `json:"data"`
}

func (ms *MonoStatement) IsStatement() bool {
	return true
}

type MonoTransaction struct {
	Id        string    `json:"_id"`
	Amount    int       `json:"amount"`
	Date      time.Time `json:"date"`
	Narration string    `json:"narration"`
	Type      string    `json:"type"`
	Category  string    `json:"category"`
}

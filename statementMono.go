package goDecide

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
	Id        string `json:"_id"`
	Amount    int    `json:"amount"`
	Date      string `json:"date"`
	Narration string `json:"narration"`
	Type      string `json:"type"`
	Category  string `json:"category"`
}

func NewMonoStatement() *MonoStatement {
	return &MonoStatement{}
}

func (ms *MonoStatement) AddTransaction(
	amt int,
	id, date, narr, debitOrCredit,
	category string) {
	ms.Data = append(ms.Data, MonoTransaction{
		Id:        id,
		Amount:    amt,
		Date:      date,
		Narration: narr,
		Type:      debitOrCredit,
		Category:  category,
	})
}

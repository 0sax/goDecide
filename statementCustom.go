package goDecide

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

func NewCustomStatement() *CustomStatement {
	return &CustomStatement{}
}

func (ms *CustomStatement) AddTransaction(
	amt, balance int,
	id, date, narr, debitOrCredit string) {
	ms.Statement = append(ms.Statement, CustomTransaction{
		Id:        id,
		Amount:    amt,
		Date:      date,
		Narration: narr,
		Type:      debitOrCredit,
		Balance:   balance,
	})
}

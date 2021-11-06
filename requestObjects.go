package indicina

type Client struct {
	token    string
	baseurl  string
	signedIn bool
}

func (c *Client) signIn(signedIn bool) {
	c.signedIn = signedIn
}
func (c *Client) setToken(token string) {
	c.token = token
}

type loginCredentials struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type AnalysisRequest struct {
	Customer      Customer       `json:"customer,omitempty"`
	BankStatement *BankStatement `json:"bankStatement,omitempty"`
	Pdf           string         `json:"pdf,omitempty"`
	BankCode      string         `json:"bank_code,omitempty"`
	CustomerId    string         `json:"customer_id,omitempty"`
}

type BankStatement struct {
	Type    string    `json:"type"`
	Content Statement `json:"content"` //Statement Type
}

type Statement interface {
	IsStatement() bool
}

type Customer struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	LastName  string `json:"lastName"`
	FirstName string `json:"firstName"`
	Phone     string `json:"phone"`
}

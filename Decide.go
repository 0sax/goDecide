package indicina

import (
	"encoding/base64"
	"github.com/0sax/err2"
)

const (
	// statement types
	Mono   = "mono"
	Custom = "custom"
)

func Login(clientId, clientSecret, baseUrl string) (cl *Client, err error) {

	lc := &loginCredentials{
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}

	c := &Client{
		token:    "",
		baseurl:  baseUrl,
		signedIn: false,
	}

	var lr loginResponse

	err = c.standardRequest(
		"POST", "/api/login",
		lc, &lr)
	if err != nil {
		err2.LogErr1("decide: login", "c.standardRequest", err)
		return
	}

	if lr.isSuccess() {
		c.setToken(lr.Data.Token)
		c.signIn(true)
		cl = c
		return
	}

	err = err2.NewClientErr(nil, lr.Message, lr.Code)
	return

}

//TOMORROW, create functions that parse Mono JSON and Custom JSon to Mono and Cusotm Types Respectively

func (cl *Client) ParseMonoStatement(
	customer *Customer, statement *MonoStatement) (
	sum *StatementSummary, err error) {

	arq := &AnalysisRequest{
		Customer: *customer,
		BankStatement: &BankStatement{
			Mono,
			statement,
		},
	}

	var ar analysisResponse

	err = cl.standardRequest(
		"POST", "/bsp",
		arq, &ar)
	if err != nil {
		err2.LogErr1("decide: parseMonoStatement", "c.standardRequest", err)
		return
	}

	if ar.isSuccess() {
		return &ar.Data, nil
	}

	err = err2.NewClientErr(nil, ar.Message, ar.Code)
	return

}

func (cl *Client) ParseCustomStatement(
	customer *Customer, statement *CustomStatement) (
	sum *StatementSummary, err error) {

	arq := &AnalysisRequest{
		Customer: *customer,
		BankStatement: &BankStatement{
			Custom,
			statement,
		},
	}

	var ar analysisResponse

	err = cl.standardRequest(
		"POST", "/bsp",
		arq, &ar)
	if err != nil {
		err2.LogErr1("decide: parseCustomStatement", "c.standardRequest", err)
		return
	}

	if ar.isSuccess() {
		return &ar.Data, nil
	}

	err = err2.NewClientErr(nil, ar.Message, ar.Code)
	return

}

func (cl *Client) ParsePDFUpload(file []byte, bankCode, customerId string) (
	sum *PDFStatementResponse, err error) {

	pdf := base64.StdEncoding.EncodeToString(file)

	arq := &AnalysisRequest{
		Pdf:        pdf,
		BankCode:   bankCode,
		CustomerId: customerId,
	}

	var ar analysisResponsePDF

	err = cl.standardRequest(
		"POST", "/bsp",
		arq, &ar)
	if err != nil {
		err2.LogErr1("decide: parseCustomStatement", "c.standardRequest", err)
		return
	}

	if ar.isSuccess() {
		return &ar.Data, nil
	}

	err = err2.NewClientErr(nil, ar.Message, ar.Code)
	return

}

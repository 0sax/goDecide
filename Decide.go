package goDecide

import (
	"encoding/base64"
	"fmt"
	"github.com/0sax/err2"
)

const (
	// v2 Base Url
	V2BaseUrl = "https://api.indicina.co/api/v2/client"

	// statement types
	Mono   = "mono"
	Custom = "custom"
)

// Login uses your goDecide ID and Key and returns a client with methods for
// performing Decide API functions
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

func (cl *Client) ParseMonoStatement(
	customer *Customer, statement *MonoStatement) (
	sum *StatementSummary, err error) {

	arq := &AnalysisRequest{
		Customer: customer,
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

	err = err2.NewClientErr(nil, ar.Message, 400)
	return

}

func (cl *Client) ParseCustomStatement(
	customer *Customer, statement *CustomStatement) (
	sum *StatementSummary, err error) {

	arq := &AnalysisRequest{
		Customer: customer,
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

func (cl *Client) ParsePDFUpload(file []byte, fileString *string, bankCode, customerId string) (
	sum *PDFStatementResponse, err error) {

	var pdf []byte

	if file != nil {
		pdf = file
	} else if fileString != nil {
		pdf, err = base64.StdEncoding.DecodeString(*fileString)
		if err != nil {
			err = err2.NewClientErr(err, "unable to decode filestring", 400)
			return
		}
	} else {
		err = err2.NewClientErr(nil, "no file provided", 400)
		return
	}

	mpfs := []multipartField{
		{"pdf", "", pdf, "statement101.pdf"},
		{"bank_code", bankCode, nil, ""},
		{"customer_id", customerId, nil, ""},
	}

	var ar analysisResponsePDF

	//change this to a non standard request

	err = cl.multiPartFormRequest("/pdf/extract",
		mpfs, &ar)
	if err != nil {
		err2.LogErr1("decide: parsePDFStatement", "c.standardRequest", err)
		return
	}

	if ar.isSuccess() {
		return &ar.Data, nil
	}

	err = err2.NewClientErr(nil, ar.Message, 400)
	return

}

func (cl *Client) GetPDFJobStatus(jobId string) (
	sum *PDFStatementResponse, err error) {

	var ar analysisResponsePDF

	err = cl.standardRequest("GET",
		fmt.Sprintf("/pdf/extract/%v/status", jobId),
		nil, &ar)
	if err != nil {
		err2.LogErr1("decide: pollPDFStatus", "c.standardRequest", err)
		return
	}

	if ar.isSuccess() {
		return &ar.Data, nil
	}

	err = err2.NewClientErr(nil, ar.Message, ar.Code)
	return

}

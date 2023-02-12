package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var UrlUsdBrl string = "https://economia.awesomeapi.com.br/json/last/USD-BRL"

type Quotation struct {
	Code       string `json:"code"`
	Codein     string `json:"codein"`
	Name       string `json:"name"`
	High       string `json:"high"`
	Low        string `json:"low"`
	VarBid     string `json:"varBid"`
	PctChange  string `json:"pctChange"`
	Bid        string `json:"bid"`
	Ask        string `json:"ask"`
	Timestamp  string `json:"timestamp"`
	CreateDate string `json:"create_date"`
}

type QuotationWrapper struct {
	USDBRL Quotation `json:"USDBRL"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Bem-vindo à API de cotações"))
	})
	mux.HandleFunc("/cotacao", QuotationHandler)
	http.ListenAndServe(":8080", mux)
}

func QuotationHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cotacao" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ctx := r.Context()

	quotation, err := SearchQuotation(ctx)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(quotation.Bid)
}

func SearchQuotation(ctx context.Context) (*Quotation, error) {
	res, err := http.Get(UrlUsdBrl)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var quotation QuotationWrapper
	err = json.Unmarshal(resBody, &quotation)
	if err != nil {
		return nil, err
	}

	return &quotation.USDBRL, err
}

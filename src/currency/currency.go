package currency

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type exchangeRate struct {
	Rates struct {
		USD float32
		GBP float32
	}
	CurrencyCode  string  `json:"code"`
	CurrencyValue float32 `json:"value"`
	Base          string  `json:"base"`
	Date          string  `json:"date"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/get", getATodo)
	return router
}

func getATodo(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	code := queryParams["currency"][0]

	if code == "" {
		code = "GBP"
	}

	if !isValidCurrencyCode(code) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request! API only accepts USD and GBP as currency"))
		return
	}

	queryRequest := buildAPIRequest(code)
	response, err := http.Get(queryRequest)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Internal Server Error"))
		return
	}

	var er exchangeRate
	data, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(data, &er)

	er.CurrencyCode = code
	if code == "GBP" {
		er.CurrencyValue = er.Rates.GBP
	} else if code == "USD" {
		er.CurrencyValue = er.Rates.USD
	}

	render.JSON(w, r, er)
}

func buildAPIRequest(code string) string {
	req, err := http.NewRequest("GET", "https://api.exchangeratesapi.io/latest", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("base", "EUR")
	q.Add("symbols", code)

	req.URL.RawQuery = q.Encode()
	return req.URL.String()
}

func isValidCurrencyCode(code string) bool {
	switch code {
	case
		"USD",
		"GBP":
		return true
	}
	return false
}

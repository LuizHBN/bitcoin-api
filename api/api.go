package api

import (
	"bitcoin-klever-api/models"
	"encoding/json"
	"errors"
	"io"
	"os"

	"net/http"
)

type Config struct {
	BaseURL     string `json:"baseURL"`
	Credentials struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"credentials"`
}

var config Config

func init() {
	loadConfig()
}

func loadConfig() {
	file, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
}

func FetchBitcoinFromAPI(address string) (models.BitcoinAddress, error) {
	var bitcoin models.BitcoinAddress
	err := fetchData("/address/"+address, &bitcoin)

	return bitcoin, err
}

func FetchTxFromAPI(txID string) (models.Transaction, error) {
	var tx models.Transaction
	err := fetchData("/tx/"+txID, &tx)
	return tx, err
}

func makeRequest(url string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", config.BaseURL+url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(config.Credentials.Username, config.Credentials.Password)

	return client.Do(req)
}

func fetchData(url string, model interface{}) error {
	response, err := makeRequest(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return errors.New("request failed with status: " + response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, model)
}

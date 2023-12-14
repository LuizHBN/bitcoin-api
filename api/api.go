package api

import (
	"bitcoin-klever-api/models"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
)

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
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	basicUrl := os.Getenv("URL")
	client := &http.Client{}
	req, err := http.NewRequest("GET", basicUrl+url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)

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

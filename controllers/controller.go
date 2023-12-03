package controllers

import (
	"bitcoin-klever-api/api"
	"bitcoin-klever-api/models"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/gorilla/mux"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status": "healthy"}`))
}

func GetBitcoinData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	bitcoin, err := api.FetchBitcoinFromAPI(address)

	if err != nil {
		http.Error(w, "Invalid Address", http.StatusBadRequest)
		return
	}

	response := models.ConvertBitcoinAdressToBitcoinResponse(bitcoin)

	json.NewEncoder(w).Encode(response)

}

func GetBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	bitcoin, err := api.FetchBitcoinFromAPI(address)

	if err != nil {
		http.Error(w, "Invalid Address", http.StatusBadRequest)
	}

	response := models.CalculateBalance(bitcoin)

	json.NewEncoder(w).Encode(response)

}

func SendHandler(w http.ResponseWriter, r *http.Request) {
	var body models.SendRequest
	json.NewDecoder(r.Body).Decode(&body)
	address := body.Address
	amount := body.Amount

	if address == "" || amount == 0 {
		http.Error(w, "Address and amount are required", http.StatusBadRequest)
		return
	}
	utxos, err := CalculateUtxo(address, amount)
	if err != nil {
		http.Error(w, "Error getting UTXOs", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(utxos)

}

func CalculateUtxo(address string, amount int64) (models.Utxos, error) {
	bitcoin, err := api.FetchBitcoinFromAPI(address)
	if err != nil {
		return models.Utxos{}, err
	}

	requiredAmount := big.NewInt(amount)
	var currentAmount big.Int
	var usedUTXOs []models.Utxo

	for _, txid := range bitcoin.Txids {
		if currentAmount.Cmp(requiredAmount) >= 0 {
			break
		}

		transaction, err := api.FetchTxFromAPI(txid)
		if err != nil {
			return models.Utxos{}, err
		}

		for _, vin := range transaction.Vin {
			if currentAmount.Cmp(requiredAmount) < 0 {
				amountBigInt, success := new(big.Int).SetString(vin.Value, 10)
				if !success {
					return models.Utxos{}, fmt.Errorf("failed to convert amount to big.Int: %s", vin.Value)
				}

				usedUTXOs = append(usedUTXOs, models.Utxo{
					TxID:   vin.Txid,
					Amount: vin.Value,
				})
				currentAmount.Add(&currentAmount, amountBigInt)
			} else {
				break
			}
		}
	}

	response := models.Utxos{
		Utxo: usedUTXOs,
	}
	return response, nil
}

func GetTransactionData(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	txID := vars["tx"]
	tx, err := api.FetchTxFromAPI(txID)

	if err != nil {
		http.Error(w, "Invalid TxID", http.StatusBadRequest)
	}

	addresses := models.GetAddresses(tx)
	response := models.ConvertTransactionToTransactionResponse(tx, addresses)

	json.NewEncoder(w).Encode(response)
}

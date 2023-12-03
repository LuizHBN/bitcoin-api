package models

import (
	"bitcoin-klever-api/utils"
)

type Vin struct {
	Txid      string   `json:"txid"`
	Sequence  uint32   `json:"sequence"`
	N         uint     `json:"n"`
	Addresses []string `json:"addresses"`
	IsAddress bool     `json:"isAddress"`
	Value     string   `json:"value"`
	Vout      uint     `json:"vout,omitempty"`
}

type Vout struct {
	Value     string   `json:"value"`
	N         uint     `json:"n"`
	Spent     bool     `json:"spent"`
	Hex       string   `json:"hex"`
	Addresses []string `json:"addresses"`
	IsAddress bool     `json:"isAddress"`
}

type Transaction struct {
	Txid          string `json:"txid"`
	Version       int    `json:"version"`
	Vin           []Vin  `json:"vin"`
	Vout          []Vout `json:"vout"`
	BlockHash     string `json:"blockHash"`
	BlockHeight   uint   `json:"blockHeight"`
	Confirmations uint   `json:"confirmations"`
	BlockTime     int64  `json:"blockTime"`
	Size          uint   `json:"size"`
	Vsize         uint   `json:"vsize"`
	Value         string `json:"value"`
	ValueIn       string `json:"valueIn"`
	Fees          string `json:"fees"`
	Hex           string `json:"hex"`
}

type TransactionResponse struct {
	Addresses []Addresses `json:"addresses"`
	Block     uint        `json:"block"`
	TxId      string      `json:"txID"`
}

type Addresses struct {
	Address string `json:"address"`
	Value   int64  `json:"value"`
}

func ConvertTransactionToTransactionResponse(tx Transaction, addresses []Addresses) TransactionResponse {
	response := TransactionResponse{
		Addresses: addresses,
		Block:     tx.BlockHeight,
		TxId:      tx.Txid,
	}
	return response
}

func GetAddresses(tx Transaction) []Addresses {
	var addresses []Addresses
	for _, vout := range tx.Vout {
		for _, address := range vout.Addresses {
			addresses = append(addresses, Addresses{
				Address: address,
				Value:   utils.ParseStringToInt(vout.Value),
			})
		}
	}
	for _, vin := range tx.Vin {
		for _, address := range vin.Addresses {
			addresses = append(addresses, Addresses{
				Address: address,
				Value:   utils.ParseStringToInt(vin.Value),
			})
		}
	}
	return addresses

}

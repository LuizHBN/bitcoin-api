package models

type Utxos struct {
	Utxo []Utxo `json:"utxos"`
}
type Utxo struct {
	TxID   string `json:"txid"`
	Amount string `json:"amount"`
}

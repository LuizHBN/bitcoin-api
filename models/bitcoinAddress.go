package models

import "bitcoin-klever-api/utils"

type BalanceDetails struct {
	Confirmed   string `json:"confirmed"`
	Unconfirmed string `json:"unconfirmed"`
}

type BitcoinAddressResponse struct {
	Address        string         `json:"address"`
	Balance        string         `json:"balance"`
	TotalTx        int            `json:"totalTx"`
	BalanceDetails BalanceDetails `json:"Balance"`
	Total          TxTotal        `json:"total"`
}
type BitcoinAddress struct {
	Page               int      `json:"page"`
	TotalPages         int      `json:"totalPages"`
	ItemsOnPage        int      `json:"itemsOnPage"`
	Address            string   `json:"address"`
	Balance            string   `json:"balance"`
	TotalReceived      string   `json:"totalReceived"`
	TotalSent          string   `json:"totalSent"`
	UnconfirmedBalance string   `json:"unconfirmedBalance"`
	UnconfirmedTxs     int      `json:"unconfirmedTxs"`
	Txs                int      `json:"txs"`
	Txids              []string `json:"txids"`
	Total              TxTotal  `json:"total"`
}

type TxTotal struct {
	Sent     string `json:"sent"`
	Received string `json:"received"`
}

type SendRequest struct {
	Address string `json:"address"`
	Amount  int64  `json:"amount"`
}

func ConvertBitcoinAdressToBitcoinResponse(bitcoin BitcoinAddress) BitcoinAddressResponse {
	response := BitcoinAddressResponse{
		Address:        bitcoin.Address,
		Balance:        bitcoin.Balance,
		TotalTx:        bitcoin.Txs,
		BalanceDetails: CalculateBalance(bitcoin),
		Total: TxTotal{
			Sent:     bitcoin.TotalSent,
			Received: bitcoin.TotalReceived,
		},
	}
	return response
}

func CalculateBalance(bitcoin BitcoinAddress) BalanceDetails {
	confirmed := bitcoin.Txs - bitcoin.UnconfirmedTxs
	unconfirmed := bitcoin.UnconfirmedTxs

	response := BalanceDetails{
		Confirmed:   utils.IntToString(confirmed),
		Unconfirmed: utils.IntToString(unconfirmed),
	}

	return response
}

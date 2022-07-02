package datastructs

import "time"

type PortfolioResults struct {
	Isin                   string      `json:"isin"`
	IsinTitle              string      `json:"isin_title"`
	Quantity               int         `json:"quantity"`
	BuyQuantity            int         `json:"buy_quantity"`
	SellQuantity           int         `json:"sell_quantity"`
	BuyPriceAvg            int         `json:"buy_price_avg"`
	BuyPriceMin            int         `json:"buy_price_min"`
	BuyPriceMax            int         `json:"buy_price_max"`
	BuyPriceAvgHistorical  int         `json:"buy_price_avg_historical"`
	SellPriceMin           interface{} `json:"sell_price_min"`
	SellPriceMax           interface{} `json:"sell_price_max"`
	SellPriceAvgHistorical interface{} `json:"sell_price_avg_historical"`
	OrdersTotal            int         `json:"orders_total"`
	SellOrdersTotal        int         `json:"sell_orders_total"`
	BuyOrdersTotal         int         `json:"buy_orders_total"`
}

type GetAccountResult struct {
	CreatedAt             time.Time   `json:"created_at"`
	AccountID             string      `json:"account_id"`
	Firstname             string      `json:"firstname"`
	Lastname              string      `json:"lastname"`
	Email                 string      `json:"email"`
	Phone                 string      `json:"phone"`
	Address               string      `json:"address"`
	BillingAddress        string      `json:"billing_address"`
	BillingEmail          interface{} `json:"billing_email"`
	BillingName           interface{} `json:"billing_name"`
	BillingVat            interface{} `json:"billing_vat"`
	Mode                  string      `json:"mode"`
	DepositID             string      `json:"deposit_id"`
	ClientID              string      `json:"client_id"`
	AccountNumber         string      `json:"account_number"`
	IbanBrokerage         string      `json:"iban_brokerage"`
	IbanOrigin            string      `json:"iban_origin"`
	BankNameOrigin        string      `json:"bank_name_origin"`
	Balance               int         `json:"balance"`
	CashToInvest          int         `json:"cash_to_invest"`
	CashToWithdraw        int         `json:"cash_to_withdraw"`
	AmountBoughtIntraday  int         `json:"amount_bought_intraday"`
	AmountSoldIntraday    int         `json:"amount_sold_intraday"`
	AmountOpenOrders      int         `json:"amount_open_orders"`
	AmountOpenWithdrawals int         `json:"amount_open_withdrawals"`
	AmountEstimateTaxes   int         `json:"amount_estimate_taxes"`
	ApprovedAt            time.Time   `json:"approved_at"`
	TradingPlan           string      `json:"trading_plan"`
	DataPlan              string      `json:"data_plan"`
	TaxAllowance          int         `json:"tax_allowance"`
	TaxAllowanceStart     interface{} `json:"tax_allowance_start"`
	TaxAllowanceEnd       interface{} `json:"tax_allowance_end"`
}

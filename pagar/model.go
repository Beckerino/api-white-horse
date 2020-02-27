package pagar

//Pagar ---
type Pagarout struct {
	ID            string  `json:"id" db:"uuid"`
	Nome          string  `json:"nome" db:"nome"`
	Tipoconta     string  `json:"tipo_conta" db:"tipo"`
	Valor         float32 `json:"valor" db:"valor"`
	Valorpago     float32 `json:"valor_pago" db:"valorpago"`
	Datavenc      string  `json:"data_venc" db:"datavenc"`
	Datapagamento string  `json:"data_pagamento" db:"datapagamento"`
	Situacao      string  `json:"situacao" db:"situacao"`
}
type Pagarin struct {
	ID            string  `json:"id" db:"uuid"`
	Nome          string  `json:"nome" db:"nome"`
	Tipoconta     int     `json:"tipo_conta" db:"tipoconta_id"`
	Valor         float32 `json:"valor" db:"valor"`
	Valorpago     float32 `json:"valor_pago" db:"valorpago"`
	Datavenc      string  `json:"data_venc" db:"datavenc"`
	Datapagamento string  `json:"data_pagamento" db:"datapagamento"`
	Situacao      int     `json:"situacao" db:"situacao"`
}

//Log ---
type Log struct {
	ID   string
	Tela string
	Desc string
}

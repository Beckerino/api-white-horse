package receber

type Receberout struct {
	ID           string  `json:"id" db:"uuid"`
	Nome         string  `json:"nome" db:"nome"`
	Tipoconta    string  `json:"tipo_conta" db:"tipo"`
	Valor        float32 `json:"valor" db:"valor"`
	CpfCnpj      string  `json:"cpfcnpj" db:"cpfcnpj"`
	DataReceber  string  `json:"data_receber" db:"datareceber"`
	DataRecebido string  `json:"data_recebido" db:"datarecebido"`
	Situacao     string  `json:"situacao" db:"situacao"`
}
type Receberin struct {
	ID           string  `json:"id" db:"uuid"`
	Nome         string  `json:"nome" db:"nome"`
	Tipoconta    string  `json:"tipo_conta" db:"tipo"`
	Valor        float32 `json:"valor" db:"valor"`
	CpfCnpj      string  `json:"cpfcnpj" db:"cpfcnpj"`
	DataReceber  string  `json:"data_receber" db:"datareceber"`
	DataRecebido string  `json:"data_recebido" db:"datarecebido"`
	Situacao     string  `json:"situacao" db:"situacao"`
}

//Log ---
type Log struct {
	ID   string
	Tela string
	Desc string
}

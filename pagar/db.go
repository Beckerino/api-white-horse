package pagar

import (
	"encoding/json"
	"fmt"
)

func pagarRead() (result []Pagarout, err error) {
	query := `select uuid, 
		p.nome as tipo,
		c.nome,
		valor::money::numeric::float8,
		valorpago::money::numeric::float8,
		datavenc,
		datapagamento,
		s.nome as situacao
		from public.contapagar as c
		left join tipoconta as p on p.id = tipoconta_id
		left join situacao as s on s.id = c.situacao;`
	db, err := createDB()
	if err != nil {
		println(err)
	}
	err = db.Select(&result, query)
	if err != nil {
		println(err)
	}

	auditor := new(Log)
	auditor.Tela = "Pagar/Read"

	err = auditoria(auditor, result)
	if err != nil {
		fmt.Println("error:", err)
	}

	return result, err
}

func pagarCreate(data *Pagarin) (result string, err error) {
	query := `INSERT INTO public.contapagar
		(nome, tipoconta_id, valor, valorpago, datavenc, datapagamento, situacao)
		VALUES($1, $2, $3, $4, $5, $6, $7);
			`
	db, err := createDB()
	if err != nil {
		println(err)
	}
	_, err = db.Exec(query, data.Nome, data.Tipoconta, data.Valor, data.Valorpago, data.Datavenc, data.Datapagamento, data.Situacao)
	if err != nil {
		println(err)
		return
	}

	auditor := new(Log)
	auditor.Tela = "Pagar/Create"

	err = auditoria(auditor, data)
	if err != nil {
		fmt.Println("error:", err)
	}
	result = "Item Criado : " + data.ID
	return result, err
}

func pagarUpdate(data *Pagarin) (result string, err error) {

	query := `UPDATE public.contapagar SET nome=$1, tipoconta_id=$2, valor=$3, valorpago=$4, datavenc=$5, datapagamento=$6, situacao=$7, uuid=$8 
	where uuid = $8
;
			`
	db, err := createDB()
	if err != nil {
		println(err)
	}
	_, err = db.Exec(query, data.Nome, data.Tipoconta, data.Valor, data.Valorpago, data.Datavenc, data.Datapagamento, data.Situacao, data.ID)
	if err != nil {
		println(err)
		return
	}

	auditor := new(Log)
	auditor.Tela = "Pagar/Update"

	err = auditoria(auditor, data)
	if err != nil {
		fmt.Println("error:", err)
	}
	result = "Item Atualizado : " + data.ID
	return result, err
}

func pagarRemove(data *Pagarin) (result string, err error) {
	query := `DELETE FROM public.contapagar WHERE uuid = $1;`
	db, err := createDB()
	if err != nil {
		println(err)
	}
	_, err = db.Exec(query, data.ID)
	if err != nil {
		println(err)
		return
	}
	auditor := new(Log)
	auditor.Tela = "Pagar/Remove"

	err = auditoria(auditor, data)
	if err != nil {
		fmt.Println("error:", err)
	}
	result = "Item deletado : " + data.ID
	return result, err
}

func auditoria(auditoria *Log, data interface{}) error {
	query := `INSERT INTO public.logauditoria
		(tela, descricao)
		VALUES($1, $2);
		`
	db, err := createDB()
	if err != nil {
		println(err)
	}
	Desc, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error:", err)
	}

	_, err = db.Exec(query, auditoria.Tela, []byte(Desc))
	if err != nil {
		println(err)
	}
	return err

}

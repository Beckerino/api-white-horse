package pagar

import (
	"encoding/json"
	"fmt"
	"log"
)

func pagarRead() (result []Receber, err error) {
	query := `select uuid, 
		tipoconta_id,
		nome,
		valor::money::numeric::float8,
		valorpago::money::numeric::float8,
		datavenc,
		datapagamento,
		situacao 
		from public.contapagar`
	db, err := createDB()
	if err != nil {
		log.Fatalln(err)
	}
	err = db.Select(&result, query)
	if err != nil {
		log.Fatalln(err)
	}

	auditor := new(Log)
	auditor.Tela = "Receber/Read"

	err = auditoria(auditor, result)
	if err != nil {
		fmt.Println("error:", err)
	}

	return result, err
}

func pagarCreate(data *Receber) (result string, err error) {
	query := `INSERT INTO public.contapagar
		(nome, tipoconta_id, valor, valorpago, datavenc, datapagamento, situacao)
		VALUES($1, $2, $3, $4, $5, $6, $7);
			`
	db, err := createDB()
	if err != nil {
		log.Fatalln(err)
	}
	_, err = db.Exec(query, data.Nome, data.Tipoconta, data.Valor, data.Valorpago, data.Datavenc, data.Datapagamento, data.Situacao)
	if err != nil {
		log.Fatalln(err)
		return
	}

	auditor := new(Log)
	auditor.Tela = "Receber/Create"

	err = auditoria(auditor, data)
	if err != nil {
		fmt.Println("error:", err)
	}
	result = "Item Criado : " + data.ID
	return result, err
}

func pagarUpdate(data *Receber) (result string, err error) {

	query := `UPDATE public.contapagar SET nome=$1, tipoconta_id=$2, valor=$3, valorpago=$4, datavenc=$5, datapagamento=$6, situacao=$7, uuid=$8 
	where uuid = $8
;
			`
	db, err := createDB()
	if err != nil {
		log.Fatalln(err)
	}
	_, err = db.Exec(query, data.Nome, data.Tipoconta, data.Valor, data.Valorpago, data.Datavenc, data.Datapagamento, data.Situacao, data.ID)
	if err != nil {
		log.Fatalln(err)
		return
	}

	auditor := new(Log)
	auditor.Tela = "Receber/Update"

	err = auditoria(auditor, data)
	if err != nil {
		fmt.Println("error:", err)
	}
	result = "Item Atualizado : " + data.ID
	return result, err
}

func pagarRemove(data *Receber) (result string, err error) {
	query := `DELETE FROM public.contapagar WHERE uuid = $1;`
	db, err := createDB()
	if err != nil {
		log.Fatalln(err)
	}
	_, err = db.Exec(query, data.ID)
	if err != nil {
		log.Fatalln(err)
		return
	}
	auditor := new(Log)
	auditor.Tela = "Receber/Remove"

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
		log.Fatalln(err)
	}
	Desc, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error:", err)
	}

	_, err = db.Exec(query, auditoria.Tela, []byte(Desc))
	if err != nil {
		log.Fatalln(err)
	}
	return err

}

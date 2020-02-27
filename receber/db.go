package receber

import (
	"encoding/json"
	"fmt"
)

func receberRead() (result []Receberout, err error) {
	query := `SELECT uuid,
		c.nome, 
		cpfcnpj,
		t.nome as tipo, 
		valor::money::numeric::float8,
		datareceber,
		datarecebido, 
		s.nome as situacao
		FROM public.contareceber as c
		left join tipoconta as t on t.id = tipoconta_id
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
	auditor.Tela = "Receber/Read"

	err = auditoria(auditor, result)
	if err != nil {
		Println("error:", err)
	}

	return result, err
}

func receberCreate(data *Receberin) (result string, err error) {
	query := `INSERT INTO public.contareceber
		(nome, cpfcnpj, tipoconta_id, valor, datareceber, datarecebido, situacao)
		VALUES($1, $2, $3, $4, $5, $6, $7);
;
			`
	db, err := createDB()
	if err != nil {
		println(err)
	}
	_, err = db.Exec(query, data.Nome, data.CpfCnpj, data.Tipoconta, data.Valor, data.DataReceber, data.DataRecebido, data.Situacao)
	if err != nil {
		println(err)
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

func receberUpdate(data *Receberin) (result string, err error) {

	query := `UPDATE public.contareceber SET nome=$1, cpfcnpj=$2, tipoconta_id=$3, valor=$4, datareceber=$5, datarecebido=$6, situacao=$7 
	where uuid = $8
;
			`
	db, err := createDB()
	if err != nil {
		println(err)
	}
	_, err = db.Exec(query, data.Nome, data.CpfCnpj, data.Tipoconta, data.Valor, data.DataReceber, data.DataRecebido, data.Situacao, data.ID)
	if err != nil {
		println(err)
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

func receberRemove(data *Receberin) (result string, err error) {
	query := `DELETE FROM public.contareceber WHERE uuid = $1;`
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

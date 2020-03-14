package detail

import (
	conf "dpm/conf"
	tools "dpm/models/tools"
	"fmt"
	"log"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/lib/pq"
)

type PreddogovorKaskoLeasing struct {
	FilePath    string
	summaryrows []prec_KaskoLeasingSummaryRow
}

type prec_KaskoLeasingSummaryRow struct {
	id      string
	date    string
	filial  string
	user    string
	insurer string
	vehicle string
	vin     string
	result  string
	userid  string
}

func (pk *PreddogovorKaskoLeasing) Read() {

	var current prec_KaskoLeasingSummaryRow

	var id []byte
	var date []byte
	var filial []byte
	var user []byte
	var insurer []byte
	var vehicle []byte
	var result []byte
	var userid []byte

	query := ""
	query += "SELECT "
	query += "	ЛизингАВТО.Код, ЛизингАВТО.ДатаОбращения, Филиалы.Название,  "
	query += "	ЛизингАВТО.КтоРассматривал, ЛизингАВТО.СтраховательОрганизация, [Модель] & ' ' & [Марка] & ' ' & [ГодАвто] AS машина, "
	//	query += "	Результат.Название, Сотрудники.Код, Филиалы.Код, Результат.Код "
	query += "	Результат.Название, Сотрудники.Код "
	query += "FROM "
	query += "	((ЛизингАВТО LEFT JOIN Филиалы ON ЛизингАВТО.Филиал = Филиалы.Код) "
	query += "				 LEFT JOIN Результат ON ЛизингАВТО.РезультатСогласования = Результат.Код) "
	query += "				 LEFT JOIN Сотрудники ON ЛизингАВТО.КтоРассматривал = Сотрудники.Фамилия "
	query += ";"

	rows, err := conf.DB_access.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &date, &filial, &user, &insurer, &vehicle, &result, &userid)

		if err != nil {
			log.Fatal(err)
		} else {

			current.id = string(id)
			current.date = string(date)
			current.filial = string(filial)
			current.user = string(user)
			current.insurer = string(insurer)
			current.vehicle = string(vehicle)
			current.result = string(result)
			current.userid = string(userid)

			pk.summaryrows = append(pk.summaryrows, current)
		}
	}
	rows.Close()

}

func (pk *PreddogovorKaskoLeasing) Load() {

	var lastInsertId string

	query := "DELETE FROM details.prec_kasko_leasing;"
	_, err := conf.DB_postgres.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(pk.summaryrows); i++ {

		query := "INSERT INTO details.prec_kasko_leasing "
		query += "(id, date, filial, insurer, vehicle, result, \"user\",userid)"
		query += "VALUES ("
		query += pk.summaryrows[i].id + ","
		query += tools.Query_StringOrNull(pk.summaryrows[i].date) + ", "
		query += tools.Query_StringOrNull(pk.summaryrows[i].filial) + ", "
		query += tools.Query_StringOrNull(pk.summaryrows[i].insurer) + ", "
		query += tools.Query_StringOrNull(pk.summaryrows[i].vehicle) + ", "
		query += tools.Query_StringOrNull(pk.summaryrows[i].result) + ", "
		query += tools.Query_StringOrNull(pk.summaryrows[i].user) + ", "
		query += tools.Query_StringOrNull(pk.summaryrows[i].userid) + " "
		query += ") returning id; "

		err = conf.DB_postgres.QueryRow(query).Scan(&lastInsertId)

		if err != nil {
			fmt.Println(query)
			log.Fatal(err)
		}

	}

}

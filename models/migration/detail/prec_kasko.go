package detail

import (
	conf "dpm/conf"
	tools "dpm/models/tools"
	"fmt"
	"log"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/lib/pq"
)

type PreddogovorKasko struct {
	FilePath    string
	summaryrows []prec_KaskoSummaryRow
}

type prec_KaskoSummaryRow struct {
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

func (pk *PreddogovorKasko) Read() {

	var current prec_KaskoSummaryRow

	var id []byte
	var date []byte
	var filial []byte
	var user []byte
	var insurer []byte
	var vehicle []byte
	var vin []byte
	var result []byte
	var userid []byte

	/*
		SELECT
			[Преддоговор КАСКО].Код, [Преддоговор КАСКО].ДатаОбращения, Филиалы.Название,
			[Преддоговор КАСКО].КтоРассматривал, [СтраховательФИО] & [СтраховательОрганизация] AS Страх,
			[Модель] & ' ' & [Марка] & ' ' & [ГодВыпуска] AS машина, [Преддоговор КАСКО].VIN,
			Результат.Название, Сотрудники.Код
		FROM
			(([Преддоговор КАСКО] LEFT JOIN Результат ON [Преддоговор КАСКО].РезультатСогласования = Результат.Код)
								  LEFT JOIN Филиалы ON [Преддоговор КАСКО].Филиал = Филиалы.Код)
								  LEFT JOIN Сотрудники ON [Преддоговор КАСКО].КтоРассматривал = Сотрудники.Фамилия;
	*/

	query := ""
	query += "SELECT "
	query += "	[Преддоговор КАСКО].Код, [Преддоговор КАСКО].ДатаОбращения, Филиалы.Название, "
	query += "	[Преддоговор КАСКО].КтоРассматривал, [СтраховательФИО] & [СтраховательОрганизация] AS Страх, "
	query += "	[Модель] & ' ' & [Марка] & ' ' & [ГодВыпуска] AS машина, [Преддоговор КАСКО].VIN, "
	query += "	Результат.Название, Сотрудники.Код "
	query += "FROM "
	query += "	(([Преддоговор КАСКО] LEFT JOIN Результат ON [Преддоговор КАСКО].РезультатСогласования = Результат.Код) "
	query += "						  LEFT JOIN Филиалы ON [Преддоговор КАСКО].Филиал = Филиалы.Код) "
	query += "						  LEFT JOIN Сотрудники ON [Преддоговор КАСКО].КтоРассматривал = Сотрудники.Фамилия "
	query += ";"

	rows, err := conf.DB_access.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &date, &filial, &user, &insurer, &vehicle, &vin, &result, &userid)

		if err != nil {
			log.Fatal(err)
		} else {

			current.id = string(id)
			current.date = string(date)
			current.filial = string(filial)
			current.user = string(user)
			current.insurer = string(insurer)
			current.vehicle = string(vehicle)
			current.vin = string(vin)
			current.result = string(result)
			current.userid = string(userid)

			pk.summaryrows = append(pk.summaryrows, current)
		}
	}
	rows.Close()

}

func (pk *PreddogovorKasko) Load() {

	var lastInsertId string

	query := "DELETE FROM details.prec_kasko;"
	_, err := conf.DB_postgres.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(pk.summaryrows); i++ {

		query := "INSERT INTO details.prec_kasko "
		query += "(id, date, filial, \"user\", insurer, vehicle, vin, result, userid)"
		query += "VALUES ("
		query += pk.summaryrows[i].id + ","
		query += tools.Query_StringOrNull(pk.summaryrows[i].date) + ", "
		query += tools.Query_StringOrNull(pk.summaryrows[i].filial) + ", "
		query += tools.Query_StringOrNull(pk.summaryrows[i].user) + ", "
		query += tools.Query_StringOrNull(pk.summaryrows[i].insurer) + ", "
		query += tools.Query_StringOrNull(pk.summaryrows[i].vehicle) + ", "
		query += tools.Query_StringOrNull(pk.summaryrows[i].vin) + ", "
		query += tools.Query_StringOrNull(pk.summaryrows[i].result) + ", "
		query += tools.Query_StringOrNull(pk.summaryrows[i].userid) + " "
		query += ") returning id; "

		err = conf.DB_postgres.QueryRow(query).Scan(&lastInsertId)

		if err != nil {
			fmt.Println(query)
			log.Fatal(err)
		}

	}

}

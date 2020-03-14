package detail

import (
	conf "dpm/conf"
	tools "dpm/models/tools"
	"fmt"
	"log"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/lib/pq"
)

type RegulationImushestvo struct {
	FilePath    string
	summaryrows []reg_ImushestvoSummaryRow
}

type reg_ImushestvoSummaryRow struct {
	id       string
	created  string
	filial   string
	user     string
	ins_type string
	insurer  string
	closed   string
	usherb   string
	money    string
	control  string
	krupnie  string
	moshen   string
	result   string
}

func (ri *RegulationImushestvo) Read() {

	//fmt.Println("Here")
	var current reg_ImushestvoSummaryRow
	ri.summaryrows = ri.summaryrows[:0]

	var id []byte
	var created []byte
	var filial []byte
	var user []byte
	var ins_type []byte
	var insurer []byte
	var closed []byte
	var usherb []byte
	var money []byte
	var control []byte
	var krupnie []byte
	var moshen []byte
	var result []byte

	/*
		SELECT
			Материалы.Код, Материалы.[Дата поступления], Филиалы.Название,
			Сотрудники.Фамилия, [Виды страхования].Название,
			[СтраховательФИО] & "" & [СтраховательНазвание] AS [ФИО/Организация],
			Материалы.[Дата закрытия], Материалы.[Сумма ущерба],
			Материалы.ФинРезультат, Материалы.Контроль
		FROM (
				(
					Материалы LEFT JOIN Сотрудники ON Материалы.[Текущий исполнитель] = Сотрудники.Код)
					LEFT JOIN Филиалы ON Материалы.Направлен = Филиалы.Код)
				LEFT JOIN [Виды страхования] ON Материалы.[Вид страхования] = [Виды страхования].Код;
	*/

	//Материалы.[Важный материал], Материалы.[Признаки мошенничества]

	query := ""
	query += "SELECT "
	query += "	Материалы.Код, Материалы.[Дата поступления], Филиалы.Название, "
	query += "	Сотрудники.Фамилия, [Виды страхования].Название, "
	query += "	[СтраховательФИО] & [СтраховательНазвание] AS [ФИО/Организация], "
	query += "	Материалы.[Дата закрытия], Материалы.[Сумма ущерба], "
	query += "	Материалы.ФинРезультат, Материалы.Контроль, "
	query += "	Материалы.[Важный материал], Материалы.[Признаки мошенничества], Материалы.[Отказ/минимизация] "
	query += "FROM 	("
	query += "			( "
	query += "				Материалы LEFT JOIN Сотрудники ON Материалы.[Текущий исполнитель] = Сотрудники.Код) "
	query += "				LEFT JOIN Филиалы ON Материалы.Направлен = Филиалы.Код) "
	query += "			LEFT JOIN [Виды страхования] ON Материалы.[Вид страхования] = [Виды страхования].Код "
	query += ";"

	//fmt.Println(query)

	rows, err := conf.DB_access.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("next")

	for rows.Next() {
		err := rows.Scan(&id, &created, &filial, &user, &ins_type, &insurer, &closed, &usherb, &money, &control, &krupnie, &moshen, &result)

		if err != nil {
			log.Fatal(err)
		} else {

			current.id = string(id)
			current.created = string(created)
			current.filial = string(filial)
			current.user = string(user)
			current.ins_type = string(ins_type)
			current.insurer = string(insurer)
			current.closed = string(closed)
			current.usherb = string(usherb)
			current.money = string(money)
			current.control = string(control)
			current.krupnie = string(krupnie)
			current.moshen = string(moshen)
			current.result = string(result)

			ri.summaryrows = append(ri.summaryrows, current)
		}
	}
	rows.Close()

}

func (ri *RegulationImushestvo) Load() {

	var lastInsertId string

	query := "DELETE FROM details.reg_imushestvo;"
	_, err := conf.DB_postgres.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(ri.summaryrows); i++ {

		query := "INSERT INTO details.reg_imushestvo "
		query += "(id, created, filial, \"user\", ins_type, insurer, closed, usherb, money, control, krupnie, moshen, result)"
		query += "VALUES ("
		query += ri.summaryrows[i].id + ","
		query += tools.Query_StringOrNull(ri.summaryrows[i].created) + ", "
		query += tools.Query_StringOrNull(ri.summaryrows[i].filial) + ", "
		query += tools.Query_StringOrNull(ri.summaryrows[i].user) + ", "
		query += tools.Query_StringOrNull(ri.summaryrows[i].ins_type) + ", "
		query += tools.Query_StringOrNull(ri.summaryrows[i].insurer) + ", "
		query += tools.Query_StringOrNull(ri.summaryrows[i].closed) + ", "

		if ri.summaryrows[i].usherb == "" {
			query += "NULL ,"
		} else {
			query += ri.summaryrows[i].usherb + ",  "
		}

		if ri.summaryrows[i].money == "" {
			query += "NULL ,"
		} else {
			query += ri.summaryrows[i].money + ",  "
		}

		query += tools.Query_StringOrNull(ri.summaryrows[i].control) + ", "
		query += tools.Query_StringOrNull(ri.summaryrows[i].krupnie) + ", "
		query += tools.Query_StringOrNull(ri.summaryrows[i].moshen) + ", "
		query += tools.Query_StringOrNull(ri.summaryrows[i].result) + " "
		query += ") returning id; "

		err = conf.DB_postgres.QueryRow(query).Scan(&lastInsertId)

		if err != nil {
			fmt.Println(query)
			log.Fatal(err)
		}

	}

}

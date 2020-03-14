package detail

import (
	conf "dpm/conf"
	tools "dpm/models/tools"
	"fmt"
	"log"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/lib/pq"
)

type Debitorka struct {
	FilePath    string
	summaryrows []debitorkaRow
}

type debitorkaRow struct {
	id      string
	created string
	filial  string
	user    string
	debtor  string
	money   string
	closed  string
	control string
	debp    string
}

func (d *Debitorka) Read() {

	var current debitorkaRow
	d.summaryrows = d.summaryrows[:0]

	var id []byte
	var created []byte
	var filial []byte
	var user []byte
	var debtor []byte
	var money []byte
	var closed []byte
	var control []byte
	var debp []byte

	query := ""
	query += "SELECT "
	query += "	ДЗ.[Код], "
	query += "	ДЗ.[Дата работы в ДПМ], "
	query += "	ДЗ.[Филиал],  "
	query += "	ДЗ.[Исполнитель], "
	query += "	ДЗ.[ФИО] & ДЗ.[Наименование],  "
	query += "	ROUND(ДЗ.[Фин результат(денежный)],0), "
	query += "	ДЗ.[Закрыт], "
	query += "	ДЗ.[Контроль], "
	query += "	ДЗ.[Сумма долга] "
	query += "FROM 	ДЗ;"

	rows, err := conf.DB_access.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("next")

	for rows.Next() {
		err := rows.Scan(&id, &created, &filial, &user, &debtor, &money, &closed, &control, &debp)

		if err != nil {
			log.Fatal(err)
		} else {

			current.id = string(id)
			current.created = string(created)
			current.filial = string(filial)
			current.user = string(user)
			current.debtor = string(debtor)
			current.money = string(money)
			current.closed = string(closed)
			current.control = string(control)
			current.debp = string(debp)

			d.summaryrows = append(d.summaryrows, current)
		}
	}
	rows.Close()

}

func (d *Debitorka) Load() {

	var lastInsertId string

	query := "DELETE FROM details.debtor;"
	_, err := conf.DB_postgres.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	/*
		INSERT INTO details.debtor(
		id, created, filial, "user", debtor, money, closed, control)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);
	*/
	for i := 0; i < len(d.summaryrows); i++ {

		query := "INSERT INTO details.debtor "
		query += "(id, created, filial, \"user\", debtor, money, debt, closed, control)"
		query += "VALUES ("
		query += d.summaryrows[i].id + ","
		query += tools.Query_StringOrNull(d.summaryrows[i].created) + ", "
		query += tools.Query_StringOrNull(d.summaryrows[i].filial) + ", "
		query += tools.Query_StringOrNull(d.summaryrows[i].user) + ", "
		query += tools.Query_StringOrNull(d.summaryrows[i].debtor) + ", "

		if d.summaryrows[i].money == "" {
			query += "NULL ,"
		} else {
			query += d.summaryrows[i].money + ",  "

		}

		if d.summaryrows[i].debp == "" {
			query += "NULL ,"
		} else {
			query += d.summaryrows[i].debp + ",  "

		}

		query += tools.Query_StringOrNull(d.summaryrows[i].closed) + ", "
		query += tools.Query_StringOrNull(d.summaryrows[i].control) + " "
		query += ") returning id; "

		err = conf.DB_postgres.QueryRow(query).Scan(&lastInsertId)

		if err != nil {
			fmt.Println(query)
			log.Fatal(err)
		}

	}

}

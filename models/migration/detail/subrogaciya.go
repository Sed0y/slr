package detail

import (
	conf "dpm/conf"
	tools "dpm/models/tools"
	"fmt"
	"log"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/lib/pq"
)

type Subrogaciya struct {
	FilePath    string
	summaryrows []sub_Row
}

type sub_Row struct {
	id      string
	created string
	filial  string
	user    string
	debtor  string
	money   string
	closed  string
	control string
}

func (s *Subrogaciya) Read() {

	var current sub_Row
	s.summaryrows = s.summaryrows[:0]

	var id []byte
	var created []byte
	var filial []byte
	var user []byte
	var debtor []byte
	var money []byte
	var closed []byte
	var control []byte

	/*
	   SELECT
	   	Суброгация.Код,
	   	Суброгация.[Дата работы в ДПМ],
	   	Суброгация.Филиал,
	   	Суброгация.[Сотрудник дирекции],
	   	Суброгация.Виновник,
	   	Суброгация.[Финансовый результат (денежный)],
	   	Суброгация.Закрыт,
	   	Суброгация.Контроль
	   FROM
	   	Суброгация;

	*/

	query := ""
	query += "SELECT "
	query += "	Суброгация.Код,  "
	query += "	Суброгация.[Дата работы в ДПМ], "
	query += "	Суброгация.Филиал,   "
	query += "	Суброгация.[Сотрудник дирекции], "
	query += "	Суброгация.Виновник,  "
	query += "	ROUND(Суброгация.[Финансовый результат (денежный)],0), "
	query += "	Суброгация.Закрыт, "
	query += "	Суброгация.Контроль  "
	query += "FROM 	Суброгация;"

	rows, err := conf.DB_access.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("next")

	for rows.Next() {
		err := rows.Scan(&id, &created, &filial, &user, &debtor, &money, &closed, &control)

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

			s.summaryrows = append(s.summaryrows, current)
		}
	}
	rows.Close()

}

func (s *Subrogaciya) Load() {

	var lastInsertId string

	query := "DELETE FROM details.subrogaciya;"
	_, err := conf.DB_postgres.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(s.summaryrows); i++ {

		query := "INSERT INTO details.subrogaciya "
		query += "(id, created, filial, \"user\", debtor, money, closed, control)"
		query += "VALUES ("
		query += s.summaryrows[i].id + ","
		query += tools.Query_StringOrNull(s.summaryrows[i].created) + ", "
		query += tools.Query_StringOrNull(s.summaryrows[i].filial) + ", "
		query += tools.Query_StringOrNull(s.summaryrows[i].user) + ", "
		query += tools.Query_StringOrNull(s.summaryrows[i].debtor) + ", "

		if s.summaryrows[i].money == "" {
			query += "NULL ,"
		} else {
			query += s.summaryrows[i].money + ",  "
		}
		query += tools.Query_StringOrNull(s.summaryrows[i].closed) + ", "
		query += tools.Query_StringOrNull(s.summaryrows[i].control) + " "

		query += ") returning id; "

		err = conf.DB_postgres.QueryRow(query).Scan(&lastInsertId)

		if err != nil {
			fmt.Println(query)
			log.Fatal(err)
		}

	}

}

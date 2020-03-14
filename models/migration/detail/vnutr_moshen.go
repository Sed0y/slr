package detail

import (
	conf "dpm/conf"
	tools "dpm/models/tools"
	"fmt"
	"log"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/lib/pq"
)

type VnutrenneeMoshennichestvo struct {
	FilePath    string
	summaryrows []debitorkaRow
}

type vm_Row struct {
	id      string
	created string
	filial  string
	user    string
	debtor  string
	money   string
	closed  string
	control string
}

func (v *VnutrenneeMoshennichestvo) Read() {

	var current debitorkaRow
	v.summaryrows = v.summaryrows[:0]

	var id []byte
	var created []byte
	var filial []byte
	var user []byte
	var debtor []byte
	var money []byte
	var closed []byte
	var control []byte

	query := ""
	query += "SELECT "
	query += "	Материалы_Внутренний_Мошеничество.[Код],  "
	query += "	Материалы_Внутренний_Мошеничество.[Дата поступл в ДПМ], "
	query += "	Материалы_Внутренний_Мошеничество.[Филиал],   "
	query += "	Материалы_Внутренний_Мошеничество.[Сотрудник дирекции], "
	query += "	Материалы_Внутренний_Мошеничество.[Должность нарушителя], "
	query += "	ROUND(Материалы_Внутренний_Мошеничество.[Финансовый результат (денежный)],0), "
	query += "	Материалы_Внутренний_Мошеничество.[Закрыт], "
	query += "	Материалы_Внутренний_Мошеничество.[Контроль] "
	query += "FROM 	Материалы_Внутренний_Мошеничество;"

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

			v.summaryrows = append(v.summaryrows, current)
		}
	}
	rows.Close()

}

func (v *VnutrenneeMoshennichestvo) Load() {

	var lastInsertId string

	query := "DELETE FROM details.vm;"
	_, err := conf.DB_postgres.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(v.summaryrows); i++ {

		query := "INSERT INTO details.vm "
		query += "(id, created, filial, \"user\", debtor, money, closed, control)"
		query += "VALUES ("
		query += v.summaryrows[i].id + ","
		query += tools.Query_StringOrNull(v.summaryrows[i].created) + ", "
		query += tools.Query_StringOrNull(v.summaryrows[i].filial) + ", "
		query += tools.Query_StringOrNull(v.summaryrows[i].user) + ", "
		query += tools.Query_StringOrNull(v.summaryrows[i].debtor) + ", "

		if v.summaryrows[i].money == "" {
			query += "NULL ,"
		} else {
			query += v.summaryrows[i].money + ",  "
		}
		query += tools.Query_StringOrNull(v.summaryrows[i].closed) + ", "
		query += tools.Query_StringOrNull(v.summaryrows[i].control) + " "
		query += ") returning id; "

		err = conf.DB_postgres.QueryRow(query).Scan(&lastInsertId)

		if err != nil {
			fmt.Println(query)
			log.Fatal(err)
		}

	}

}

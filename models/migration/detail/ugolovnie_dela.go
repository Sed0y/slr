package detail

import (
	conf "dpm/conf"
	tools "dpm/models/tools"
	"fmt"
	"log"

	_ "github.com/alexbrainman/odbc"
	_ "github.com/lib/pq"
)

type UgolovnieDela struct {
	FilePath    string
	summaryrows []ud_Row
}

type ud_Row struct {
	id                     string
	created                string
	filial                 string
	user                   string
	debtor                 string
	money                  string
	closed                 string
	control                string
	ins_type               string
	zayavlenie_date        string
	material_date          string
	material_reshenie_date string
	vozbujdenie_date       string
	ud_reshsenie_date      string
	reshenie               string
	reshenie_date          string
	vozbujdenie_sd_date    string
	reshenie_ud            string
	opg                    string
}

func (u *UgolovnieDela) Read() {

	var current ud_Row
	u.summaryrows = u.summaryrows[:0]

	var id []byte
	var created []byte
	var filial []byte
	var user []byte
	var debtor []byte
	var money []byte
	var closed []byte
	var control []byte
	var ins_type []byte
	var zayavlenie_date []byte        //Дата отправления заявления
	var material_date []byte          //Дата КУСП
	var material_reshenie_date []byte //ДатаРешенияПоМатериалу
	var vozbujdenie_date []byte       //Дата возбуждения УД
	var ud_reshsenie_date []byte      //ДатаРешенияПоУД
	var reshenie []byte               //Решение по СД
	var reshenie_UD []byte            //Решение по УД
	var reshenie_date []byte          //Дата решения по СД
	var vozbujdenie_sd_date []byte    //Дата возбуждения СД
	var opg []byte                    //ОПГ

	query := ""
	query += "SELECT "
	query += "	Материалы_Внутренний_контроль.[Код],  "
	query += "	Материалы_Внутренний_контроль.[Дата работы в ДПМ], "
	query += "	Материалы_Внутренний_контроль.[Филиал],   "
	query += "	Материалы_Внутренний_контроль.[Сотрудник дирекции], "
	query += "	Материалы_Внутренний_контроль.[Данные о мошеннике], "
	query += "	ROUND(Материалы_Внутренний_контроль.[Финансовый результат (денежный)],0), "
	query += "	Null AS Выражение1, "
	query += "	Материалы_Внутренний_контроль.[Контроль],  "
	query += "	Материалы_Внутренний_контроль.[Вид страхования], "
	query += "	Материалы_Внутренний_контроль.[Дата отправления заявления], "
	query += "	Материалы_Внутренний_контроль.[Дата возбуждения УД], "
	query += "	Материалы_Внутренний_контроль.[Решение по СД], "
	query += "	Материалы_Внутренний_контроль.[Дата решения по СД], "
	query += "	Материалы_Внутренний_контроль.[Дата КУСП], "
	query += "	Материалы_Внутренний_контроль.[ДатаРешенияПоМатериалу], "
	query += "	Материалы_Внутренний_контроль.[ДатаРешенияПоУД], "
	query += "	Материалы_Внутренний_контроль.[Дата возбуждения СД], "
	query += "	Материалы_Внутренний_контроль.[Решение по материалу], "
	query += "	Материалы_Внутренний_контроль.[ОПГ] "
	query += "FROM 	Материалы_Внутренний_контроль;"

	rows, err := conf.DB_access.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("next")

	for rows.Next() {
		err := rows.Scan(&id, &created, &filial, &user, &debtor, &money,
			&closed, &control, &ins_type, &zayavlenie_date, &vozbujdenie_date,
			&reshenie, &reshenie_date, &material_date, &material_reshenie_date,
			&ud_reshsenie_date, &vozbujdenie_sd_date, &reshenie_UD, &opg)

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
			current.ins_type = string(ins_type)
			current.zayavlenie_date = string(zayavlenie_date)
			current.vozbujdenie_date = string(vozbujdenie_date)
			current.reshenie = string(reshenie)
			current.reshenie_date = string(reshenie_date)
			current.material_date = string(material_date)
			current.material_reshenie_date = string(material_reshenie_date)
			current.ud_reshsenie_date = string(ud_reshsenie_date)
			current.vozbujdenie_sd_date = string(vozbujdenie_sd_date)
			current.reshenie_ud = string(reshenie_UD)
			current.opg = string(opg)

			u.summaryrows = append(u.summaryrows, current)
		}
	}
	rows.Close()

}

func (u *UgolovnieDela) Load() {

	var lastInsertId string

	query := "DELETE FROM details.ud;"
	_, err := conf.DB_postgres.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(u.summaryrows); i++ {

		query := "INSERT INTO details.ud "
		query += "(id, created, filial, \"user\", debtor, money, closed, control, ins_type, zayavlenie_date, vozbujdenie_date, reshenie, reshenie_date, material_date, material_reshenie_date, ud_reshsenie_date, vozbujdenie_sd_date, reshenie_ud ,opg )"
		query += "VALUES ("
		query += u.summaryrows[i].id + ","
		query += tools.Query_StringOrNull(u.summaryrows[i].created) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].filial) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].user) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].debtor) + ", "

		if u.summaryrows[i].money == "" {
			query += "NULL ,"
		} else {
			query += u.summaryrows[i].money + ",  "
		}
		query += tools.Query_StringOrNull(u.summaryrows[i].closed) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].control) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].ins_type) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].zayavlenie_date) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].vozbujdenie_date) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].reshenie) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].reshenie_date) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].material_date) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].material_reshenie_date) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].ud_reshsenie_date) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].vozbujdenie_sd_date) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].reshenie_ud) + ", "
		query += tools.Query_StringOrNull(u.summaryrows[i].opg) + " "
		query += ") returning id; "

		err = conf.DB_postgres.QueryRow(query).Scan(&lastInsertId)

		if err != nil {
			fmt.Println(query)
			log.Fatal(err)
		}

	}

}

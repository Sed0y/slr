package entities

import (
	"database/sql"
	"fmt"
	"log"
	"solaris/conf"
	"strconv"

	_ "github.com/lib/pq"
)

type Filial struct {
	Id int

	Name    string
	Curator int // идентификатор филиала, который его курирует

	EmployerDPM bool //предусмотрена должность ДПМ

}

type FilialList struct {
	filials []Filial
}

func (fl *FilialList) Get() []Filial {
	return fl.filials
}

func (fl *FilialList) query_AllFilials() string {
	//return "SELECT Код, Название, [Должность ПМ] FROM [Филиалы] ORDER BY Название"
	return "SELECT id, name, curator, is_dpm FROM public.filial order by name;"
}

func (fl *FilialList) Item(index int) Filial {
	return fl.filials[index]
}

func (fl *FilialList) Count() int {
	return len(fl.filials)
}

func (fl *FilialList) Clear() {
	fl.filials = fl.filials[:0]
}

func (fl *FilialList) Load() {

	fl.Clear()

	var id []byte
	var name []byte
	var curator []byte
	var job []byte

	var current Filial

	rows, err := conf.DB_postgres.Query(fl.query_AllFilials())
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &name, &curator, &job)

		if err != nil {
			log.Fatal(err)
		} else {
			current.Id, err = strconv.Atoi(string(id))
			current.Name = string(name)
			c, _ := strconv.Atoi(string(curator))
			current.Curator = c
			current.EmployerDPM = false
			if string(job) == "Предусмотрена" {
				current.EmployerDPM = true
			}

			fl.filials = append(fl.filials, current)
		}
	}

	defer rows.Close()
}

func (fl *FilialList) GetById(id int) (bool, Filial) {

	var nothing Filial
	for i := 0; i < fl.Count(); i++ {
		if fl.Item(i).Id == id {
			return true, fl.Item(i)
		}
	}
	return false, nothing
}

func (fl *FilialList) MigrateToPostgres() {

	var id string
	var lastInsertId string

	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		"postgres", "vmmjn28z", "journal")

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	for i := 0; i < fl.Count(); i++ {
		id = strconv.Itoa(fl.Item(i).Id)

		query := "INSERT INTO filial"
		query += "(id, name, is_dpm)"
		query += "VALUES ("
		query += id + ","
		query += "'" + fl.Item(i).Name + "', "
		query += strconv.FormatBool(fl.Item(i).EmployerDPM) + " "
		query += ") returning id; "

		//fmt.Println(query)
		err = db.QueryRow(query).Scan(&lastInsertId)

		if err != nil {
			log.Fatal(err)
		}
	}

}

func (fl *FilialList) GetListToReport(id int) []Filial {
	var fils []Filial

	_, f := fl.GetById(id)

	if f.Name == "ГК" {

		var all_fil Filial
		all_fil.Name = "ВСЕ"
		all_fil.Id = -1

		fils = append(fils, all_fil)
		all := fl.Get()

		for f := 0; f < len(all); f++ {
			fils = append(fils, all[f])
		}

	} else {

		fils = append(fils, f)

		for i := 0; i < fl.Count(); i++ {
			if fl.filials[i].Curator == id {
				_, f = fl.GetById(fl.filials[i].Id)
				fils = append(fils, f)
			}
		}
	}

	return fils
}

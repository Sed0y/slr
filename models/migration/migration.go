package migration

import (
	"log"
	conf "solaris/conf"

	_ "github.com/lib/pq"
)

type Migration struct {
	AccessTables []accessTable
}

type accessTable struct {
	Id          string
	Name        string
	Description string
	DB_name     string
	DB_path     string
	Last_update string
}

func (s *Migration) Load() {
	s.GetAccessTables()
}

// Получить список обслуживаемых таблиц журнала
//
//
func (s *Migration) GetAccessTables() string {

	var tbl accessTable
	var result string
	result = ""

	s.AccessTables = s.AccessTables[:0]

	var id []byte
	var name []byte
	var description []byte
	var db_name []byte
	var db_path []byte
	var last_update []byte

	query := "SELECT id, name, description, db_name, db_path, to_char(last_update, 'dd.mm.YYYY HH24:MI') as last_update  from system.access_tables;"

	rows, err := conf.DB_postgres.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &db_name, &db_path, &last_update)
		if err != nil {
			log.Fatal(err)
		} else {
			tbl.Id = string(id)
			tbl.Name = string(name)
			tbl.Description = string(description)
			tbl.DB_name = string(db_name)
			tbl.DB_path = string(db_path)
			tbl.Last_update = string(last_update)
			s.AccessTables = append(s.AccessTables, tbl)
		}
	}

	rows.Close()

	return result
}

func (s *Migration) Updated(id string) {
	conf.DB_postgres.Exec("UPDATE system.access_tables SET last_update = current_timestamp WHERE id = " + id + ";")
}

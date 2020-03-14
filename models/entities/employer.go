package entities

import (
	"solaris/conf"
	//	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

type Employer struct {
	Id int

	Login string

	Surname    string
	Name       string
	FatherName string

	Password string
	EFilial  Filial

	Role string
}

func (e *Employer) isRoot() bool {

	if e.Login == conf.Root {
		return true
	}

	return false
}

type EmployerList struct {
	employers []Employer
}

func (el *EmployerList) Item(index int) Employer {
	return el.employers[index]
}

func (el *EmployerList) Count() int {
	return len(el.employers)
}

func (el *EmployerList) Clear() {
	el.employers = el.employers[:0]
}

func (el *EmployerList) Get() []Employer {
	return el.employers
}

// Загружает данные о сотрудниках (пользователях) из БД Postgres
// процесс обновлеия данных с аксессом возложен на модуль migration
// поэтому класс рабоает только с БД Postgres
//
func (el *EmployerList) Load() {

	var id string

	var login []byte
	var surname []byte
	var name []byte
	var fathername []byte
	var password []byte
	var role []byte

	var fil_name []byte
	var fil_id []byte

	var current Employer
	var current_fil Filial

	el.Clear()
	// id, login, surname,  name, fathername, password, role, fid
	rows, err := conf.DB_postgres.Query("SELECT u.id, u.login, u.surname,  u.name, u.fathername, u.password,u.fid, u.role, f.name	FROM public.users as u left join public.filial as f on u.fid = f.id;")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&id, &login, &surname, &name, &fathername, &password, &fil_id, &role, &fil_name)

		if err != nil {
			log.Fatal(err)
		} else {
			current.Id, err = strconv.Atoi(id)
			current.Surname = string(surname)
			current.Login = string(login) // !!! Логин по фамилии. надо переделывать
			current.Name = string(name)
			current.FatherName = string(fathername)
			current.Password = string(password)
			current.Role = string(role)

			current_fil.Id, _ = strconv.Atoi(string(fil_id))
			current_fil.Name = string(fil_name)

			current.EFilial = current_fil
			el.employers = append(el.employers, current)
		}
	}

	defer rows.Close()
}

// Возвращает экземпляр по идентификатору
func (el *EmployerList) GetById(id int) (bool, Employer) {

	var nothing Employer
	for i := 0; i < el.Count(); i++ {
		if el.Item(i).Id == id {
			return true, el.Item(i)
		}
	}
	return false, nothing
}

// Авторизацию пользователя
// проверяет пару логин-пароль
// возвращает:
// - идентификатор роли сотрудника
// если сотрудник не найден, то возвращает -1
func (el *EmployerList) CheckAuth(name string, pass string) (int, int) {

	role_id := -1
	user_id := -1

	for i := 0; i < el.Count(); i++ {
		if el.Item(i).Login == name && el.Item(i).Password == pass {
			role_id, _ = strconv.Atoi(el.Item(i).Role)
			user_id = el.Item(i).Id
		}
	}

	return user_id, role_id
}

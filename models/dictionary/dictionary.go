package dictionary

//"github.com/astaxie/beego"

//"database/sql"
//"dpm/conf"
//"log"
//"strconv"

import (
	"time"
)

// ВИД СТРАХОВАНИЯ
// на данный момент есть несколько источников,
// поэтому просто строка
//
// надо сделать единый список и тогда уже можно будет
// его обрабатывать
type TypeOfInsurance struct {
	id   int
	name string
}

func (t *TypeOfInsurance) Set(value string) bool {

	if value == "КАСКО" || value == "ОСАГО" || value == "ГО и ЗК" {
		t.name = value
		return true
	}

	if value == "KASKO" {
		t.name = "КАСКО"
		return true
	}

	if value == "OSAGO" {
		t.name = "ОСАГО"
		return true
	}

	if value == "GO" {
		t.name = "ГО и ЗК"
		return true
	}

	return false
}

func (t *TypeOfInsurance) Get() string {
	return t.name
}

func (t *TypeOfInsurance) IsKASKO() bool {
	if t.name == "КАСКО" || t.name == "KASKO" {
		return true
	}
	return false
}

func (t *TypeOfInsurance) IsOSAGO() bool {
	if t.name == "ОСАГО" || t.name == "OSAGO" {
		return true
	}
	return false
}

func (t *TypeOfInsurance) IsGO() bool {
	if t.name == "ГО и ЗК" || t.name == "GO" {
		return true
	}
	return false
}

// ВИД ДЕЯТЕЛЬНОСТИ
// PreContract - Преддоговорная работа
// Regulation - Урегулирование
type TypeOfJob struct {
	id   int
	Name string
}

func (t *TypeOfJob) IsPreContract() bool {

	if t.Name == "PreContract" {
		return true
	}

	return false
}

func (t *TypeOfJob) IsRegulation() bool {

	if t.Name == "Regulation" {
		return true
	}

	return false
}

// ПЕРИОД
//
type Period struct {
	Begin time.Time
	End   time.Time
}

// ТИП РЕЗУЛЬТАТА
//
// Поступившие, отказы и т.п.
//
type TypeOfResult struct {
	Name string
}

func (t *TypeOfResult) IsAll() bool {
	if t.Name == "All" {
		return true
	}
	return false
}

func (t *TypeOfResult) IsRefusal() bool {
	if t.Name == "Refusal" {
		return true
	}
	return false
}

func (t *TypeOfResult) IsMinimization() bool {
	if t.Name == "Minimization" {
		return true
	}
	return false
}

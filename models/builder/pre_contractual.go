package builder

import (
	conf "dpm/conf"
	Dictionary "dpm/models/dictionary"
	Entities "dpm/models/entities"
	"fmt"
	"log"
	"strconv"
)

type PreContractBuilder struct {
	query_PreContractBuilder
}

func (pc *PreContractBuilder) DoNothing() {

}

// Преддоговорная работа
// Результаты конкретного сотрудника
//
// employer - сотрудник по которому нужно получить результаты
// type_of_insurance - вид страхования
// period - период
//
// возвращает Result
//
//
func (pc *PreContractBuilder) GetEmployerResult(
	employer *Entities.Employer,
	type_of_insurance *Dictionary.TypeOfInsurance,
	period *Dictionary.Period) Result {

	var res Result

	if type_of_insurance.IsKASKO() {
		res = pc.getEmployerResult_KASKO(employer, period)
	}
	res.Id = employer.Id

	return res
}

func (pc *PreContractBuilder) getEmployerResult_KASKO(
	employer *Entities.Employer,
	period *Dictionary.Period) Result {

	var res Result
	var All int
	var Refusal int
	var result []byte

	All = -1
	Refusal = -1

	fmt.Println("AutoPostupilo")

	query := pc.query_PreContractBuilder.AutoPostupilo(
		period.Begin.Format("2006-01-02"),
		period.End.Format("2006-01-02"),
		strconv.Itoa(employer.Id))

	fmt.Println(query)

	rows, err := conf.DB_postgres.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			log.Fatal(err)
		} else {
			All, _ = strconv.Atoi(string(result))
		}
	}

	rows.Close()

	fmt.Println("AutoRefusal")

	query = pc.query_PreContractBuilder.AutoRefusal(
		period.Begin.Format("2006-01-02"),
		period.End.Format("2006-01-02"),
		strconv.Itoa(employer.Id))

	rows, err = conf.DB_postgres.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(&result)
		if err != nil {
			log.Fatal(err)
		} else {
			Refusal, _ = strconv.Atoi(string(result))
		}
	}

	rows.Close()

	res.All = All
	res.Refusal = Refusal

	return res
}

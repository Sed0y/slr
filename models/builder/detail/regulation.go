package detail

import (
	conf "dpm/conf"
	Dictionary "dpm/models/dictionary"
	Entities "dpm/models/entities"
	"fmt"
	"log"
	"strconv"

	_ "github.com/alexbrainman/odbc"
)

type RegulationDetails struct {
	query_RegulationDetails
}

func (pc *RegulationDetails) DoNothing() {

}

func (pc *RegulationDetails) GetEmployerDetails(
	type_of_result *Dictionary.TypeOfResult,
	employer *Entities.Employer,
	type_of_insurance *Dictionary.TypeOfInsurance,
	period *Dictionary.Period) DetailResults {

	var res DetailResults

	res = pc.getEmployerDetails(type_of_insurance, employer, period, type_of_result)

	return res
}

func (pc *RegulationDetails) getEmployerDetails(
	type_of_insurance *Dictionary.TypeOfInsurance,
	employer *Entities.Employer,
	period *Dictionary.Period,
	type_of_result *Dictionary.TypeOfResult) DetailResults {

	var res DetailResults
	var columns []string

	query := ""

	if type_of_result.IsAll() {
		query = pc.AutoPostupilo(
			type_of_insurance.Get(),
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(employer.Id))
	}

	if type_of_result.IsRefusal() {
		query = pc.AutoRefusal(
			type_of_insurance.Get(),
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(employer.Id))
	}

	if type_of_result.IsMinimization() {
		query = pc.AutoMinimize(
			type_of_insurance.Get(),
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(employer.Id))
	}

	fmt.Println(query)

	rows, err := conf.DB_postgres.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	//var current []string
	columns, _ = rows.Columns()

	fmt.Println(columns)

	count := len(columns)
	//	values := make([]interface{}, count)

	pointers := make([]interface{}, count)
	container := make([][]byte, count)

	for _, val := range columns {
		res.header = append(res.header, val)
	}
	fmt.Println(res.header)

	for i, _ := range pointers {
		pointers[i] = &container[i]
	}
	//	var result [][]string

	for rows.Next() {
		err := rows.Scan(pointers...)
		if err != nil {
			log.Fatal(err)
		} else {

			//tmp := make([]byte, len(container))
			var tmp_str []string

			for i := 0; i < len(container); i++ {
				tmp_str = append(tmp_str, string(container[i]))
			}

			//copy(tmp, container)
			res.data = append(res.data, tmp_str)
		}
	}

	rows.Close()
	fmt.Println(res.data)
	return res
}

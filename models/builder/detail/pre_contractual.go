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

type PreContractDetails struct {
	query_PreContractDetails
}

func (pc *PreContractDetails) DoNothing() {

}

func (pc *PreContractDetails) GetEmployerDetails(
	type_of_result *Dictionary.TypeOfResult,
	employer *Entities.Employer,
	type_of_insurance *Dictionary.TypeOfInsurance,
	period *Dictionary.Period) DetailResults {

	var res DetailResults

	if type_of_insurance.IsKASKO() {
		res = pc.getEmployerDetails_KASKO(employer, period, type_of_result)
	}

	return res
}

func (pc *PreContractDetails) getEmployerDetails_KASKO(
	employer *Entities.Employer,
	period *Dictionary.Period,
	type_of_result *Dictionary.TypeOfResult) DetailResults {

	var res DetailResults
	var columns []string

	query := ""

	if type_of_result.IsAll() {
		query = pc.query_PreContractDetails.AutoAll(
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(employer.Id))
	}

	if type_of_result.IsRefusal() {
		query = pc.query_PreContractDetails.AutoRefusal(
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(employer.Id))
	}

	//fmt.Println(query)

	rows, err := conf.DB_postgres.Query(query)
	if err != nil {
		log.Fatal(err)
	}

	//var current []string
	columns, _ = rows.Columns()

	//fmt.Println(columns)

	count := len(columns)
	//	values := make([]interface{}, count)

	pointers := make([]interface{}, count)
	container := make([][]byte, count)

	for _, val := range columns {
		res.header = append(res.header, val)
	}
	//fmt.Println(res.header)

	for i, _ := range pointers {
		pointers[i] = &container[i]
	}

	for rows.Next() {
		err := rows.Scan(pointers...)
		if err != nil {
			log.Fatal(err)
		} else {

			var tmp_str []string
			for i := 0; i < len(container); i++ {
				tmp_str = append(tmp_str, string(container[i]))
			}
			res.data = append(res.data, tmp_str)
		}
	}

	rows.Close()
	fmt.Println(res.data)
	return res
}

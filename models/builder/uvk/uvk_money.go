package uvk

import (
	conf "dpm/conf"
	general "dpm/models/builder"
	dictionary "dpm/models/dictionary"
	entities "dpm/models/entities"
	"encoding/json"
	"fmt"
	"strconv"
)

type Results struct {
	Debitorka                general.Result
	VnutrnneeMoshennichestvo general.Result
	UgolovnieDela            general.Result
	Subrogaciya              general.Result
}

func (res *Results) ToJSON() string {

	jsonResult, _ := json.Marshal(res)
	return string(jsonResult)
}

type Money struct {
}

func (r *Money) GetFilialResults(
	filial *entities.Filial,
	period *dictionary.Period,
	curator bool) Results {

	var qu queries
	var q string
	var scan_result []byte

	var results Results

	var Debitorka general.Result
	var VnutrnneeMoshennichestvo general.Result
	var UgolovnieDela general.Result
	var Subrogaciya general.Result

	// ****************************************
	// 1. Получить результаты по Дебиторке ****
	// ****************************************

	if true {

		q = qu.FilialMoney.Debitorka(
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err := conf.DB_postgres.Query(q)

		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					Debitorka.Financial, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}
	}

	// ********************************************************
	// 2. Получить результаты по Внутреннему мошенничеству ****
	// ********************************************************

	if true {

		q = qu.FilialMoney.VnutrenneeMoshennichestvo(
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err := conf.DB_postgres.Query(q)

		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					VnutrnneeMoshennichestvo.Financial, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}

	}

	// **********************************************
	// 3. Получить результаты по Уголовным делам ****
	// **********************************************

	if true {

		q = qu.FilialMoney.UgolovnieDela(
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err := conf.DB_postgres.Query(q)

		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					UgolovnieDela.Financial, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}
	}

	// *****************************************
	// 4. Получить результаты по Суброгации ****
	// *****************************************

	if true {

		q = qu.FilialMoney.Subrogaciya(
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err := conf.DB_postgres.Query(q)

		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					Subrogaciya.Financial, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}
	}

	results.Debitorka = Debitorka
	results.VnutrnneeMoshennichestvo = VnutrnneeMoshennichestvo
	results.UgolovnieDela = UgolovnieDela
	results.Subrogaciya = Subrogaciya

	return results

}

func (r *Money) GetFilialResultsDetails(
	filial *entities.Filial,
	period *dictionary.Period,
	curator bool,
	debtor_type string) general.DetailResults {

	var qu queries
	var res general.DetailResults
	var columns []string

	query := ""

	fmt.Println(query)
	switch debtor_type {

	case "Debitorka":
		query = qu.FilialMoney.DebitorkaDetails(
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)
	case "Subrogaciya":
		query = qu.FilialMoney.SubrogaciyaDetails(
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)
	case "VnutrnneeMoshennichestvo":
		query = qu.FilialMoney.VnutrenneeMoshennichestvoDetails(
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)
	case "UgolovnieDela":
		query = qu.FilialMoney.UgolovnieDelaDetails(
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)
	default:
		fmt.Println("??")
	}

	rows, err := conf.DB_postgres.Query(query)
	if err != nil {
		fmt.Println(err)
	}

	columns, _ = rows.Columns()
	count := len(columns)

	pointers := make([]interface{}, count)
	container := make([][]byte, count)

	for _, val := range columns {
		res.Header = append(res.Header, val)
	}

	for i, _ := range pointers {
		pointers[i] = &container[i]
	}

	for rows.Next() {
		err := rows.Scan(pointers...)
		if err != nil {
			fmt.Println(err)
		} else {
			var tmp_str []string
			for i := 0; i < len(container); i++ {
				tmp_str = append(tmp_str, string(container[i]))
			}
			res.Data = append(res.Data, tmp_str)
		}
	}

	rows.Close()
	//fmt.Println(res.Data)
	return res

}

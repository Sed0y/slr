package upma

import (
	conf "dpm/conf"
	general "dpm/models/builder"
	dictionary "dpm/models/dictionary"
	entities "dpm/models/entities"
	"fmt"
	//"log"
	"encoding/json"
	"strconv"
)

type Results struct {
	KASKO general.Result
	OSAGO general.Result
	GO    general.Result
	ZK    general.Result
}

func (res *Results) ToJSON() string {

	jsonResult, _ := json.Marshal(res)
	return string(jsonResult)
}

type Regulation struct {
	q_builder query_Regulation
}

func (r *Regulation) GetFilialResults(
	filial *entities.Filial,
	period *dictionary.Period,
	curator bool) Results {

	var qu queries
	var q string
	var scan_result []byte

	var results Results
	var KASKO general.Result
	var OSAGO general.Result
	var GO general.Result
	var ZK general.Result

	// ************************************
	// 1. Получить результаты по КАСКО ****
	// ************************************

	if true {

		q = qu.FilialRegulation.Postupilo(
			"КАСКО",
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
					KASKO.All, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}

		q = qu.FilialRegulation.Refusal(
			"КАСКО",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					KASKO.Refusal, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.Minimize(
			"КАСКО",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					KASKO.Minimization, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.RefusalFinance(
			"КАСКО",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					KASKO.RefusalFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.MinimizeFinance(
			"КАСКО",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					KASKO.MinimizationFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

	}

	// ************************************
	// 2. Получить результаты по ОСАГО ****
	// ************************************

	if true {

		q = qu.FilialRegulation.Postupilo(
			"ОСАГО",
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
					OSAGO.All, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}

		//fmt.Println(q)

		q = qu.FilialRegulation.Refusal(
			"ОСАГО",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					OSAGO.Refusal, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.Minimize(
			"ОСАГО",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					OSAGO.Minimization, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.RefusalFinance(
			"ОСАГО",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					OSAGO.RefusalFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.MinimizeFinance(
			"ОСАГО",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					OSAGO.MinimizationFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

	}

	// ************************************
	// 3. Получить результаты по ГО ****
	// ************************************

	if true {

		q = qu.FilialRegulation.Postupilo(
			"ГО",
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
					GO.All, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}

		//fmt.Println(q)

		q = qu.FilialRegulation.Refusal(
			"ГО",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					GO.Refusal, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.Minimize(
			"ГО",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					GO.Minimization, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.RefusalFinance(
			"ГО",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					GO.RefusalFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.MinimizeFinance(
			"ГО",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					GO.MinimizationFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

	}

	// ************************************
	// 4. Получить результаты по ЗК ****
	// ************************************

	if true {

		q = qu.FilialRegulation.Postupilo(
			"ЗК",
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
					ZK.All, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}

		//fmt.Println(q)

		q = qu.FilialRegulation.Refusal(
			"ЗК",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					ZK.Refusal, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.Minimize(
			"ЗК",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					ZK.Minimization, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.RefusalFinance(
			"ЗК",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					ZK.RefusalFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.MinimizeFinance(
			"ЗК",
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)

		rows, err = conf.DB_postgres.Query(q)
		if err != nil {
			fmt.Println(err)
		} else {
			for rows.Next() {
				err := rows.Scan(&scan_result)
				if err != nil {
					fmt.Println(err)
				} else {
					ZK.MinimizationFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

	}

	results.KASKO = KASKO
	results.OSAGO = OSAGO
	results.GO = GO
	results.ZK = ZK

	results.KASKO.Financial = KASKO.MinimizationFinancialResult + KASKO.RefusalFinancialResult
	results.OSAGO.Financial = OSAGO.MinimizationFinancialResult + OSAGO.RefusalFinancialResult
	results.GO.Financial = GO.MinimizationFinancialResult + GO.RefusalFinancialResult
	results.ZK.Financial = ZK.MinimizationFinancialResult + ZK.RefusalFinancialResult

	results.KASKO.RefusalAndMinimization = KASKO.Refusal + KASKO.Minimization
	results.OSAGO.RefusalAndMinimization = OSAGO.Refusal + OSAGO.Minimization
	results.GO.RefusalAndMinimization = GO.Refusal + GO.Minimization
	results.ZK.RefusalAndMinimization = ZK.Refusal + ZK.Minimization

	return results
}

/*
* minimized - если true, то выдаёт отказы=минимизации
*             если false, то поступившие
 */
func (r *Regulation) GetFilialResultsDetails(
	filial *entities.Filial,
	period *dictionary.Period,
	curator bool,
	ins_type string,
	minimized bool) general.DetailResults {

	var qu queries
	var res general.DetailResults
	var columns []string

	query := ""

	/*
	   q = qu.FilialRegulation.Postupilo(
	   			"ЗК",
	   			period.Begin.Format("2006-01-02"),
	   			period.End.Format("2006-01-02"),
	   			strconv.Itoa(filial.Id),
	   			curator)
	*/
	if minimized {
		query = qu.FilialRegulation.RefusalMinimizeDetails(
			ins_type,
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)
	} else {
		query = qu.FilialRegulation.PostupiloDetails(
			ins_type,
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)
	}

	//fmt.Println(query)

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

/*

// Работа на этапе урегулирования
// Результаты конкретного сотрудника
//
// employer - сотрудник по которому нужнонужно получить результаты
// type_of_insurance - вид страхования
// period = период
//
// возвращает result.Result
//
//
func (r *Regulation) GetEmployerResult(
	employer *entities.Employer,
	type_of_insurance *dictionary.TypeOfInsurance,
	period *dictionary.Period) result.Result {

	var res result.Result

	fmt.Println("GetEmployerResult")

	fmt.Println("GetEmployerResult IsKASKO")
	if type_of_insurance.IsKASKO() {
		res = r.getEmployerAutoResult(type_of_insurance, employer, period)
	}

	fmt.Println("GetEmployerResult IsOSAGO")
	if type_of_insurance.IsOSAGO() {
		res = r.getEmployerAutoResult(type_of_insurance, employer, period)
	}

	fmt.Println("GetEmployerResult IsGO")
	if type_of_insurance.IsGO() {
		res = r.getEmployerAutoResult(type_of_insurance, employer, period)
	}

	res.Id = employer.Id

	return res
}

// Урегулирование по КАСКО
//
func (r *Regulation) getEmployerAutoResult(
	type_of_insurance *dictionary.TypeOfInsurance,
	employer *entities.Employer,
	period *dictionary.Period) result.Result {

	var res result.Result

	var All int            //всего отработано
	var Refusal int        //отказы
	var Minimized int      //минимизировано
	var RefusalMoney int   //отказы
	var MinimizedMoney int //отказы


		//var Finance int   //сумма

	var result []byte

	All = -1
	Refusal = -1
	Minimized = -1

	query := r.q_builder.AutoPostupilo(
		type_of_insurance.Get(),
		period.Begin.Format("2006-01-02"),
		period.End.Format("2006-01-02"),
		strconv.Itoa(employer.Id))

	rows, err := conf.DB_postgres.Query(query)
	if err != nil {
		fmt.Println(err)
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

	query = r.q_builder.AutoRefusal(
		type_of_insurance.Get(),
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

	query = r.q_builder.AutoMinimize(
		type_of_insurance.Get(),
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
			Minimized, _ = strconv.Atoi(string(result))
		}
	}

	rows.Close()

	query = r.q_builder.AutoRefusalFinance(
		type_of_insurance.Get(),
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
			RefusalMoney, _ = strconv.Atoi(string(result))
		}
	}

	rows.Close()

	query = r.q_builder.AutoMinimizeFinance(
		type_of_insurance.Get(),
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
			MinimizedMoney, _ = strconv.Atoi(string(result))
		}
	}

	rows.Close()

	res.All = All
	res.Refusal = Refusal
	res.Minimization = Minimized
	res.RefusalFinancialResult = RefusalMoney
	res.MinimizationFinancialResult = MinimizedMoney
	res.Financial = RefusalMoney + MinimizedMoney

	return res
}

*/

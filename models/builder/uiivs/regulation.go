package uiivs

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
	ImushestvoUL   general.Result
	ImushestvoFL   general.Result
	NS             general.Result
	Gruz           general.Result
	Otvetstvennost general.Result
	Spectehnika    general.Result
	Titul          general.Result
	Others         general.Result
}

func (res *Results) ToJSON() string {

	jsonResult, _ := json.Marshal(res)
	return string(jsonResult)
}

type Regulation struct {
}

func (r *Regulation) GetFilialResults(
	filial *entities.Filial,
	period *dictionary.Period,
	curator bool) Results {

	var qu queries
	var q string
	var scan_result []byte

	var results Results

	var ImushestvoUL general.Result
	var ImushestvoFL general.Result
	var Otvetstvennost general.Result
	var Gruz general.Result
	var NS general.Result
	var Spectehnika general.Result
	var Titul general.Result
	var Others general.Result

	/*
		1	Имущество ЮЛ
		2	Имущество ФЛ
		3	Ответственность
		4	Грузы
		5	НС
		6	Госконтракты		true
		7	Иные				true
		8	Титул
		9	Спецтехника
		10	ВЗР					true
		12	Авиа- и косм- риски	true
	*/

	// ************************************
	// 1. Получить результаты по Имущество ЮЛ ****
	// ************************************
	//fmt.Println("Получить результаты по Имущество Ю")
	if true {

		q = qu.FilialRegulation.Postupilo(
			"Имущество ЮЛ",
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
					ImushestvoUL.All, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}

		q = qu.FilialRegulation.Refusal(
			"Имущество ЮЛ",
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
					ImushestvoUL.Refusal, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		fmt.Println("Получить результаты по Имущество Ю - 3")

		q = qu.FilialRegulation.Minimize(
			"Имущество ЮЛ",
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
					ImushestvoUL.Minimization, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.RefusalFinance(
			"Имущество ЮЛ",
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
					ImushestvoUL.RefusalFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.MinimizeFinance(
			"Имущество ЮЛ",
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
					ImushestvoUL.MinimizationFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

	}

	// ************************************
	// 2. Получить результаты по Имущество ФЛ ****
	// ************************************

	if true {

		q = qu.FilialRegulation.Postupilo(
			"Имущество ФЛ",
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
					ImushestvoFL.All, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}

		q = qu.FilialRegulation.Refusal(
			"Имущество ФЛ",
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
					ImushestvoFL.Refusal, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.Minimize(
			"Имущество ФЛ",
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
					ImushestvoFL.Minimization, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.RefusalFinance(
			"Имущество ФЛ",
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
					ImushestvoFL.RefusalFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.MinimizeFinance(
			"Имущество ФЛ",
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
					ImushestvoFL.MinimizationFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

	}

	// ************************************
	// 3. Получить результаты по Ответственность ****
	// ************************************

	if true {

		q = qu.FilialRegulation.Postupilo(
			"Ответственность",
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
					Otvetstvennost.All, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}

		q = qu.FilialRegulation.Refusal(
			"Ответственность",
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
					Otvetstvennost.Refusal, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.Minimize(
			"Ответственность",
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
					Otvetstvennost.Minimization, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.RefusalFinance(
			"Ответственность",
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
					Otvetstvennost.RefusalFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.MinimizeFinance(
			"Ответственность",
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
					Otvetstvennost.MinimizationFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

	}

	// ************************************
	// 4. Получить результаты по Грузы ****
	// ************************************

	if true {

		q = qu.FilialRegulation.Postupilo(
			"Грузы",
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
					Gruz.All, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}

		q = qu.FilialRegulation.Refusal(
			"Грузы",
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
					Gruz.Refusal, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.Minimize(
			"Грузы",
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
					Gruz.Minimization, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.RefusalFinance(
			"Грузы",
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
					Gruz.RefusalFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.MinimizeFinance(
			"Грузы",
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
					Gruz.MinimizationFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

	}

	// ************************************
	// 5. Получить результаты по НС ****
	// ************************************

	if true {

		q = qu.FilialRegulation.Postupilo(
			"НС",
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
					NS.All, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}

		//fmt.Println(q)

		q = qu.FilialRegulation.Refusal(
			"НС",
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
					NS.Refusal, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.Minimize(
			"НС",
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
					NS.Minimization, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.RefusalFinance(
			"НС",
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
					NS.RefusalFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.MinimizeFinance(
			"НС",
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
					NS.MinimizationFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

	}

	// ************************************
	// 6. Получить результаты по Спецтехника ****
	// ************************************

	if true {

		q = qu.FilialRegulation.Postupilo(
			"Спецтехника",
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
					Spectehnika.All, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}

		q = qu.FilialRegulation.Refusal(
			"Спецтехника",
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
					Spectehnika.Refusal, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.Minimize(
			"Спецтехника",
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
					Spectehnika.Minimization, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.RefusalFinance(
			"Спецтехника",
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
					Spectehnika.RefusalFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.MinimizeFinance(
			"Спецтехника",
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
					Spectehnika.MinimizationFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

	}

	// ************************************
	// 7. Получить результаты по Титул ****
	// ************************************

	if true {

		q = qu.FilialRegulation.Postupilo(
			"Титул",
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
					Titul.All, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}

		q = qu.FilialRegulation.Refusal(
			"Титул",
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
					Titul.Refusal, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.Minimize(
			"Титул",
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
					Titul.Minimization, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.RefusalFinance(
			"Титул",
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
					Titul.RefusalFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.MinimizeFinance(
			"Титул",
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
					Titul.MinimizationFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

	}

	// ************************************
	// 8. Получить результаты по Остальным ****
	// ************************************

	if true {

		q = qu.FilialRegulation.Postupilo(
			"Иные",
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
					Others.All, _ = strconv.Atoi(string(scan_result))
				}
			}

			rows.Close()
		}

		q = qu.FilialRegulation.Refusal(
			"Иные",
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
					Others.Refusal, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.Minimize(
			"Иные",
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
					Others.Minimization, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.RefusalFinance(
			"Иные",
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
					Others.RefusalFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

		q = qu.FilialRegulation.MinimizeFinance(
			"Иные",
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
					Others.MinimizationFinancialResult, _ = strconv.Atoi(string(scan_result))
				}
			}
		}

		rows.Close()

	}

	results.ImushestvoUL = ImushestvoUL
	results.ImushestvoFL = ImushestvoFL
	results.NS = NS
	results.Gruz = Gruz
	results.Otvetstvennost = Otvetstvennost
	results.Spectehnika = Spectehnika
	results.Titul = Titul
	results.Others = Others

	results.ImushestvoUL.Financial = ImushestvoUL.MinimizationFinancialResult + ImushestvoUL.RefusalFinancialResult
	results.ImushestvoFL.Financial = ImushestvoFL.MinimizationFinancialResult + ImushestvoFL.RefusalFinancialResult
	results.NS.Financial = NS.MinimizationFinancialResult + NS.RefusalFinancialResult
	results.Gruz.Financial = Gruz.MinimizationFinancialResult + Gruz.RefusalFinancialResult
	results.Otvetstvennost.Financial = Otvetstvennost.MinimizationFinancialResult + Otvetstvennost.RefusalFinancialResult
	results.Spectehnika.Financial = Spectehnika.MinimizationFinancialResult + Spectehnika.RefusalFinancialResult
	results.Titul.Financial = Titul.MinimizationFinancialResult + Titul.RefusalFinancialResult
	results.Others.Financial = Others.MinimizationFinancialResult + Others.RefusalFinancialResult

	results.ImushestvoUL.RefusalAndMinimization = ImushestvoUL.Refusal + ImushestvoUL.Minimization
	results.ImushestvoFL.RefusalAndMinimization = ImushestvoFL.Refusal + ImushestvoFL.Minimization
	results.NS.RefusalAndMinimization = NS.Refusal + NS.Minimization
	results.Gruz.RefusalAndMinimization = Gruz.Refusal + Gruz.Minimization
	results.Otvetstvennost.RefusalAndMinimization = Otvetstvennost.Refusal + Otvetstvennost.Minimization
	results.Spectehnika.RefusalAndMinimization = Spectehnika.Refusal + Spectehnika.Minimization
	results.Titul.RefusalAndMinimization = Titul.Refusal + Titul.Minimization
	results.Others.RefusalAndMinimization = Others.Refusal + Others.Minimization

	return results
}

/*
 * restype - тип результата - поступившие, отказы или минимизация
 * */
func (r *Regulation) GetFilialResultsDetails(
	filial *entities.Filial,
	period *dictionary.Period,
	curator bool,
	ins_type string,
	restype string) general.DetailResults {

	var qu queries
	var res general.DetailResults
	var columns []string

	query := ""

	if restype == "All" {
		query = qu.FilialRegulation.PostupiloDetails(
			ins_type,
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)
	} else if restype == "Refusal" {
		query = qu.FilialRegulation.RefusalDetails(
			ins_type,
			period.Begin.Format("2006-01-02"),
			period.End.Format("2006-01-02"),
			strconv.Itoa(filial.Id),
			curator)
	} else if restype == "Minimization" {
		query = qu.FilialRegulation.MinimizeDetails(
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

	return res

}

package builder

import (
	"encoding/json"
	"strings"
)

/*
* Количественные результаты
* Id - идентификатор филиала или сотрудника
* All - сколько поступилов в работу
* Refusal - сколько отказано
* Minimization - сколько минимизировано
* RefusalAndMinimization - минимизировано и отказано вместе
* RefusalFinancialResult - сумма по отказам
* MinimizationFinancialResult - сумма по минимизации
* Financial - итоговая сумма минимизации
*
 */
type Result struct {
	Id                          int
	All                         int
	Refusal                     int
	Minimization                int
	RefusalAndMinimization      int
	RefusalFinancialResult      int
	MinimizationFinancialResult int
	Financial                   int
}

func (er *Result) ToJSON() string {

	jsonResult, _ := json.Marshal(er)
	return string(jsonResult)
}

/*
* детализация результатов
* header - названия столбцов
* data - данные
*
*	   {
*	     "header": [
*	       "1",
*	       "2",
*	       "3"
*	     ],
*	     "data": [
*	       [
*	         "1",
*	         "2",
*	         "3"
*	       ],
*	       [
*	         "4",
*	         "5",
*	         "6"
*	       ]
*	     ]
*	   }
*
*
 */
type DetailResults struct {
	Header []string
	Data   [][]string
}

func (r *DetailResults) ToJSON() string {

	var result string
	result = ""

	result += "{"
	result += "\"header\":"
	result += "["
	for i, val := range r.Header {
		if i == len(r.Header)-1 {
			result += "\"" + string(val) + "\""
		} else {
			result += "\"" + string(val) + "\","
		}
	}
	result += "],"

	result += "\"data\":"
	result += "["
	for i, val := range r.Data {
		result += "["

		for j, row := range val {
			if j == len(val)-1 {
				result += "\"" + ClearData(string(row)) + "\""
			} else {
				result += "\"" + ClearData(string(row)) + "\","
			}
		}

		if i == len(r.Data)-1 {
			result += "]"
		} else {
			result += "],"
		}
	}
	result += "]"

	result += "}"

	return result
}

func ClearData(value string) string {
	result := strings.Replace(value, "\"", "", -1)
	return result
}

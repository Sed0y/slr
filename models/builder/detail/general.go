package detail

import (
	//	"database/sql"

	_ "fmt"
	"strings"
	//	"dpm/models/dictionary"
	//	"dpm/models/entities"
)

type DetailResults struct {
	header []string
	data   [][]string
}

func (r *DetailResults) ToJSON() string {

	var result string
	result = ""
	/*
	   {
	     "header": [
	       "1",
	       "2",
	       "3"
	     ],
	     "data": [
	       [
	         "1",
	         "2",
	         "3"
	       ],
	       [
	         "4",
	         "5",
	         "6"
	       ]
	     ]
	   }
	*/

	result += "{"
	result += "\"header\":"
	result += "["
	for i, val := range r.header {
		if i == len(r.header)-1 {
			result += "\"" + string(val) + "\""
		} else {
			result += "\"" + string(val) + "\","
		}
	}
	result += "],"

	result += "\"data\":"
	result += "["
	for i, val := range r.data {
		result += "["

		for j, row := range val {
			if j == len(val)-1 {
				result += "\"" + ClearData(string(row)) + "\""
			} else {
				result += "\"" + ClearData(string(row)) + "\","
			}
		}

		if i == len(r.data)-1 {
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

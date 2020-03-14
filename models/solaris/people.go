package solaris

import (
	//conf "dpm/conf"
	//	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	//"time"

	_ "github.com/lib/pq"
)

type Information struct {
	Data []string `json:"D"`
}

type People struct {
	Surname        string
	Name           string
	FatherName     string
	BirthdateDay   string
	BirthdateMonth string
	BirthdateYear  string
	Inn            string

	Response       string
	ParcedResponse []PeopleResponse
}

type PeopleResponse struct {
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`

	Telephone string `json:"telephone"`
	Car       string `json:"car"`
	Passport  string `json:"passport"`

	DayBirth   int `json:"dayBirth"`
	MonthBirth int `json:"monthBirth"`
	YearBirth  int `json:"yearBirth"`

	Snils string `json:"snils"`
	Inn   string `json:"inn"`

	Information string `json:"information"`

	Base   string `json:"base"`
	BaseId int    `json:"baseId"`
	Id     int64  `json:"id"`

	Base_schema Schema
}

/*
func (ppl *People) RequestByFio(service_url string, proxy string, proxy_auth string, token string) bool {

	if ppl.Surname == "" ||
		ppl.Name == "" ||
		ppl.FatherName == "" ||
		ppl.Birthdate == "" {

		fmt.Println("Solaris:RequestByFio: Не задан один из параметров")
		ppl.Response = "#error#wrong parametres"
		return false
	}

	bd, err := time.Parse("02.01.2006", ppl.Birthdate)

	if err != nil {
		fmt.Println("Solaris:RequestByFio: Некорректная дата")
		ppl.Response = "#error#wrong parametres"
		return false
	}

	proxyURL, err := url.Parse(proxy)
	if err != nil {
		fmt.Println("Solaris:RequestByFio: Ошибка прокси сервера ", err)
		return false
	}

	Url, err := url.Parse(service_url)
	if err != nil {
		fmt.Println("Solaris:RequestByFio: ошибка URL сервиса ", err)
		return false
	}

	Url.Path += api_GetPeople
	parameters := url.Values{}
	parameters.Add("lastName", ppl.Surname)
	parameters.Add("firstName", ppl.Name)
	parameters.Add("middleName", ppl.FatherName)
	parameters.Add("yearBirth", strconv.Itoa(bd.Year()))
	parameters.Add("monthBirth", strconv.Itoa(int(bd.Month())))
	parameters.Add("dayBirth", strconv.Itoa(bd.Day()))
	parameters.Add("token", token)

	Url.RawQuery = parameters.Encode()

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(proxy_auth))

	transport.ProxyConnectHeader = http.Header{}
	transport.ProxyConnectHeader.Add("Proxy-Authorization", basicAuth)

	client := &http.Client{
		Transport: transport,
	}
	request, req_err := http.NewRequest("GET", Url.String(), nil)
	if req_err != nil {
		fmt.Println("Solaris:RequestByFio: Ошибка запроса ", req_err)
		return false
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Solaris:RequestByFio: Ошибка ответа ", err)
		return false
	}

	if response.StatusCode != 200 {
		fmt.Println("Solaris:GetSchema: Ошибка запроса ", response.StatusCode)
		fmt.Println("Solaris:GetSchema: Ответ сервера ", response)
		return false
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Solaris:GetSchema: Ошибка чтения данных ответа сервера ", err)
		return false
	}

	ppl.Response = string(data)
	//fmt.Println(ppl.Response)
	json.Unmarshal([]byte(ppl.Response), &ppl.ParcedResponse)

	return true
}
*/

func (ppl *People) FullRequest(service_url string, proxy string, proxy_auth string, token string) bool {

	check_parameters := false

	if ppl.Inn != "" {
		check_parameters = true
	}

	if ppl.Surname != "" && ppl.Name != "" {
		check_parameters = true
	}

	if ppl.Surname != "" && ppl.BirthdateYear != "" {
		check_parameters = true
	}

	if ppl.Name != "" && ppl.FatherName != "" && ppl.BirthdateDay != "" && ppl.BirthdateMonth != "" && ppl.BirthdateYear != "" {
		check_parameters = true
	}

	if !check_parameters {

		fmt.Println("Solaris:People FullRequest: Недостаточно параметров для поиска")
		ppl.Response = "#error#not enogth parametres"
		return false
	}
	/*
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			fmt.Println("Solaris:People FullRequest: Ошибка прокси сервера ", err)
			ppl.Response = "#error#proxy error: " + err.Error()
			return false
		}
	*/
	Url, err := url.Parse(service_url)
	if err != nil {
		fmt.Println("Solaris:People FullRequest: ошибка URL сервиса ", err)
		ppl.Response = "#error#service url error: " + err.Error()
		return false
	}

	Url.Path += api_GetPeople
	parameters := url.Values{}

	fmt.Println("*****  Составление запроса *****")
	fmt.Println(ppl)

	if ppl.Surname != "" {
		parameters.Add("lastName", ppl.Surname)
	}

	if ppl.Name != "" {
		parameters.Add("firstName", ppl.Name)
	}

	if ppl.FatherName != "" {
		parameters.Add("middleName", ppl.FatherName)
	}

	if ppl.BirthdateYear != "" {
		parameters.Add("yearBirth", ppl.BirthdateYear)
	}

	if ppl.BirthdateMonth != "" {
		parameters.Add("monthBirth", ppl.BirthdateMonth)
	}

	if ppl.BirthdateDay != "" {
		parameters.Add("dayBirth", ppl.BirthdateDay)
	}

	if ppl.Inn != "" {
		parameters.Add("inn", ppl.Inn)
	}

	parameters.Add("token", token)

	Url.RawQuery = parameters.Encode()

	fmt.Println(Url)
	/*
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	*/
	//	basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(proxy_auth))
	/*
		transport.ProxyConnectHeader = http.Header{}
		transport.ProxyConnectHeader.Add("Proxy-Authorization", basicAuth)
	*/
	client := &http.Client{}

	request, req_err := http.NewRequest("GET", Url.String(), nil)
	if req_err != nil {
		fmt.Println("Solaris:People FullRequest: Ошибка запроса ", req_err)
		ppl.Response = "#error#request error: " + err.Error()
		return false
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Solaris:People FullRequest: Ошибка ответа ", err)
		ppl.Response = "#error#response error: " + err.Error()
		return false
	}

	if response.StatusCode != 200 {
		fmt.Println("Solaris:People FullRequest: Ошибка запроса ", response.StatusCode)
		fmt.Println("Solaris:People FullRequest: Ответ сервера ", response)
		ppl.Response = "#error#response error: " + string(response.Status)
		return false
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Solaris:People FullRequest: Ошибка чтения данных ответа сервера ", err)
		ppl.Response = "#error#response error: Ошибка чтения данных ответа сервера - " + err.Error()
		return false
	}

	if string(data) == "[]" {
		fmt.Println("Solaris:People FullRequest: Пустой ответ ", err)
		ppl.Response = "nothing"
		return true
	}

	ppl.Response = string(data)
	json.Unmarshal([]byte(ppl.Response), &ppl.ParcedResponse)

	return true
}

func (ppl *People) RenderResults(sl SchemaList) string {

	var html string
	var odd bool

	odd = true
	//html = "<div>Start</div></br>"

	for index, value := range ppl.ParcedResponse {

		html += "<table id=\"res_table_" + strconv.Itoa(index) + "\">"

		html += "   <thead>"
		html += "      <tr>"
		html += "         <th class = \"base-name\" colspan=\"2\">(" + strconv.Itoa(index+1) + ")  " + value.Base + "</th>"
		html += "      </tr>"
		html += "      <tr>"
		html += "         <th>Поле</th>"
		html += "         <th>Значение</th>"
		html += "      </tr>"
		html += "   </thead>"

		html += "   <tbody>"
		html += "      <tr class = \"even\">"
		html += "         <td class = \"field\">ФИО</td><td>" + value.LastName + " " +
			value.FirstName + " " +
			value.MiddleName + " "

		if value.DayBirth < 10 {
			html += "0" + strconv.Itoa(value.DayBirth) + "."
		} else {
			html += strconv.Itoa(value.DayBirth) + "."
		}

		if value.MonthBirth < 10 {
			html += "0" + strconv.Itoa(value.MonthBirth) + "."
		} else {
			html += strconv.Itoa(value.MonthBirth) + "."
		}

		html += strconv.Itoa(value.YearBirth) + "</td>"
		html += "      </tr>"

		if (value.Passport != "") && (len(value.Passport) > 7) {

			if odd {
				html += "      <tr class = \"odd\">"
			} else {
				html += "      <tr class = \"even\">"
			}
			odd = !odd
			html += "         <td class = \"field\">Паспорт</td><td>" + value.Passport + "</td>"
			html += "      </tr>"

		}

		if value.Inn != "" {

			if odd {
				html += "      <tr class = \"odd\">"
			} else {
				html += "      <tr class = \"even\">"
			}
			odd = !odd

			html += "         <td class = \"field\">ИНН</td><td>" + value.Inn + "</td>"
			html += "      </tr>"
		}

		if value.Snils != "" {

			if odd {
				html += "      <tr class = \"odd\">"
			} else {
				html += "      <tr class = \"even\">"
			}
			odd = !odd

			html += "         <td class = \"field\">Снилс</td><td>" + value.Snils + "</td>"
			html += "      </tr>"
		}

		if value.Telephone != "" {

			if odd {
				html += "      <tr class = \"odd\">"
			} else {
				html += "      <tr class = \"even\">"
			}
			odd = !odd

			html += "         <td class = \"field\">Телефон</td><td>" + value.Telephone + "</td>"
			html += "      </tr>"
		}

		if value.Car != "" {

			if odd {
				html += "      <tr class = \"odd\">"
			} else {
				html += "      <tr class = \"even\">"
			}
			odd = !odd

			html += "         <td class = \"field\">Автомобиль</td><td>" + value.Car + "</td>"
			html += "      </tr>"
		}

		html += "      <tr class = \"to_referat\">"
		html += "      <td></td><td></td>"
		html += "      </tr>"
		/*
			html += "      <tr>"
			html += "         <td>Реферат</td><td>" + value.Information + "</td>"
			html += "      </tr>"

			html += "      <tr>"
			html += "         <td>Схема</td><td>" + value.Base_schema.Value + "</td>"
			html += "      </tr>"
		*/
		var sch_fields ParcedSchema
		var sch_values Information

		json.Unmarshal([]byte(value.Base_schema.Value), &sch_fields)
		json.Unmarshal([]byte(value.Information), &sch_values)

		if len(sch_fields.Fields) == len(sch_values.Data) {
			for index2, value2 := range sch_fields.Fields {

				if strings.TrimSpace(sch_values.Data[index2]) == "" {
					continue
				}

				if odd {
					html += "      <tr class = \"odd\">"
				} else {
					html += "      <tr class = \"even\">"
				}
				odd = !odd

				html += "         <td class = \"field\">" +
					strings.ToUpper(value2[:2]) +
					strings.ToLower(value2[2:]) + "</td><td>" + sch_values.Data[index2] + "</td>"
				html += "      </tr>"

			}
		} else {
			// ЛОГИРОВАТЬ ЭТУ ОШИБКУ
		}

		html += "   </tbody>"
		html += "</table>"

	}

	//html += "<div>End</div></br>"

	return html

}

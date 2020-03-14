package solaris

import (
	conf "solaris/conf"
	//	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	_ "github.com/lib/pq"
)

type Schema struct {
	Table string
	Value string
}

type ParcedSchema struct {
	Table  string   `json:"table"`
	BaseId int      `json:"baseId"`
	Fields []string `json:"fields"`
}

type SchemaList struct {
	schemas []Schema
}

func (sl *SchemaList) Clear() {
	sl.schemas = sl.schemas[:0]
}

func (sl *SchemaList) Add(new_one Schema) {
	sl.schemas = append(sl.schemas, new_one)
	new_one.AddToCache()
}

func (sl *SchemaList) GetList() []Schema {
	return sl.schemas
}

func (sl *SchemaList) GetSchema(table string) Schema {

	var empty Schema
	//fmt.Println("GetSchema fom List")
	for i := 0; i < len(sl.schemas); i++ {
		if sl.schemas[i].Table == table {
			return sl.schemas[i]
		}
	}

	//fmt.Println("GetSchema fom List - empty")
	return empty
}

func (sl *SchemaList) Load() {

	sl.Clear()

	var table []byte
	var json []byte

	var current Schema

	rows, err := conf.DB_postgres.Query("SELECT table_name, \"schema\" FROM system.solaris_table_cache;")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		err := rows.Scan(&table, &json)

		if err != nil {
			fmt.Println(err)
		} else {

			current.Table = string(table)
			current.Value = string(json)

			sl.schemas = append(sl.schemas, current)
		}
	}

	defer rows.Close()

}

func (sc *Schema) CheckCache() bool {

	var scan_result []byte
	var result bool

	result = false

	query := "SELECT count(*) FROM system.solaris_table_cache "
	query += "WHERE table_name = '" + sc.Table + "';"

	rows, err := conf.DB_postgres.Query(query)

	if err != nil {
		fmt.Println(err)
	} else {
		for rows.Next() {
			err := rows.Scan(&scan_result)
			if err != nil {
				fmt.Println(err)
			} else {
				count, _ := strconv.Atoi(string(scan_result))
				if count == 1 {
					result = true
				}
			}
		}

		rows.Close()
	}

	return result

}

func (sc *Schema) GetSchema(service_url string, proxy string, proxy_auth string, token string) bool {

	if sc.Table == "" {
		fmt.Println("Solaris:GetSchema: Не задано имя таблицы")
		return false
	}
	/*
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			fmt.Println("Solaris:GetSchema: Ошибка прокси сервера ", err)
			return false
		}
	*/
	Url, err := url.Parse(service_url)
	if err != nil {
		fmt.Println("Solaris:GetSchema: ошибка URL сервиса ", err)
		return false
	}

	Url.Path += api_GetSchema
	parameters := url.Values{}
	parameters.Add("table", sc.Table)
	parameters.Add("token", token)

	Url.RawQuery = parameters.Encode()
	/*
		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyURL),
		}
	*/
	//basicAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte(proxy_auth))

	//transport.ProxyConnectHeader = http.Header{}
	//transport.ProxyConnectHeader.Add("Proxy-Authorization", basicAuth)

	client := &http.Client{
		//	Transport: transport,
	}
	request, req_err := http.NewRequest("GET", Url.String(), nil)
	if req_err != nil {
		fmt.Println("Solaris:GetSchema: Ошибка запроса ", req_err)
		return false
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Solaris:GetSchema: Ошибка ответа ", err)
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

	sc.Value = string(data)
	sc.AddToCache()
	//fmt.Println(string(data))

	return true
}

func (sc *Schema) AddToCache() bool {

	if (sc.Table == "") || (sc.Value == "") {
		fmt.Println("Solaris:AddToCache: Пустое имя или описание таблицы")
		return false
	}

	query := "DELETE FROM system.solaris_table_cache WHERE table_name = '" + sc.Table + "';"
	_, err := conf.DB_postgres.Exec(query)

	if err != nil {
		fmt.Println("Solaris:AddToCache: Ошибка удаления таблицы из кеша ", sc, err)
	}

	/*
	   INSERT INTO system.solaris_table_cache(
	   	table_name, schema)
	   	VALUES ('table2', '[{"paramm2":"Value2"}]');
	*/

	var lastInsertId string

	query = "INSERT INTO system.solaris_table_cache(table_name, schema) "
	query += "VALUES ('" + sc.Table + "', '" + sc.Value + "') returning tid;;"

	err = conf.DB_postgres.QueryRow(query).Scan(&lastInsertId)

	if err != nil {
		fmt.Println("Solaris:AddToCache: Ошибка добавления таблицы в кеш", sc, err)
	}

	return true
}

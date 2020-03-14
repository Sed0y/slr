package activity

import (
	"dpm/conf"
	"fmt"
	"strconv"
)

const (
	TableName = "log.urls"
)

type ActionLogger struct {
}

func (c *ActionLogger) LogUrl(
	userid int,
	url string,
	parameters string) {

	query := ""

	query += "INSERT INTO journal.jlogs.urls"
	query += " ( uid, url, params, dt) "
	query += "VALUES ( "
	query += " " + strconv.Itoa(userid) + ", "
	query += " '" + url + "', "
	query += " '" + parameters + "', "
	query += " current_timestamp "
	query += ");"

	conf.DB_postgres.Exec(query)

}

func (c *ActionLogger) LogSolaris(
	user_id int,
	request_type string,
	parameters string) {

	/*
		INSERT INTO jlogs.solaris(
		dt, request_type, params, uid)
		VALUES ('2019-09-09','req_1', 'par1 par2', 4);
	*/
	query := ""

	query += "INSERT INTO journal.jlogs.solaris( "
	query += "   dt, request_type, params, uid) "
	query += "VALUES ( "
	query += " current_timestamp, "
	query += " '" + request_type + "', "
	query += " '" + parameters + "', "
	query += " " + strconv.Itoa(user_id) + " "
	query += ");"

	res, err := conf.DB_postgres.Exec(query)
	if err != nil {
		fmt.Println(res, err)
	}

}

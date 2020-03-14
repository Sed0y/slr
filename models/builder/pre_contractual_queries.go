package builder

type query_PreContractBuilder struct {
}

func (q *query_PreContractBuilder) AutoPostupilo(
	date_begin string,
	date_end string,
	employer_id string) string {

	var query string

	query = ""

	query += "select count(*) from deem.prec_kasko "
	query += "where "
	query += "	employer = " + employer_id + " and "
	query += "	date between '" + date_begin + "' and '" + date_end + "' "
	query += "; "

	return query
}

func (q *query_PreContractBuilder) AutoRefusal(
	date_begin string,
	date_end string,
	employer_id string) string {

	var query string

	query = ""

	query += "select count(*) from deem.prec_kasko "
	query += "where "
	query += "	employer = " + employer_id + " and "
	query += "	date between '" + date_begin + "' and '" + date_end + "' and "
	query += "	result = 2 "
	query += "; "

	return query
}

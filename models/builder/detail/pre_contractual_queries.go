package detail

type query_PreContractDetails struct {
}

func (q *query_PreContractDetails) AutoColumns() string {

	var query string

	query = ""

	query += "SELECT column_name "
	query += "FROM information_schema.columns "
	query += "	WHERE table_schema = 'public' "
	query += "	AND table_name = 'summary_preddogovor_kasko' "
	query += "; "

	return query
}

func (q *query_PreContractDetails) AutoAll(
	date_begin string,
	date_end string,
	employer_id string) string {

	var query string

	query = ""

	query += "select id, to_char(date, 'dd.mm.YYYY') as date, filial, insurer, vehicle, vin, result, \"user\" from details.prec_kasko "
	query += "where "
	query += "	userid = " + employer_id + " and "
	query += "	date between '" + date_begin + "' and '" + date_end + "'  "
	query += "; "

	return query
}

func (q *query_PreContractDetails) AutoRefusal(
	date_begin string,
	date_end string,
	employer_id string) string {

	var query string

	query = ""

	query += "select id, to_char(date, 'dd.mm.YYYY') as date, filial, insurer, vehicle, vin, result, \"user\" from details.prec_kasko "
	query += "where "
	query += "	userid = " + employer_id + " and "
	query += "	date between '" + date_begin + "' and '" + date_end + "' and "
	query += "	result = 'отказ' "
	query += "; "

	return query
}

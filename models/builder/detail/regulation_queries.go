package detail

type query_RegulationDetails struct {
}

func (q *query_RegulationDetails) AutoPostupilo(
	ins_type string,
	date_begin string,
	date_end string,
	employer_id string) string {

	var query string

	query = ""

	query += "select * from oisuu_auto as oa "
	query += "where "
	query += "	oa.employer_trass != '' and "
	query += "	oa.trass_date between '" + date_begin + "' and '" + date_end + "' and "
	query += "	oa.zone in ( "
	query += "		select z.name from ( "
	query += "			select * from filial as f where "
	query += "				   f.id in (select u.filial from users as u where u.id = " + employer_id + ")"
	query += "				or f.curator in (select u.filial from users as u where u.id = " + employer_id + ") ) as fils "
	query += "			left join oisuu_zones as z on (fils.oisuu_zone = z.id) "
	query += "		)  and "
	query += "	oa.ins_type in ( "
	query += "		select ot.name from journal_oisuu_types as oft "
	query += "			left join ins_types as it on (oft.journal_type = it.id) "
	query += "			left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
	query += "			where it.name = '" + ins_type + "'"
	query += "	); "

	return query
}

func (q *query_RegulationDetails) AutoRefusal(
	ins_type string,
	date_begin string,
	date_end string,
	employer_id string) string {

	var query string

	query = ""

	query += "select * from oisuu_auto as oa "
	query += "where "
	query += "	oa.employer_trass != '' and "
	query += "	oa.trass_result_date between '" + date_begin + "' and '" + date_end + "' and "
	query += "	oa.trass_result = 'Отказ' and "
	query += "	oa.zone in ( "
	query += "		select z.name from ( "
	query += "			select * from filial as f where "
	query += "				   f.id in (select u.filial from users as u where u.id = " + employer_id + ")"
	query += "				or f.curator in (select u.filial from users as u where u.id = " + employer_id + ") ) as fils "
	query += "			left join oisuu_zones as z on (fils.oisuu_zone = z.id) "
	query += "		)  and "
	query += "	oa.ins_type in ( "
	query += "		select ot.name from journal_oisuu_types as oft "
	query += "			left join ins_types as it on (oft.journal_type = it.id) "
	query += "			left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
	query += "			where it.name = '" + ins_type + "'"
	query += "	); "

	return query
}

func (q *query_RegulationDetails) AutoMinimize(
	ins_type string,
	date_begin string,
	date_end string,
	employer_id string) string {

	var query string

	query = ""

	query += "select * from oisuu_auto as oa "
	query += "where "
	query += "	oa.employer_trass != '' and "
	query += "	oa.trass_result_date between '" + date_begin + "' and '" + date_end + "' and "
	query += "	oa.trass_result = 'Частичный отказ' and "
	query += "	oa.zone in ( "
	query += "		select z.name from ( "
	query += "			select * from filial as f where "
	query += "				   f.id in (select u.filial from users as u where u.id = " + employer_id + ")"
	query += "				or f.curator in (select u.filial from users as u where u.id = " + employer_id + ") ) as fils "
	query += "			left join oisuu_zones as z on (fils.oisuu_zone = z.id) "
	query += "		)  and "
	query += "	oa.ins_type in ( "
	query += "		select ot.name from journal_oisuu_types as oft "
	query += "			left join ins_types as it on (oft.journal_type = it.id) "
	query += "			left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
	query += "			where it.name = '" + ins_type + "'"
	query += "	); "

	return query
}

package upma

type queries struct {
	FilialRegulation query_FilialRegulation
	//EmployerRegulation query_EmployerRegulation
}

type query_FilialRegulation struct {
}

// ----------------------------------------------------------------------
// подстчёт количества
// ----------------------------------------------------------------------

func (q *query_FilialRegulation) Postupilo(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	if curator == true && filial_id != "-1" {

		query = ""
		query += " select count(*) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id in ( "
		query += " 													select id from public.filial where "
		query += " 														curator = " + filial_id + " union select " + filial_id + " "
		query += " 												) "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	if curator == false && filial_id != "-1" {

		query = ""
		query += " select count(*) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id = " + filial_id + " "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	if filial_id == "-1" {

		query = ""
		query += " select count(*) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	return query

}

func (q *query_FilialRegulation) Refusal(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	query = ""

	if curator == true && filial_id != "-1" {

		query = ""
		query += " select count(*) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    oa.trass_result = 'Отказ' and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id in ( "
		query += " 													select id from public.filial where "
		query += " 														curator = " + filial_id + " union select " + filial_id + " "
		query += " 												) "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	if curator == false && filial_id != "-1" {

		query = ""
		query += " select count(*) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    oa.trass_result = 'Отказ' and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id = " + filial_id + " "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	if filial_id == "-1" {

		query = ""
		query += " select count(*) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    oa.trass_result = 'Отказ' and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	/*
		query += "select count(*) from oisuu_auto as oa "
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
	*/
	return query
}

func (q *query_FilialRegulation) Minimize(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	query = ""

	if curator == true && filial_id != "-1" {

		query = ""
		query += " select count(*) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    oa.trass_result = 'Частичный отказ' and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id in ( "
		query += " 													select id from public.filial where "
		query += " 														curator = " + filial_id + " union select " + filial_id + " "
		query += " 												) "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	if curator == false && filial_id != "-1" {

		query = ""
		query += " select count(*) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    oa.trass_result = 'Частичный отказ' and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id = " + filial_id + " "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	if filial_id == "-1" {

		query = ""
		query += " select count(*) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    oa.trass_result = 'Частичный отказ' and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	return query
}

func (q *query_FilialRegulation) RefusalFinance(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	query = ""

	if curator == true && filial_id != "-1" {

		query = ""
		query += " select sum(finanse) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    oa.trass_result = 'Отказ' and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id in ( "
		query += " 													select id from public.filial where "
		query += " 														curator = " + filial_id + " union select " + filial_id + " "
		query += " 												) "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	if curator == false && filial_id != "-1" {

		query = ""
		query += " select sum(finanse) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    oa.trass_result = 'Отказ' and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id = " + filial_id + " "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	if filial_id == "-1" {

		query = ""
		query += " select sum(finanse) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    oa.trass_result = 'Отказ' and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	return query
}

func (q *query_FilialRegulation) MinimizeFinance(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	query = ""

	if curator == true && filial_id != "-1" {

		query = ""
		query += " select sum(finanse) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    oa.trass_result = 'Частичный отказ' and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id in ( "
		query += " 													select id from public.filial where "
		query += " 														curator = " + filial_id + " union select " + filial_id + " "
		query += " 												) "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	if curator == false && filial_id != "-1" {

		query = ""
		query += " select sum(finanse) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    oa.trass_result = 'Частичный отказ' and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id = " + filial_id + " "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	if filial_id == "-1" {

		query = ""
		query += " select sum(finanse) "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    oa.trass_result = 'Частичный отказ' and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	return query
}

// ----------------------------------------------------------------------
// детализация
// ----------------------------------------------------------------------

func (q *query_FilialRegulation) PostupiloDetails(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	if curator {

		query = ""
		query += " select * "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id in ( "
		query += " 													select id from public.filial where "
		query += " 														curator = " + filial_id + " union select " + filial_id + " "
		query += " 												) "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	} else {

		query = ""
		query += " select * "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id = " + filial_id + " "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	return query

}

func (q *query_FilialRegulation) RefusalMinimizeDetails(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	query = ""

	if curator {

		query = ""
		query += " select * "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    (oa.trass_result = 'Частичный отказ' or oa.trass_result = 'Отказ') and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id in ( "
		query += " 													select id from public.filial where "
		query += " 														curator = " + filial_id + " union select " + filial_id + " "
		query += " 												) "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	} else {

		query = ""
		query += " select * "
		query += " 	from oisuu_auto as oa where "
		query += " 		oa.employer_trass != '' and "
		query += " 		oa.trass_result_date between '" + date_begin + "' and '" + date_end + "'  and "
		query += "	    (oa.trass_result = 'Частичный отказ' or oa.trass_result = 'Отказ') and "
		query += " 		oa.zone in "
		query += " 		( "
		query += " 			select z.name from ( "
		query += " 								select * "
		query += " 									from public.filial as f where "
		query += " 										f.id = " + filial_id + " "
		query += " 									) as fils "
		query += " 								left join oisuu_zones as z on "
		query += " 									(fils.oisuu_zone = z.id) "
		query += " 		) and "
		query += " 		oa.ins_type in ( "
		query += " 						select ot.name from journal_oisuu_types as oft "
		query += " 							left join ins_types as it on (oft.journal_type = it.id) "
		query += " 							left join oisuu_ins_types as ot on (oft.oisuu_type = ot.id) "
		query += " 								where it.name = '" + ins_type + "' "
		query += " 						) "
		query += " ;"

	}

	return query
}

type query_Regulation struct {
}

/*

func (q *query_Regulation) AutoPostupilo(
	ins_type string,
	date_begin string,
	date_end string,
	employer_id string) string {

	var query string

	query = ""

	query += "select count(*) from oisuu_auto as oa "
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



func (q *query_Regulation) AutoMinimize(
	ins_type string,
	date_begin string,
	date_end string,
	employer_id string) string {

	var query string

	query = ""

	query += "select count(*) from oisuu_auto as oa "
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

func (q *query_Regulation) AutoRefusalFinance(
	ins_type string,
	date_begin string,
	date_end string,
	employer_id string) string {

	var query string

	query = ""

	query += "select sum(finanse) from oisuu_auto as oa "
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

func (q *query_Regulation) AutoMinimizeFinance(
	ins_type string,
	date_begin string,
	date_end string,
	employer_id string) string {

	var query string

	query = ""

	query += "select sum(finanse) from oisuu_auto as oa "
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

*/

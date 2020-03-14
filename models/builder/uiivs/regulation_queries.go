package uiivs

//"fmt"

/*

id  name            other
---------------------------
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

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == false {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.created BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " ;"

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == false {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.created BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.filial = " + filial_id + " "
		query += " ;"

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == true {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.created BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " ;"

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == true {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.created BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " ;"

	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type != "Иные" {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.created BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' "
		query += " ;"

	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type == "Иные" {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.created BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) "
		query += " ;"
	}

	//fmt.Println(query)
	return query

}

func (q *query_FilialRegulation) Refusal(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == false {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " ;"

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == false {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " ;"
	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == true {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " ;"

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == true {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " ;"
	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type != "Иные" {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' "
		query += " ;"
	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type == "Иные" {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' "
		query += " ;"
	}

	return query
}

func (q *query_FilialRegulation) Minimize(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == false {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " ;"

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == false {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " ;"
	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == true {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " ;"

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == true {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " ;"
	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type != "Иные" {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' "
		query += " ;"
	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type == "Иные" {

		query = ""
		query += "  SELECT count(*) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' "
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

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == false {

		query = ""
		query += "  SELECT sum(money) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " ;"

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == false {

		query = ""
		query += "  SELECT sum(money) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " ;"
	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == true {

		query = ""
		query += "  SELECT sum(money) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " ;"

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == true {

		query = ""
		query += "  SELECT sum(money) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " ;"
	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type != "Иные" {

		query = ""
		query += "  SELECT sum(money) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' "
		query += " ;"
	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type == "Иные" {

		query = ""
		query += "  SELECT sum(money) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' "
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

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == false {

		query = ""
		query += "  SELECT sum(money) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " ;"

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == false {

		query = ""
		query += "  SELECT sum(money) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " ;"
	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == true {

		query = ""
		query += "  SELECT sum(money) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " ;"

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == true {

		query = ""
		query += "  SELECT sum(money) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " ;"
	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type != "Иные" {

		query = ""
		query += "  SELECT sum(money) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' "
		query += " ;"
	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type == "Иные" {

		query = ""
		query += "  SELECT sum(money) "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' "
		query += " ;"
	}

	return query

}

// ----------------------------------------------------------------------
// возвращает идентификаторы количества
// ----------------------------------------------------------------------

func (q *query_FilialRegulation) postupiloID(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == false {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.created BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " "

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == false {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.created BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.filial = " + filial_id + " "
		query += " "

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == true {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.created BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " "

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == true {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.created BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " "

	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type != "Иные" {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.created BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' "
		query += " "

	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type == "Иные" {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.created BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) "
		query += " "
	}

	//fmt.Println(query)
	return query

}

func (q *query_FilialRegulation) refusalID(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == false {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " "

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == false {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " "
	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == true {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " "

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == true {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " "
	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type != "Иные" {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' "
		query += " "
	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type == "Иные" {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Отказ' AND  "
		query += " 		im.control = 'Снят с контроля' "
		query += " "
	}

	return query
}

func (q *query_FilialRegulation) minimizeID(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == false {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " "

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == false {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial = " + filial_id + " "
		query += " "
	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type != "Иные" && curator == true {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " "

	}

	// Запрос по конктретному филиалу
	if filial_id != "-1" && ins_type == "Иные" && curator == true {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' AND  "
		query += " 		im.filial in "
		query += " 		    ( "
		query += " 		    	select id from public.filial where curator = " + filial_id + " union select " + filial_id + " "
		query += " 		    ) "
		query += " "
	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type != "Иные" {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name = '" + ins_type + "' AND  "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' "
		query += " "
	}

	// Запрос по всем филиалам
	if filial_id == "-1" && ins_type == "Иные" {

		query = ""
		query += "  SELECT im.id "
		query += " 	FROM    deem.reg_imushestvo AS im "
		query += " 		 LEFT JOIN "
		query += " 		    public.uiivs_ins_types AS uit "
		query += " 		 ON im.instype = uit.journal_instype_id "
		query += " 	WHERE "
		query += " 		im.closed BETWEEN '" + date_begin + "' AND '" + date_end + "'  AND "
		query += " 		uit.name IN ( "
		query += " 			SELECT name FROM public.uiivs_ins_types "
		query += " 			WHERE other = true "
		query += " 		) AND "
		query += " 		im.result = 'Минимизация' AND  "
		query += " 		im.control = 'Снят с контроля' "
		query += " "
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

	query = ""

	query += " SELECT ri.* from details.reg_imushestvo as ri "
	query += " where ri.id in ( "
	query += q.postupiloID(ins_type, date_begin, date_end, filial_id, curator)
	query += " ); "

	//fmt.Println(query)
	return query

}

func (q *query_FilialRegulation) RefusalDetails(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	query = ""

	query += " SELECT ri.* from details.reg_imushestvo as ri "
	query += " where ri.id in ( "
	query += q.refusalID(ins_type, date_begin, date_end, filial_id, curator)
	query += " ); "

	//fmt.Println(query)
	return query
}

func (q *query_FilialRegulation) MinimizeDetails(
	ins_type string,
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	query = ""

	query += " SELECT ri.* from details.reg_imushestvo as ri "
	query += " where ri.id in ( "
	query += q.minimizeID(ins_type, date_begin, date_end, filial_id, curator)
	query += " ); "

	//fmt.Println(query)
	return query

}

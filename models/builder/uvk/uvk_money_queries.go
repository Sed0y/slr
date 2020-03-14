package uvk

//"fmt"

type queries struct {
	FilialMoney query_FilialMoney
	//EmployerRegulation query_EmployerRegulation
}

type query_FilialMoney struct {
}

func (q *query_FilialMoney) Debitorka(
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	if filial_id == "-1" {

		query = ""
		query += " SELECT sum(gr.money) "
		query += " 	FROM deem.grafik as gr "
		query += " 		left join details.debtor as d on gr.code = d.id "
		query += " 		left join filial as f on (d.filial = f.id) "
		query += " 		left join users as u on (d.user = u.id) "
		query += " where "
		query += " 		gr.code_type = 'д' "
		query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
		query += " ; "

	} else {

		if curator == false {

			query = ""
			query += " SELECT sum(gr.money) "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.debtor as d on gr.code = d.id "
			query += " 		left join filial as f on (d.filial = f.id) "
			query += " 		left join users as u on (d.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'д' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and d.filial = " + filial_id + " "
			query += " ; "

		} else {

			query = ""
			query += " SELECT sum(gr.money) "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.debtor as d on gr.code = d.id "
			query += " 		left join filial as f on (d.filial = f.id) "
			query += " 		left join users as u on (d.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'д' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and d.filial  in ( "
			query += " 							select id from public.filial where "
			query += " 							curator = " + filial_id + " "
			query += " 							union select " + filial_id + "  "
			query += " 		) "
			query += " ; "

		}

	}

	return query
}

func (q *query_FilialMoney) VnutrenneeMoshennichestvo(
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	if filial_id == "-1" {

		query = ""
		query += " SELECT sum(gr.money) "
		query += " 	FROM deem.grafik as gr "
		query += " 		left join details.vm as vm on gr.code = vm.id "
		query += " 		left join filial as f on (vm.filial = f.id) "
		query += " 		left join users as u on (vm.user = u.id) "
		query += " where "
		query += " 		gr.code_type = 'вк' "
		query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
		query += " ; "

	} else {

		if curator == false {

			query = ""
			query += " SELECT sum(gr.money) "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.vm as vm on gr.code = vm.id "
			query += " 		left join filial as f on (vm.filial = f.id) "
			query += " 		left join users as u on (vm.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'вк' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and vm.filial = " + filial_id + " "
			query += " ; "

		} else {

			query = ""
			query += " SELECT sum(gr.money) "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.vm as vm on gr.code = vm.id "
			query += " 		left join filial as f on (vm.filial = f.id) "
			query += " 		left join users as u on (vm.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'вк' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and vm.filial in ( "
			query += " 							select id from public.filial where "
			query += " 							curator = " + filial_id + " "
			query += " 							union select " + filial_id + "  "
			query += " 		) "
			query += " ; "

		}

	}

	return query
}

func (q *query_FilialMoney) UgolovnieDela(
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	if filial_id == "-1" {

		query = ""
		query += " SELECT sum(gr.money) "
		query += " 	FROM deem.grafik as gr "
		query += " 		left join details.ud as ud on gr.code = ud.id "
		query += " 		left join filial as f on (ud.filial = f.id) "
		query += " 		left join users as u on (ud.user = u.id) "
		query += " where "
		query += " 		gr.code_type = 'з' "
		query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
		query += " ; "

	} else {

		if curator == false {

			query = ""
			query += " SELECT sum(gr.money) "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.ud as ud on gr.code = ud.id "
			query += " 		left join filial as f on (ud.filial = f.id) "
			query += " 		left join users as u on (ud.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'з' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and ud.filial = " + filial_id + " "
			query += " ; "

		} else {

			query = ""
			query += " SELECT sum(gr.money) "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.ud as ud on gr.code = ud.id "
			query += " 		left join filial as f on (ud.filial = f.id) "
			query += " 		left join users as u on (ud.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'з' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and ud.filial in ( "
			query += " 							select id from public.filial where "
			query += " 							curator = " + filial_id + " "
			query += " 							union select " + filial_id + "  "
			query += " 		) "
			query += " ; "

		}

	}

	return query
}

func (q *query_FilialMoney) Subrogaciya(
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	if filial_id == "-1" {

		query = ""
		query += " SELECT sum(gr.money) "
		query += " 	FROM deem.grafik as gr "
		query += " 		left join details.subrogaciya as s on gr.code = s.id "
		query += " 		left join filial as f on (s.filial = f.id) "
		query += " 		left join users as u on (s.user = u.id) "
		query += " where "
		query += " 		gr.code_type = 'с' "
		query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
		query += " ; "

	} else {

		if curator == false {

			query = ""
			query += " SELECT sum(gr.money) "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.subrogaciya as s on gr.code = s.id "
			query += " 		left join filial as f on (s.filial = f.id) "
			query += " 		left join users as u on (s.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'с' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and s.filial = " + filial_id + " "
			query += " ; "

		} else {

			query = ""
			query += " SELECT sum(gr.money) "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.subrogaciya as s on gr.code = s.id "
			query += " 		left join filial as f on (s.filial = f.id) "
			query += " 		left join users as u on (s.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'с' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and s.filial in ( "
			query += " 							select id from public.filial where "
			query += " 							curator = " + filial_id + " "
			query += " 							union select " + filial_id + "  "
			query += " 		) "
			query += " ; "

		}

	}

	return query
}

func (q *query_FilialMoney) DebitorkaDetails(
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	if filial_id == "-1" {

		query = ""
		query += " SELECT "
		query += " 		d.id as code, "
		query += " 		gr.date, "
		query += " 		f.name as filial, "
		query += " 		u.login as user, "
		query += " 		d.debtor, "
		query += " 		gr.money "
		query += " 	FROM deem.grafik as gr "
		query += " 		left join details.debtor as d on gr.code = d.id "
		query += " 		left join filial as f on (d.filial = f.id) "
		query += " 		left join users as u on (d.user = u.id) "
		query += " where "
		query += " 		gr.code_type = 'д' "
		query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
		query += " order by gr.date desc "
		query += " ; "

	} else {

		if curator == false {

			query = ""
			query += " SELECT "
			query += " 		d.id as code, "
			query += " 		gr.date, "
			query += " 		f.name as filial, "
			query += " 		u.login as user, "
			query += " 		d.debtor, "
			query += " 		gr.money "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.debtor as d on gr.code = d.id "
			query += " 		left join filial as f on (d.filial = f.id) "
			query += " 		left join users as u on (d.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'д' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and d.filial = " + filial_id + " "
			query += " order by gr.date desc "
			query += " ; "

		} else {

			query = ""
			query += " SELECT "
			query += " 		d.id as code, "
			query += " 		gr.date, "
			query += " 		f.name as filial, "
			query += " 		u.login as user, "
			query += " 		d.debtor, "
			query += " 		gr.money "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.debtor as d on gr.code = d.id "
			query += " 		left join filial as f on (d.filial = f.id) "
			query += " 		left join users as u on (d.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'д' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and d.filial  in ( "
			query += " 							select id from public.filial where "
			query += " 							curator = " + filial_id + " "
			query += " 							union select " + filial_id + "  "
			query += " 		) "
			query += " order by gr.date desc "
			query += " ; "

		}

	}

	return query
}

func (q *query_FilialMoney) VnutrenneeMoshennichestvoDetails(
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	if filial_id == "-1" {

		query = ""
		query += " SELECT "
		query += " 		vm.id as code, "
		query += " 		gr.date, "
		query += " 		f.name as filial, "
		query += " 		u.login as user, "
		query += " 		vm.debtor, "
		query += " 		gr.money "
		query += " 	FROM deem.grafik as gr "
		query += " 		left join details.vm as vm on gr.code = vm.id "
		query += " 		left join filial as f on (vm.filial = f.id) "
		query += " 		left join users as u on (vm.user = u.id) "
		query += " where "
		query += " 		gr.code_type = 'вк' "
		query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
		query += " ; "

	} else {

		if curator == false {

			query = ""
			query += " SELECT "
			query += " 		vm.id as code, "
			query += " 		gr.date, "
			query += " 		f.name as filial, "
			query += " 		u.login as user, "
			query += " 		vm.debtor, "
			query += " 		gr.money "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.vm as vm on gr.code = vm.id "
			query += " 		left join filial as f on (vm.filial = f.id) "
			query += " 		left join users as u on (vm.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'вк' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and vm.filial = " + filial_id + " "
			query += " ; "

		} else {

			query = ""
			query += " SELECT "
			query += " 		vm.id as code, "
			query += " 		gr.date, "
			query += " 		f.name as filial, "
			query += " 		u.login as user, "
			query += " 		vm.debtor, "
			query += " 		gr.money "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.vm as vm on gr.code = vm.id "
			query += " 		left join filial as f on (vm.filial = f.id) "
			query += " 		left join users as u on (vm.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'вк' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and vm.filial in ( "
			query += " 							select id from public.filial where "
			query += " 							curator = " + filial_id + " "
			query += " 							union select " + filial_id + "  "
			query += " 		) "
			query += " ; "

		}

	}

	return query
}

func (q *query_FilialMoney) UgolovnieDelaDetails(
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	if filial_id == "-1" {

		query = ""
		query += " SELECT "
		query += " 		ud.id as code, "
		query += " 		gr.date, "
		query += " 		f.name as filial, "
		query += " 		u.login as user, "
		query += " 		ud.debtor, "
		query += " 		it.name, "
		query += " 		gr.money "
		query += " 	FROM deem.grafik as gr "
		query += " 		left join details.ud as ud on gr.code = ud.id "
		query += " 		left join filial as f on (ud.filial = f.id) "
		query += " 		left join users as u on (ud.user = u.id) "
		query += " 		left join ud_ins_types as it on (ud.ins_type = it.id) "
		query += " where "
		query += " 		gr.code_type = 'з' "
		query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
		query += " ; "

	} else {

		if curator == false {

			query = ""
			query += " SELECT "
			query += " 		ud.id as code, "
			query += " 		gr.date, "
			query += " 		f.name as filial, "
			query += " 		u.login as user, "
			query += " 		ud.debtor, "
			query += " 		it.name, "
			query += " 		gr.money "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.ud as ud on gr.code = ud.id "
			query += " 		left join filial as f on (ud.filial = f.id) "
			query += " 		left join users as u on (ud.user = u.id) "
			query += " 		left join ud_ins_types as it on (ud.ins_type = it.id) "
			query += " where "
			query += " 		gr.code_type = 'з' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and ud.filial = " + filial_id + " "
			query += " ; "

		} else {

			query = ""
			query += " SELECT "
			query += " 		ud.id as code, "
			query += " 		gr.date, "
			query += " 		f.name as filial, "
			query += " 		u.login as user, "
			query += " 		ud.debtor, "
			query += " 		it.name, "
			query += " 		gr.money "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.ud as ud on gr.code = ud.id "
			query += " 		left join filial as f on (ud.filial = f.id) "
			query += " 		left join users as u on (ud.user = u.id) "
			query += " 		left join ud_ins_types as it on (ud.ins_type = it.id) "
			query += " where "
			query += " 		gr.code_type = 'з' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and ud.filial in ( "
			query += " 							select id from public.filial where "
			query += " 							curator = " + filial_id + " "
			query += " 							union select " + filial_id + "  "
			query += " 		) "
			query += " ; "

		}

	}

	return query
}

func (q *query_FilialMoney) SubrogaciyaDetails(
	date_begin string,
	date_end string,
	filial_id string,
	curator bool) string {

	var query string

	if filial_id == "-1" {

		query = ""
		query += " SELECT "
		query += " 		s.id as code, "
		query += " 		gr.date, "
		query += " 		f.name as filial, "
		query += " 		u.login as user, "
		query += " 		s.debtor, "
		query += " 		gr.money "
		query += " 	FROM deem.grafik as gr "
		query += " 		left join details.subrogaciya as s on gr.code = s.id "
		query += " 		left join filial as f on (s.filial = f.id) "
		query += " 		left join users as u on (s.user = u.id) "
		query += " where "
		query += " 		gr.code_type = 'с' "
		query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
		query += " ; "

	} else {

		if curator == false {

			query = ""
			query += " SELECT "
			query += " 		s.id as code, "
			query += " 		gr.date, "
			query += " 		f.name as filial, "
			query += " 		u.login as user, "
			query += " 		s.debtor, "
			query += " 		gr.money "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.subrogaciya as s on gr.code = s.id "
			query += " 		left join filial as f on (s.filial = f.id) "
			query += " 		left join users as u on (s.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'с' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and s.filial = " + filial_id + " "
			query += " ; "

		} else {

			query = ""
			query += " SELECT "
			query += " 		s.id as code, "
			query += " 		gr.date, "
			query += " 		f.name as filial, "
			query += " 		u.login as user, "
			query += " 		s.debtor, "
			query += " 		gr.money "
			query += " 	FROM deem.grafik as gr "
			query += " 		left join details.subrogaciya as s on gr.code = s.id "
			query += " 		left join filial as f on (s.filial = f.id) "
			query += " 		left join users as u on (s.user = u.id) "
			query += " where "
			query += " 		gr.code_type = 'с' "
			query += " 		and gr.date between '" + date_begin + "' and '" + date_end + "' "
			query += " 		and s.filial in ( "
			query += " 							select id from public.filial where "
			query += " 							curator = " + filial_id + " "
			query += " 							union select " + filial_id + "  "
			query += " 		) "
			query += " ; "

		}

	}

	return query
}

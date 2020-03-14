


SELECT 
	sub2.*,
	f.name as filial
FROM
(
	SELECT
		sub.rid,
		sub.uid,
		us.login,
		sub.url,
		sub.params,
		sub.dt,
		CASE WHEN sub.r~E'^\\d+$' THEN CAST (sub.r AS INTEGER) ELSE 0 END as sfil	
	FROM
	(
		SELECT 	ur.*, 
			t.* FROM jlogs.urls as ur, unnest(string_to_array(ur.params, '#')) WITH ORDINALITY AS t(r,n)
		WHERE t.n = 1
	) as sub
	LEFT JOIN public.users as us on sub.uid = us.id
) as sub2
LEFT JOIN public.filial as f on sub2.sfil = f.id

WHERE 
	--sub2.dt > timestamp '2018-08-20 10:41:00'
	--sub2.dt > timestamp '2018-08-30 10:41:00'
	--sub2.dt > timestamp '2018-08-30 11:01:00'
	--sub2.dt > timestamp '2018-08-30 13:52:00'
	sub2.dt > timestamp '2018-10-04 00:00:00'
	and sub2.uid <> 230
	and sub2.uid <> 64
	--and sub2.url = '/report/filial/excel'
order by sub2.rid desc

/*
	0aj8q6X9
	17148742685
	http://rzbpm.ru/knowledge/bpmn-2-0-iz-chego-sostoit-model-biznes-processa.html
	
*/



select
   CONNECTIONPROPERTY('net_transport') AS net_transport,
   CONNECTIONPROPERTY('protocol_type') AS protocol_type,
   CONNECTIONPROPERTY('auth_scheme') AS auth_scheme,
   CONNECTIONPROPERTY('local_net_address') AS local_net_address,
   CONNECTIONPROPERTY('local_tcp_port') AS local_tcp_port,
   CONNECTIONPROPERTY('client_net_address') AS client_net_address 

   

=======================================================================================================




SELECT onum as №, cat as Категория, period as "За период", year as "За год" FROM
(
	
		SELECT '3' as onum, q1.cat, q1.period, q2.year FROM 
			(
				SELECT 'Внутреннее мошенничество' as cat, sum(money) as period
				FROM deem.grafik
				where 
					date between $__timeFrom() and $__timeTo()
					and code_type = 'вк'
			) as q1
		LEFT JOIN 
			(
				SELECT 'Внутреннее мошенничество' as cat, sum(money) as year
				FROM deem.grafik as g2
				where 
					date between (Extract(YEAR from now()) || '-01-01')::date and $__timeTo()
					and code_type = 'вк'											
			) as q2
		ON (q1.cat = q2.cat)

	UNION
		
		SELECT '2' as onum, q1.cat, q1.period, q2.year FROM 
		(
			SELECT 'Уголовные дела' as cat, sum(g.money) as period
			FROM 
				deem.grafik as g
			LEFT JOIN
				details.ud as d
			ON g.code = d.id
				where 
					g.code_type='з' and 
					d.ins_type = 11 and
					g.date between $__timeFrom() and $__timeTo()
		) as q1
		LEFT JOIN 
		(
			SELECT 'Уголовные дела' as cat, sum(g.money) as year
			FROM 
				deem.grafik as g
			LEFT JOIN
				details.ud as d
			ON g.code = d.id
				where 
					g.code_type='з' and 
					d.ins_type = 11 and
					g.date between (Extract(YEAR from now()) || '-01-01')::date and $__timeTo()										
		) as q2
		on (q1.cat = q2.cat)
		
	UNION
		
		SELECT '1' as onum, q1.cat, q1.period, q2.year FROM 
		(
		SELECT 'Дебиторская задолженность' as cat, sum(money) as period
		FROM deem.grafik
		where 
			date between $__timeFrom() and $__timeTo()
			and code_type = 'д'
		) as q1
		LEFT JOIN 
		(
			SELECT 'Дебиторская задолженность' as cat, sum(money) as year
			FROM deem.grafik as g2
			where 
				date between (Extract(YEAR from now()) || '-01-01')::date and $__timeTo()
				and code_type = 'д'											
		) as q2
		on (q1.cat = q2.cat)

	UNION
			
		
	SELECT '4' as onum,'Итого' as cat, sum(period) as period, sum(year) as year
	FROM
	(
		SELECT q1.cat, q1.period, q2.year FROM 
		(
		SELECT 'ВМ' as cat, sum(money) as period
		FROM deem.grafik
		where 
			date between $__timeFrom() and $__timeTo()
			and code_type = 'вк'
		) as q1
		LEFT JOIN 
		(
			SELECT 'ВМ' as cat, sum(money) as year
			FROM deem.grafik as g2
			where 
				date between (Extract(YEAR from now()) || '-01-01')::date and $__timeTo()
				and code_type = 'вк'											
		) as q2
		on (q1.cat = q2.cat)

	UNION
		
		SELECT q1.cat, q1.period, q2.year FROM 
		(
			SELECT 'УД' as cat, sum(g.money) as period
			FROM 
				deem.grafik as g
			LEFT JOIN
				details.ud as d
			ON g.code = d.id
				where 
					g.code_type='з' and 
					d.ins_type = 11 and
					g.date between $__timeFrom() and $__timeTo()
		) as q1
		LEFT JOIN 
		(
			SELECT 'УД' as cat, sum(g.money) as year
			FROM 
				deem.grafik as g
			LEFT JOIN
				details.ud as d
			ON g.code = d.id
				where 
					g.code_type='з' and 
					d.ins_type = 11 and
					g.date between (Extract(YEAR from now()) || '-01-01')::date and $__timeTo()										
		) as q2
		on (q1.cat = q2.cat)
		
	UNION
		
		SELECT q1.cat, q1.period, q2.year FROM 
		(
		SELECT 'ПДЗ' as cat, sum(money) as period
		FROM deem.grafik
		where 
			date between $__timeFrom() and $__timeTo()
			and code_type = 'д'
		) as q1
		LEFT JOIN 
		(
			SELECT 'ПДЗ' as cat, sum(money) as year
			FROM deem.grafik as g2
			where 
				date between (Extract(YEAR from now()) || '-01-01')::date and $__timeTo()
				and code_type = 'д'											
		) as q2
		on (q1.cat = q2.cat)
	) as subq
) as res
order by res.onum asc


























SELECT q1.cat, q1.period, q2.year FROM
	(
		SELECT 'Направлено заявлений' as cat, count(*) as period
		FROM details.ud
		WHERE
			ins_type = 11 AND							
			zayavlenie_date BETWEEN $__timeFrom() and $__timeTo() 
	) as q1
	LEFT JOIN
	(
		SELECT 'Направлено заявлений' as cat, count(*) as year
		FROM details.ud
		WHERE
			ins_type = 11 AND							
			zayavlenie_date BETWEEN (Extract(YEAR from now()) || '-01-01')::date and $__timeTo()
	) as q2
	ON q1.cat = q2.cat
	
	==================================================================================================
	
	
	
SELECT * FROM 
(
	SELECT 1 as order_number, q1.cat, q1.period, q2.year FROM
	(
		SELECT 'Направлено заявлений' as cat, count(*) as period
		FROM details.ud
		WHERE
			ins_type = 11 AND							
			zayavlenie_date BETWEEN $__timeFrom() and $__timeTo() 
	) as q1
	LEFT JOIN
	(
		SELECT 'Направлено заявлений' as cat, count(*) as year
		FROM details.ud
		WHERE
			ins_type = 11 AND							
			zayavlenie_date BETWEEN (Extract(YEAR from now()) || '-01-01')::date and $__timeTo()
	) as q2
	ON q1.cat = q2.cat
	
UNION
	
	SELECT 2 as order_number, q1.cat, q1.period, q2.year FROM
	(
		SELECT 'Возбуждено уголовных дел' as cat, count(*) as period
		FROM details.ud
		WHERE
			ins_type = 11 AND							
			vozbujdenie_date BETWEEN $__timeFrom() and $__timeTo() 
	) as q1
	LEFT JOIN
	(
		SELECT 'Возбуждено уголовных дел' as cat, count(*) as year
		FROM details.ud
		WHERE
			ins_type = 11 AND							
			vozbujdenie_date BETWEEN (Extract(YEAR from now()) || '-01-01')::date and $__timeTo()
	) as q2
	ON q1.cat = q2.cat

UNION
	
	SELECT 3 as order_number, q1.cat, q1.period, q2.year FROM
	(
		SELECT 'Вынесено приговоров' as cat, count(*) as period
		FROM details.ud
		WHERE
			ins_type = 11 AND
			reshenie = 'Приговор' AND
			reshenie_date BETWEEN $__timeFrom() and $__timeTo()
	) as q1
	LEFT JOIN
	(
		SELECT 'Вынесено приговоров' as cat, count(*) as year
		FROM details.ud
		WHERE
			ins_type = 11 AND
			reshenie = 'Приговор' AND
			reshenie_date BETWEEN (Extract(YEAR from now()) || '-01-01')::date and $__timeTo()
	) as q2
	ON q1.cat = q2.cat
	
UNION	
	
	SELECT 4 as order_number, q1.cat, q1.period, q2.year FROM
	(
		SELECT 'Проведено служебных проверок' as cat, count(*) as period
		FROM details.vm
		WHERE
			closed BETWEEN $__timeFrom() and $__timeTo()
	) as q1
	LEFT JOIN
	(
		SELECT 'Проведено служебных проверок' as cat, count(*) as year
		FROM details.vm
		WHERE
			closed BETWEEN (Extract(YEAR from now()) || '-01-01')::date and $__timeTo()
	) as q2
	ON q1.cat = q2.cat	
	
	
	
) as res_query
order by order_number asc







SELECT * FROM 
(

	SELECT 1 as order_number, 'Служебных проверок' as cat, count(*) as count
		FROM details.vm
		WHERE 
			closed is null
			and control <> 'Снят с контроля'

	UNION
			
	SELECT 2 as order_number, 'Количество материалов ПДЗ' as cat, count(*) as count
		FROM details.debtor
		WHERE 
			closed is null
			and control <> 'Снят с контроля'

	UNION

	SELECT 3 as order_number, 'Материалов' as cat, count(*) as count
		FROM details.ud
		WHERE
			ins_type = 11 and
			material_date is not null and 
			material_reshenie_date is null

	UNION

	SELECT 4 as order_number, 'Уголовных дел' as cat, count(*) as count
		FROM details.ud
		WHERE
			ins_type = 11 and
			vozbujdenie_date is not null and 
			ud_reshsenie_date is null
) as resq
order by order_number asc
		
		
		
Зеньков Вадим Николаевич <zenkov@VSK.RU>

=======================================================================================================




Возвраты по ПДЗ, УД и ВМ
Преддоговорная проверка


http://hd.vsk.ru/WorkOrder.do?reqTemplate=49302&requestServiceId=2701


postgres=# SELECT to_char(1210.7, 'L9G999.99');
  to_char
------------
 $ 1,210.70
















    <div class="container-fluid">
		<!-- Breadcrumbs-->
		<ol class="breadcrumb">
			<li class="breadcrumb-item"><a href="#">Система</a></li>
			<li class="breadcrumb-item active">Базы данных</li>
		</ol>
  
		<div class="row">
			<div class="col-lg-8">
			  <div class="card mb-3">
				<div class="card-header">
				  <i class="fa fa-database"></i> Синхронизация баз данных c Access</div>
				<div class="card-body">
				​
				<table id="AccessTables" style="width:100%; font-size: 0.75rem;" class="table table-bordered compact" cellspacing="0">
					<!-- AccessTables -->
					<thead>
			            <tr>			   	            
                			<th>Название</th>
               				<th>Описание</th>
                			<th>Название Access</th>
                			<th>Путь</th>
                			<th>Обновлена</th>
							<th>Действие</th>
						</tr>				
        			</thead>	
		        	<tbody>
					{{if .AccessTables -}}						
						{{ range $key, $value := .AccessTables }}										
						<tr> 							
							<td>{{$value.Name}}</td> 
							<td>{{$value.Description}}</td> 
							<td>{{$value.DB_name}}</td> 
							<td>{{$value.DB_path}}</td> 
							<td>{{$value.Last_update}}</td> 
							<td><span class="Update" id="Update{{$value.Id}}" table_id="{{$value.Id}}">Обновить</span></td> 
						</tr>
						{{end}}						
					{{- else}}
							Не удалось загрузить список сотрудников
					{{- end}}
					</tbody>
				</table>
				  
				</div>
				<div class="card-footer small text-muted">СУБД: <i>PostgreSQL</i>, домен: <i>localhost</i>, БД: <i>journal</i></div>
			  </div>
			</div>
		
			<div class="col-lg-4">
			  <div class="card mb-4">
				<div class="card-header">
				  <i class="fa fa-database"></i> Синхронизация баз данных c ОИСУУ</div>
				<div class="card-body">
					Синхронизация осуществляется через файл иксель: </br>
					\\dpm\dpm\ЖР\Общие\АвтоОИСУУ\Отчет3.xlsm</br></br>
				 	<span id="ReloadOISUU"><strong>Обновить</strong></span>
				</div>
				<div class="card-footer small text-muted">СУБД: <i>PostgreSQL</i>, домен: <i>localhost</i>, БД: <i>journal, префикс oisuu_</i></div>
			  </div>
			</div>
		</div>	
	</div>


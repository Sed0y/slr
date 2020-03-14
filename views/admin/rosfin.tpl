

    <div class="container-fluid">
		<!-- Breadcrumbs-->
		<ol class="breadcrumb">
			<li class="breadcrumb-item"><a href="#">Система</a></li>
			<li class="breadcrumb-item active">Базы данных</li>
		</ol>
  
		<div class="row">
			<div class="col-lg-9">
			  <div class="card mb-3">
				<div class="card-header">
					<i class="fa fa-database"></i> Обновление данных росфинмониторинга в стоп-листе						
				</div>
				<div class="card-body">
					<div id="RosfinUploadFile">
						<form 	enctype="multipart/form-data" 
								action="/upload/rosfin" 
								method="post" 
								onsubmit="return validate_upload_form();" 
								name="upload_form"	
						>
							<span>Актуальность:</span><input type="text" name="actual" size="10" id="rosfin-actual-data"/>;
							<span>Описание:</span><input type="text" name="description" size="20" id="rosfin-description"/>;
							<span style="float:right;">
								<input type="file" name="uploadfile" />
							    <input type="hidden" name="token" value="{{.}}"/>
							    <input type="submit" value="Загрузить"  />
							</span>
						</form>		
					</div>

					<div id="rosfin_progressbar">
						<div class="progress-label">Обработка...</div>	
					</div>
					
					<div id="rosfin_progress_console">						
						<div class="rosfin_progress_item" id="download-link" style="height:0px; overflow:hidden;">
						</div>
						<div id="progress_count"></div>
					</div>


					<table id="RosfinFiles" style="width:100%; font-size: 0.75rem;" class="table table-bordered compact" cellspacing="0">						
						<thead>
							<tr>			   	            
								<th>Файл</th>
								<th>Описание</th>
								<th>Добавлен</th>
								<th>Актуальность</th>
								<th>Загружен в СЛ</th>
								<th>В архиве</th>
								<th>Стоп-лист</th>																				
							</tr>				
						</thead>	
						<tbody>
	
	
						{{if .RosfinFiles -}}						
							{{ range $key, $value := .RosfinFiles }}										
							<tr fileid="{{$value.Id}}"
								{{if $value.Exist -}}										
									{{- else}}										
										class="RosfinNotExist"
									{{- end}}
							> 							
								<td>{{$value.FileName}}</td> 
								<td>{{$value.Description}}</td> 
								<td>{{$value.Loaded}}</td> 
								<td>{{$value.Actual}}</td> 								
								<td id = "LoadedDate{{$value.Id}}">{{$value.LoadedToStop}}</td> 
								<td style="text-align:center;">
									{{if $value.Exist -}}
										Да
									{{- else}}
										Нет 
									{{- end}}
										/ <span class="Delete" id="Rosfin-Delete-{{$value.Id}}" table_id="{{$value.Id}}">Удалить</span> 
									
								</td> 
								<td style="text-align:center;" id="rosfin_fileid_{{$value.Id}}">
									{{if $value.Exist -}}
										<span class="Update" id="Rosfin-Update-{{$value.Id}}" table_id="{{$value.Id}}">Загрузить в СЛ</span>
									{{- else}}
										
									{{- end}}
								</td> 
							</tr>
							{{end}}												
								
						{{- end}}
						</tbody>
					</table>
				​
				
				</div>
				<div class="card-footer small text-muted">СУБД: <i>MsSql</i>, домен: <i>autosqlcl</i>, БД: <i>stop_list</i></div>
			  </div>
			</div>
		
			<div class="col-lg-3">
			  <div class="card mb-4">
				<div class="card-header">
				  <i class="fa fa-database"></i> Актуальные данные</div>
				<div class="card-body">					
					{{if ne .InStopList.FileName "" -}}																			
						<strong>{{.InStopList.FileName}}</strong><br>
						{{.InStopList.Description}}<br>
						Загружен в архив: {{.InStopList.Loaded}}<br>
						Актуальность данных: {{.InStopList.Actual}}<br>
						Дата загрузки в СЛ: {{.InStopList.LoadedToStop}}<br>
					{{- else}}										
						Нет загруженных файлов
					{{- end}}
				</div>
				<div class="card-footer small text-muted"></div>
			  </div>
			</div>
		</div>	
	</div>


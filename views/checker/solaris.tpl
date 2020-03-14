

 
  
    <div class="container-fluid">
		<!-- Breadcrumbs
		<ol class="breadcrumb">
			<li class="breadcrumb-item"><a href="#">Проверка</a></li>
			<li class="breadcrumb-item active">Стоп-лист</li>
		</ol>  
		-->
		<div class="row">
    	
			<div class="col-lg-8">
				<div class="card">
					<div class="card-header">
						<i class="fa fa-user-secret"></i> Результаты поиска
					</div>


					<div class="card-body">
					
						<!--
						<div id="ItemDetails">
						
							<div class="card" style="background-color:#f5f9fb;">
							  <div class="card-body">
								<strong><span id="res_fio">-</span></br><span id="res_dr">-</span></strong>								
							  </div>							  
							</div>
							
							<p class="card-part-header"><em><i class="fa fa-exclamation-triangle" style="color:#a20b0b;"></i>Причина учёта</em></p>
							<p class="card-part-content">
								<label><strong>Добавлен:</strong></label> <span id="res_add_date">-</span>
								</br>
								<span id="res_reason">
									-
								</span>
							</p>
												
							
						</div>
						-->
						<table id="WaitTable" >
							<tbody>
								<tr>
									<td>
										<i class="fa fa-spinner fa-pulse fa-3x fa-fw"></i></br></br>Подождите, идёт поиск...
									</td>
								</tr>								
							</tbody>
						</table>
						<div id="SolarisResultTable"></div>
					</div>					
					
					
					
					<div class="card-footer small text-muted"></div>
					
					
					
				</div>
			</div>
		
		
			<div class="col-lg-4">
				<div class="card" style="min-width:420px;">
					<div class="card-header">
						<i class="fa"></i> Поиск по БД Солярис
					</div>
					<div class="card-body">
			<!--
						<div class="alert alert-light" role="alert" id="MainInfo" style="display:none;">
							<span>&nbsp; </span>
							<br>
							<i>&nbsp;</i>
						</div>
			-->
		
			<!--
						<ul class="nav nav-tabs" id="myTab" role="tablist">
						  <li class="nav-item">
							<a class="nav-link active" href="#solaris_person" data-toggle="tab">Физические лица</a>
						  </li>						  
						</ul>
			-->	
	
						<div class="tab-content">				
							
							<div class="tab-pane active" id="solaris_person">
									
								
								<p>
									 <em style="color:#484848;">&nbsp;</em>
									<em style="color:#484848;">
										<span id="CleanRequestParameters" onClick="solaris_HideAllRequestParameter('req_person_name');">Очистить параметры запроса</span>
									</em>
								</p>
								<div id="RequestParameters" class="alert alert-warning" role="alert">
									<span id="PanelTitle"><em style="color:#484848;">Состав запроса</em></span>
									<table>
										<tbody>
											<tr id="req_person_surname">
												<td class="req_label" width="25%"><em>Фамилия:</em></td>
												<td class="req_value" width="55%">-</td>
												<td class="req_clear" width="20%" onClick="solaris_HideRequestParameter('req_person_surname');">x</td>
											</tr>
											<tr id="req_person_name">
												<td class="req_label" width="25%"><em>Имя:</em></td>
												<td class="req_value" width="55%">-</td>
												<td class="req_clear" width="20%" onClick="solaris_HideRequestParameter('req_person_name');">x</td>
											</tr>
											<tr id="req_person_fathername">
												<td class="req_label" width="25%"><em>Отчество:</em></td>
												<td class="req_value" width="55%">-</td>
												<td class="req_clear" width="20%" onClick="solaris_HideRequestParameter('req_person_fathername');">x</td>
											</tr>
											<tr id="req_person_birthdate">
												<td class="req_label" width="25%"><em>День рождения:</em></td>
												<td class="req_value" width="55%">-</td>
												<td class="req_clear" width="20%" onClick="solaris_HideRequestParameter('req_person_birthdate');">x</td>
											</tr>										
											</tr>
											<tr id="req_person_inn">
												<td class="req_label" width="25%"><em>ИНН ФЛ:</em></td>
												<td class="req_value" width="55%">-</td>
												<td class="req_clear" width="20%" onClick="solaris_HideRequestParameter('req_person_inn');">x</td>
											</tr>											
										</tbody>
									</table>									
									
									
									
									
								</div>
								
								<div class="alert alert-danger" role="alert" id = "solaris_form_person_valid_message" style="display: none;"></div>	

								<div class="row">
									<div class="col-sm-11" style="text-align: right; padding-right: 27px; padding-top: 7px;" >  					
										<em>Добавить всё</em>
									</div>
									<div class="alert alert-danger col-sm-1 add_value_buttom" role="alert" id = "solaris_add_all">
										<span onClick="solaris_AddRequestParameter('all');">
											<i class="fa fa-plus"></i>
										</span>
									</div>
								</div>
								<hr></hr>
								<form >										
									<div class="form-group row">
										<label for="solaris_surname" class="col-sm-3 col-form-label" style="margin-left: 10px;" >Фамилия:</label>
										<div class="col-sm-8" style="margin-left: -10px;">  					
											<input type="text" class="form-control" id="solaris_surname"  onpaste="CheckForFullString();">
										</div>		
										
										<div class="alert alert-danger col-sm-1 add_value_buttom" role="alert" id = "solaris_add_surname">
											<span onClick="solaris_AddRequestParameter('req_person_surname');"><i class="fa fa-plus"></i></span>
										</div>
										
									</div>

									<div class="form-group row">
										<label for="solaris_name" class="col-sm-3 col-form-label" style="margin-left: 10px;">Имя:</label>
										<div class="col-sm-8" style="margin-left: -10px;">
											<input type="text" class="form-control" id="solaris_name" >
										</div>
										
										<div class="alert alert-danger col-sm-1 add_value_buttom" role="alert" id = "solaris_add_name">
											<span onClick="solaris_AddRequestParameter('req_person_name');"><i class="fa fa-plus"></i></span>
										</div>
										
									</div>

									<div class="form-group row">
										<label for="solaris_fathername" class="col-sm-3 col-form-label" style="margin-left: 10px;">Отчество:</label>
										<div class="col-sm-8" style="margin-left: -10px;">
											<input type="text" class="form-control" id="solaris_fathername" >
										</div>
										
										<div class="alert alert-danger col-sm-1 add_value_buttom" role="alert" id = "solaris_add_fathername">
											<span onClick="solaris_AddRequestParameter('req_person_fathername');"><i class="fa fa-plus"></i></span>
										</div>
										
									</div>
									  
									<div class="form-group row">
										<label class="col-sm-3 col-form-label" style="margin-left: 10px;">Дата:</label>
										
										<div class="col-sm-2" style="margin-left: -10px;">
											<input type="text" class="form-control" id="solaris_birthdate_day" >
										</div>
										<div class="alert alert-danger col-sm-1 add_value_buttom" role="alert" id = "solaris_add_inn">
											<span onClick="solaris_AddRequestParameter('req_person_birthdate');"><i class="fa fa-plus"></i></span>
										</div>
										
										<div class="col-sm-2" style="margin-left: -1px;">
											<input type="text" class="form-control" id="solaris_birthdate_month" >
										</div>
										<div class="alert alert-danger col-sm-1 add_value_buttom" role="alert" id = "solaris_add_inn">
											<span onClick="solaris_AddRequestParameter('req_person_birthdate');"><i class="fa fa-plus"></i></span>
										</div>
										
										<div class="col-sm-3 " style="margin-left: -1px;">
											<input type="text" class="form-control" id="solaris_birthdate_year"  >
										</div>
										<div class="alert alert-danger col-sm-1 add_value_buttom" role="alert" id = "solaris_add_inn">
											<span onClick="solaris_AddRequestParameter('req_person_birthdate');"><i class="fa fa-plus"></i></span>
										</div>
										
									</div>
																		
									<hr></hr>
									
									<div class="form-group row">
										<label for="solaris_inn" class="col-sm-3 col-form-label" style="margin-left: 10px;">ИНН:</label>
										<div class="col-sm-8" style="margin-left: -10px;">
											<input type="text" class="form-control" id="solaris_inn" >
										</div>										
										
										<div class="alert alert-danger col-sm-1 add_value_buttom" role="alert" id = "solaris_add_inn">
											<span onClick="solaris_AddRequestParameter('req_person_inn');"><i class="fa fa-plus"></i></span>
										</div>
										
									</div>
									
								<hr></hr>
									<div class="form-group row">
										<div class="col-sm-6">
											<em class="ClearForm">Очистить форму</em>
										</div>						
										
										
										<div class="col-sm-6">
										  <button type="submit" 
												class="btn btn-primary float-sm-right"
												onClick="return solaris_ValidPerson(this.form);" >Проверить</button>
										</div>
									</div>
									
									
									<hr></hr>
									
									<div class="alert alert-info" role="alert" id="solaris_documantation">
									  <h6><strong>Краткая справка</strong></h6>
									  Минимальный набор данных для поиска:
									  <ul>
										<li>Фамилия + Имя</li>
										<li>Фамилия + Год рождения</li>
										<li>ИНН</li>
										<li>Имя + Отчество + Полная дата рождения</li>
									  </ul>
									</div>
									
								</form>
								
								
							</div>	
								
							
						
						</div>	
						
												
			
					</div>						
				</div>
			</div>
    
			
		</div>
	</div>
 



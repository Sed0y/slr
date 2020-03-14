

 
  
    <div class="container-fluid">
		<!-- Breadcrumbs
		<ol class="breadcrumb">
			<li class="breadcrumb-item"><a href="#">Проверка</a></li>
			<li class="breadcrumb-item active">Стоп-лист</li>
		</ol>  
		-->
		<div class="row">
    	
			<div class="col-lg-6">
				<div class="card">
					<div class="card-header">
						<i class="fa"></i> Поиск по стоп-листу
					</div>
					<div class="card-body">
			
						<div class="alert alert-light" role="alert" id="MainInfo" style="display:none;">
							<span>&nbsp; </span>
							<br>
							<i>&nbsp;</i>
						</div>
		
	
						<ul class="nav nav-tabs" id="myTab" role="tablist">
						  <li class="nav-item">
							<a class="nav-link active" href="#person" data-toggle="tab">Физические лица</a>
						  </li>
						  <li class="nav-item">
							<a class="nav-link" href="#vehicle" data-toggle="tab">Транспорт</a>
						  </li>
						  <li class="nav-item">
							<a class="nav-link" href="#company" data-toggle="tab">Юридические лица</a>
						  </li>		  
						</ul>
		
						<div class="tab-content" style="padding-top:25px;">				
							
							<div class="tab-pane active" id="person">
								<div class="alert alert-danger" role="alert" id = "form_person_valid_message" style="display: none;"></div>			
								<form >										
									<div class="form-group row">
										<label for="surname" class="col-sm-3 col-form-label" >Фамилия:</label>
										<div class="col-sm-9">  					
											<input type="text" class="form-control" id="surname" >
										</div>		
									</div>

									<div class="form-group row">
										<label for="name" class="col-sm-3 col-form-label">Имя:</label>
										<div class="col-sm-9">
											<input type="text" class="form-control" id="name">
										</div>
									</div>

									<div class="form-group row">
										<label for="fathername" class="col-sm-3 col-form-label">Отчество:</label>
										<div class="col-sm-9">
											<input type="text" class="form-control" id="fathername">
										</div>
									</div>
									  
									<div class="form-group row">
										<label for="birthdate" class="col-sm-3 col-form-label">Дата:</label>
										<div class="col-sm-9">
											<input type="text" class="form-control" id="birthdate">
										</div>
									</div>
								
									<div class="form-group row">
										<div class="col-sm-12">
										  <button type="submit" 
												class="btn btn-secondary float-sm-right"
												onClick="return ValidPerson(this.form);" >Проверить</button>
										</div>
									  </div>
									
								</form>
								
								
							</div>	
								
							<div class="tab-pane" id="vehicle">
								
								<div class="alert alert-danger" role="alert" id = "form_vehicle_valid_message"></div>
								
								<form >	
									<div class="form-group row">
										<label for="gos" class="col-sm-2 col-form-label">Госномер:</label>
										<div class="col-sm-10">
										  <input type="text" class="form-control" id="gos" 						
											oninput="ToUpper(this)"
										  >
										</div>
									</div>

									<div class="form-group row">
										<label for="vin" class="col-sm-2 col-form-label">VIN:</label>
										<div class="col-sm-10">
											<input type="text" class="form-control" id="vin"
												oninput="ToUpper(this)"
												>
										</div>
									</div>
											
									<div class="form-group row">
										<div class="col-sm-12">
										  <button 	type="submit" 
													class="btn btn-secondary float-sm-right"
													onClick="return ValidVehicle(this.form);" >Проверить</button>
										</div>
									  </div>
									
								</form>
								
								
							</div>
							
							<div class="tab-pane" id="company">
								
								<div class="alert alert-danger" role="alert" id = "form_company_valid_message"></div>
								
								<form >	
									<div class="form-group row">
										<label for="inn" class="col-sm-2 col-form-label">ИНН</label>
										<div class="col-sm-10">
										  <input type="text" class="form-control" id="inn" >
										</div>
									</div>	
								
									<div class="form-group row">
										<div class="col-sm-12">
										  <button type="submit" 
												class="btn btn-secondary float-sm-right"
												onClick="return ValidCompany(this.form);" >Проверить</button>
										</div>
									  </div>
									
								</form>
							
								
							</div>		
							
						
						</div>	
						
						<hr>
						
						<div id="ResultTable">
							
						</div>
						
						<hr>
						
						<div id="ResultDetails">
							
						</div>
			
					</div>						
				</div>
			</div>
    
			<div class="col-lg-6">
				<div class="card">
					<div class="card-header">
						<i class="fa fa-user-secret"></i> Результаты поиска
					</div>


					<div class="card-body">
					
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
							<p class="card-part-header"><em><i class="fa fa-info-circle" style="color:#4f7ab1;"></i> Дополнительная информация</em></p>
							<p class="card-part-content">
								<label><strong>Категория:</strong></label> <span id="res_category">-</span>
								</br>
								<label><strong>Источник:</strong></label> <span id="res_source">-</span>
								</br>
								<label><strong>Описание:</strong></label> <span id="res_descr">-</span>
								</br>
								<label><strong>Адрес:</strong></label> <span id="res_address">-</span>
								</br>
							</p>
							
							
							
						</div>
						
					</div>	
					<div class="card-footer small text-muted"></div>
				</div>
			</div>
		</div>
	</div>
 



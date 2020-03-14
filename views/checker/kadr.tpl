

 
  
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
						<i class="fa"></i> Проведённые проверки ЦПМ
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
						<!--
						  <li class="nav-item">
							<a class="nav-link" href="#vehicle" data-toggle="tab">Транспорт</a>
						  </li>
						  <li class="nav-item">
							<a class="nav-link" href="#company" data-toggle="tab">Юридические лица</a>
						  </li>		  
						  -->
						</ul>
		
						<div class="tab-content" style="padding-top:25px;">				
							
							<div class="tab-pane active" id="kadr_person">
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
												onClick="return ValidKadrPerson(this.form);" >Проверить</button>
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
														
							<p class="card-part-header"><em><i class="fa fa-info-circle" style="color:#4f7ab1;"></i> Сводка</em></p>
							<p class="card-part-content">
								<label><strong>Источник:</strong></label> <span id="res_source">-</span>
								</br>
								<label><strong>Результат:</strong></label> <span id="res_result">-</span>
								</br>
								<label><strong>Под ответственность:</strong></label> <span id="res_guarantee">-</span>
								</br>
								<label><strong>Комментарий:</strong></label> <span id="res_comment">-</span>
								</br>
								<label><strong>Проверял:</strong></label> <span id="res_checked_by">-</span>
								</br>
								<label><strong>Дата проверки:</strong></label> <span id="res_checked_data">-</span>
								</br>
							</p>
							
							
							
						</div>
						
					</div>	
					<div class="card-footer small text-muted"></div>
				</div>
			</div>
		</div>
	</div>
 



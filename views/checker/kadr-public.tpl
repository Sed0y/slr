

 <!DOCTYPE html>

<html lang="en">

<head>
  <meta charset="utf-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">  
  <title>Стоп-лист ЦПМ</title>
  <!-- Bootstrap core CSS-->
  <link href="/static/js/vendor/bootstrap/css/bootstrap.min.css" rel="stylesheet">  
  <link href="/static/js/vendor/font-awesome/css/font-awesome.min.css" rel="stylesheet" type="text/css">  
  <link href="/static/js/vendor/datatables/dataTables.bootstrap4.css" rel="stylesheet">
  
	<style type="text/css">    

	
	
	</style>
  

</head>

<body style="background-color:#fbfbfb;">

	 <div style="width:100%;height: 120px;"></div>
	
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
						<i class="fa"></i> Проверки ЦПМ штатных сотрудников
					</div>
					<div class="card-body">
			
						<div class="alert alert-light" role="alert" id="MainInfo" style="display:none;">
							<span>&nbsp; </span>
							<br>
							<i>&nbsp;</i>
						</div>
		
	<!--
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
		-->
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
	
	<script>
		function capitalizeFirstLetter(string) {
			string = string.toLowerCase();
			return string.charAt(0).toUpperCase() + string.slice(1);
		}
	</script>
	
	<!-- Bootstrap core JavaScript-->
    <script src="/static/js/vendor/jquery/jquery.min.js"></script>
    <script src="/static/js/vendor/bootstrap/js/bootstrap.bundle.min.js"></script>
    <!-- Core plugin JavaScript-->
    <script src="/static/js/vendor/jquery-easing/jquery.easing.min.js"></script>
    <!-- Page level plugin JavaScript-->
    <!-- <script src="/static/js/vendor/chart.js/Chart.min.js"></script> -->
    <script src="/static/js/vendor/datatables/jquery.dataTables.js"></script>
    <script src="/static/js/vendor/datatables/dataTables.bootstrap4.js"></script>
    <!-- Custom scripts for all pages-->    
	<script src="/static/js/extension/jquery-ui.min.js"></script>
		<script src="/static/js/kadr.js"></script>
	
	
	</body>
</html>
 



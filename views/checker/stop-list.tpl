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
  <link href="/static/css/kadr.css" rel="stylesheet">
  
	<style type="text/css">    

    .nav-item a {      
		color: #9e9e9e;
    }
	
	.alert span {
		font-size: 0.9rem;
		font-weight: bold;
	}
	
	.alert {
		border-color:#ececec;
	}
	
	
	
	</style>
  
<!--
  <link href="/static/css/jquery-ui.theme.min.css" rel="stylesheet">
  <link href="/static/css/jquery-ui.structure.min.css" rel="stylesheet">
-->





</head>

<body style="background-color:#fbfbfb;">
  
  <div style="width:100%;height: 120px;"></div>
  <div class="container">
  <div class="row align-items-end">
  
    <div class="col-sm-3">
       
    </div>
		
	
    <div class="col-sm-6 align-self-center">
	
	<div class="alert alert-light" role="alert" id="MainInfo">
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
			<form >				
				<div class="form-group row">
					<label for="surname" class="col-sm-2 col-form-label" >Фамилия:</label>
					<div class="col-sm-10">  					
					  <input type="text" class="form-control" id="surname">
					</div>		
				</div>

				<div class="form-group row">
					<label for="name" class="col-sm-2 col-form-label">Имя:</label>
					<div class="col-sm-10">
						<input type="text" class="form-control" id="name">
					</div>
				</div>

				<div class="form-group row">
					<label for="fathername" class="col-sm-2 col-form-label">Отчество:</label>
					<div class="col-sm-10">
						<input type="text" class="form-control" id="fathername">
					</div>
				</div>
				  
				<div class="form-group row">
					<label for="birthdate" class="col-sm-2 col-form-label">Дата:</label>
					<div class="col-sm-10">
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
			
			<div class="alert alert-danger" role="alert" id = "form_person_valid_message" style="display: none;"></div>			
		</div>	
			
		<div class="tab-pane" id="vehicle">
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
			
			<div class="alert alert-danger" role="alert" id = "form_vehicle_valid_message"></div>
		</div>
		
		<div class="tab-pane" id="company">
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
		
			<div class="alert alert-danger" role="alert" id = "form_company_valid_message"></div>
		</div>		
		
	
	</div>				
		
    </div>
    <div class="col-sm-3">
      
    </div>
  </div>
</div>
  
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
		
	

 <script>
$(function () {

	$('#myTab a:first').tab('show')

	$("#form_person_valid_message").hide();
	$("#form_vehicle_valid_message").hide();
	$("#form_company_valid_message").hide();

	
	
	$("#person form").submit(function(event){

		$("#form_person_valid_message").hide();
		
		$("#MainInfo").removeClass();		
		$("#MainInfo").addClass("alert alert-light");
		$("#MainInfo span").html("&nbsp;");
		$("#MainInfo i").html("&nbsp;");
	
		event.preventDefault(); 
		
		//alert("Submit Person");			
		var v_surname = document.getElementById("surname").value;
		var v_name = document.getElementById("name").value;
		var v_fathername = document.getElementById("fathername").value;
		var v_birthdate = document.getElementById("birthdate").value;
	
		parts = v_birthdate.split('.');
		v_birthdate = parts[2] + "-" + parts[1] + "-" + parts[0];
	
		var formData = {
						surname:v_surname,												
						name:v_name,
						fathername:v_fathername,
						birthdate:v_birthdate
					}; 
		
		$.ajaxSetup({cache: false}); 	

		$.ajax({
		  type: "POST",
		  url: "/stop/simple/person",
		  data: formData,			  
		  success: function(data){	
			
			ShowPersonResults(data);
		  }
		});		
	}); 
	
	$("#vehicle form").submit(function(event){

		$("#form_vehicle_valid_message").hide();
		
		$("#MainInfo").removeClass();		
		$("#MainInfo").addClass("alert alert-light");
		$("#MainInfo span").html("&nbsp;");
		$("#MainInfo i").html("&nbsp;");
	
		event.preventDefault(); 
		
		//alert("Submit Person");			
		var v_vin = document.getElementById("vin").value;
		var v_gos = document.getElementById("gos").value;		
			
		var formData = {
						vin:v_vin,												
						gos:v_gos
					}; 
		
		$.ajaxSetup({cache: false}); 	

		$.ajax({
		  type: "POST",
		  url: "/stop/simple/vehicle",
		  data: formData,			  
		  success: function(data){				
			ShowVehicleResults(data);
		  }
		});		
	}); 
	
	$("#company form").submit(function(event){

		$("#form_company_valid_message").hide();
		
		$("#MainInfo").removeClass();		
		$("#MainInfo").addClass("alert alert-light");
		$("#MainInfo span").html("&nbsp;");
		$("#MainInfo i").html("&nbsp;");
	
		event.preventDefault(); 
		
		//alert("Submit Person");			
		var v_inn = document.getElementById("inn").value;
					
		var formData = {
						inn:v_inn
					}; 
		
		$.ajaxSetup({cache: false}); 	

		$.ajax({
		  type: "POST",
		  url: "/stop/simple/company",
		  data: formData,			  
		  success: function(data){				
			ShowCompanyResults(data);
		  }
		});		
	}); 

	
})
 

 function ShowPersonResults (data){
		
	var birthdate = document.getElementById("birthdate").value;
	var arr = data.split('@')
		
	//alert(data);
	
	if (arr[0] == 0){		
		$("#MainInfo").removeClass();		
		$("#MainInfo").addClass("alert alert-success");
		$("#MainInfo span").html(
			document.getElementById("surname").value + " "
			+ document.getElementById("name").value + " " 
			+ document.getElementById("fathername").value + " " 
			+ birthdate
		);
		$("#MainInfo i").html("Совпадений не найдено");
		return;
	}
	
	if (arr[0] == -1){		
		$("#MainInfo").removeClass();		
		$("#MainInfo").addClass("alert alert-danger");
		$("#MainInfo span").html(
			document.getElementById("surname").value + " "
			+ document.getElementById("name").value + " " 
			+ document.getElementById("fathername").value + " " 
			+ birthdate
		);
		$("#MainInfo i").html("Есть несколько совпадений в стоп-листе");
		return;
	}
	
	if (arr[0] > 0){
	
		if (arr[2] > 0 && arr[2] < 100){		
			$("#MainInfo").removeClass();		
			$("#MainInfo").addClass("alert alert-warning");
			$("#MainInfo span").html(
				document.getElementById("surname").value + " "
				+ document.getElementById("name").value + " " 
				+ document.getElementById("fathername").value + " " 
				+ birthdate
			);
			$("#MainInfo i").html(arr[0] + " " + arr[1]);
			
			return;
		}
		
		if (arr[2] == 100){		
			$("#MainInfo").removeClass();		
			$("#MainInfo").addClass("alert alert-danger");
			$("#MainInfo span").html(
				document.getElementById("surname").value + " "
				+ document.getElementById("name").value + " " 
				+ document.getElementById("fathername").value + " " 
				+ birthdate
			);
			$("#MainInfo i").html(arr[0] + " " + arr[1]);
			
			return;
		}
		
	}
 
 }
 
 
 function ShowVehicleResults (data){
		
	var birthdate = document.getElementById("birthdate").value;
	var arr = data.split('@')
		
	//alert(data);
	
	if (arr[0] == 0){		
		$("#MainInfo").removeClass();		
		$("#MainInfo").addClass("alert alert-success");
		$("#MainInfo span").html(
			document.getElementById("vin").value + " "
			+ document.getElementById("gos").value 
		);
		$("#MainInfo i").html("Совпадений не найдено");
		return;
	}
	
	if (arr[0] == -1){		
		$("#MainInfo").removeClass();		
		$("#MainInfo").addClass("alert alert-danger");
		$("#MainInfo span").html(
			document.getElementById("vin").value + " "
			+ document.getElementById("gos").value 
		);
		$("#MainInfo i").html("Есть несколько совпадений в стоп-листе");
		return;
	}
	
	if (arr[0] > 0){
	
		if (arr[2] > 0 && arr[2] < 100){		
			$("#MainInfo").removeClass();		
			$("#MainInfo").addClass("alert alert-warning");
			$("#MainInfo span").html(
				document.getElementById("vin").value + " "
				+ document.getElementById("gos").value 
			);
			$("#MainInfo i").html(arr[0] + " " + arr[1]);
			
			return;
		}
		
		if (arr[2] == 100){		
			$("#MainInfo").removeClass();		
			$("#MainInfo").addClass("alert alert-danger");
			$("#MainInfo span").html(
				document.getElementById("vin").value + " "
				+ document.getElementById("gos").value 
			);
			$("#MainInfo i").html(arr[0] + " " + arr[1]);
			
			return;
		}
		
	}
 
 }
 
 function ShowCompanyResults (data){
		
	
	var arr = data.split('@')
		
	//alert(data);
	
	if (arr[0] == 0){		
		$("#MainInfo").removeClass();		
		$("#MainInfo").addClass("alert alert-success");
		$("#MainInfo span").html(
			document.getElementById("inn").value
		);
		$("#MainInfo i").html("Совпадений не найдено");
		return;
	}
	
	if (arr[0] == -1){		
		$("#MainInfo").removeClass();		
		$("#MainInfo").addClass("alert alert-danger");
		$("#MainInfo span").html(
			document.getElementById("inn").value
		);
		$("#MainInfo i").html("Есть несколько совпадений в стоп-листе");
		return;
	}
	
	if (arr[0] > 0){
	
		if (arr[2] > 0 && arr[2] < 100){		
			$("#MainInfo").removeClass();		
			$("#MainInfo").addClass("alert alert-warning");
			$("#MainInfo span").html(
				document.getElementById("inn").value
			);
			$("#MainInfo i").html(arr[0] + " " + arr[1]);
			
			return;
		}
		
		if (arr[2] == 100){		
			$("#MainInfo").removeClass();		
			$("#MainInfo").addClass("alert alert-danger");
			$("#MainInfo span").html(
				document.getElementById("inn").value
			);
			$("#MainInfo i").html(arr[0] + " " + arr[1]);
			
			return;
		}
		
	}
 
 }
 
function ValidCompany(data){
	
	$("#form_company_valid_message").hide();
	
	var result = true;	
	
	var message  = "";
		
	var inn = document.getElementById("inn").value;	
	
	if(inn == "")
	{
		if (message != "" )
			message += "</br></br>";		
			
		message += "Не заполнены поля поиска";					
		result = false;
	}
			
	if(inn != "" && !(/^[0-9]{10}$/.test(inn)) )	{
	
		if (message != "" )
			message += "</br></br>";
			
		message = message + "ИНН должен состоять из 10 цифр</br>";				
		result = false;
	}
	
	
	if(result == false)	{
		$("#form_company_valid_message").html(message);
		$("#form_company_valid_message").show();
		
	}
		
	return result;	
}	
  
   
function ValidVehicle(data){
	
	$("#form_vehicle_valid_message").hide();
	
	var result = true;	
	
	var message  = "";
		
	var vin = document.getElementById("vin").value;
	var gos = document.getElementById("gos").value;
	
	if(vin == "" && gos == "")
	{
		if (message != "" )
			message += "</br></br>";		
			
		message += "Не заполнены поля поиска";		
			
		result = false;
	}
			
	if(vin != "" && !(/^[0-9A-HJ-NPR-Z]{17}$/.test(vin)) )
	{
		if (message != "" )
			message += "</br></br>";		
		message += "<strong>Неверный VIN номер</strong><br>Может содержать цифры и латинские символы, кроме O, Q, I ";		
		
		result = false;
	}
	
	if( gos != "")
	{
		if( !(/^[А,В,Е,И,К,М,Н,О,Р,С,Т,У,Х][0-9]{3}[А,В,Е,И,К,М,Н,О,Р,С,Т,У,Х]{2}[0-9]{2,3}$/.test(gos)) && 
			!(/^[А,В,Е,И,К,М,Н,О,Р,С,Т,У,Х]{2}[0-9]{3}[А,В,Е,И,К,М,Н,О,Р,С,Т,У,Х]{1}[0-9]{1,3}$/.test(gos)) &&
			!(/^[А,В,Е,И,К,М,Н,О,Р,С,Т,У,Х]{2}[0-9]{6,7}$/.test(gos))			
		  )
		{
			if (message != "" )
				message += "</br></br>";				
			
			message += '<strong>Неверный номер государственной регистрации</strong> </br> а) Может содержать только цифры и русские символы</br>б) Неверный формат номера';	
			
			result = false;
		}
	}
	
	if(result == false)	{
		$("#form_vehicle_valid_message").html(message);
		$("#form_vehicle_valid_message").show();
		
	}
		
	return result;	
}
 
 
function ValidPerson(data){
	
	
	$("#form_person_valid_message").hide();
	var result = true;	
	
	var error_index = 1;
	var message  = "";
	
	var surname = document.getElementById("surname").value;
	var name = document.getElementById("name").value;
	var fathername = document.getElementById("fathername").value;
	var birthdate = document.getElementById("birthdate").value;
		
	
	
	if(surname == "")
	{
		message = message + error_index + ". Фамилия обязательна для ввода</br>";
		error_index++;
		result = false;
	}
	
	if(birthdate == "")
	{
		message = message + error_index + ". Дата рождения обязательна для ввода</br>";
		error_index++;
		result = false;
	}
	
	/*
	if(birthdate == "" && (name == "" || fathername == "") )
	{
		message = message + error_index + ". Нужно ввести ещё либо ДР, либо полностью ФИО</br>";
		error_index = error_index + 1;
		result = false;
	}
	*/
		
	if(birthdate != "")
	{		
		if(birthdate != "" && !(/^[0-9]{2}.[0-9]{2}.[0-9]{4}$/.test(birthdate)) )
		{
			message = message + error_index + ". Дата введена неверно, формат дд.мм.гггг</br>";
			error_index = error_index + 1;
			result = false;
		}
		else
		{
			if(!isValidDate(birthdate))
			{
				message = message + error_index + ". Такой даты не существует</br>";
				error_index = error_index + 1;
				result = false;
			}
			else
			{
				now = new Date();
				m = birthdate.split(".");
				d_bd = new Date(m[2],m[1]-1,m[0]);
				d_now = new Date(now.getFullYear(),now.getMonth(),now.getDate());
										
				if(d_bd > d_now)
				{
					var message = message + error_index + ". День рождения в будущем?</br>";
					error_index = error_index + 1;
					result = false;
				}
			}
		}
	}
	
	if(surname != "" && !(/^[а-яёА-ЯЁ\-]+$/.test(surname)) )
	{
		message = message + error_index + ". Фамилия может содержать символы русского алфавита и тире</br>";
		error_index = error_index + 1;
		result = false;
	}
	
	if(name != "" && !(/^[а-яёА-ЯЁ\-]+$/.test(name)) )
	{
		message = message + error_index + ". Имя может содержать символы русского алфавита и тире</br>";
		error_index = error_index + 1;
		result = false;
	}
	
	if(fathername != "" && !(/^[а-яёА-ЯЁ\-]+$/.test(fathername)) )
	{
		message = message + error_index + ". Отчество может содержать символы русского алфавита и тире</br>";
		error_index = error_index + 1;
		result = false;
	}
	
	if(result == false)	{
		$("#form_person_valid_message").html(message);
		$("#form_person_valid_message").show();
		
	}		
		
	return result;	
}


function isValidDate (str){

	if (!/^\d\d\.\d\d\.\d{4}$/.test (str)) {alert ('Error, unformat'); return}
	var a0 = function (x) {return ((x < 10) ? '0' : '') + x},
	t = str.split ('.'), ndt = new Date (+t [2], t[1] - 1, +t [0]);
	with (ndt) var tst = [a0 (getDate ()), a0 (getMonth () + 1), getFullYear ()].join ('.');
	
	if (tst != str) 	
	 return false;
	else
	 return true;	
}


function ToUpper(val){
	val.value = val.value.toUpperCase();
}




 </script>
</body>
</html>

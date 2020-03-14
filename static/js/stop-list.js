

var Results


jQuery(document).ready(function() 
{
	
	
	$('#birthdate').mask('99.99.9999');
	
	$('#myTab a:first').tab('show')

	$("#form_person_valid_message").hide();
	$("#form_vehicle_valid_message").hide();
	$("#form_company_valid_message").hide();
	
	
	
	$("#person form").submit(function(event){

		CleanResultCard();		
		$("#form_person_valid_message").hide();
		
		$("#MainInfo").removeClass();		
		$("#MainInfo").addClass("alert alert-light");
		$("#MainInfo span").html("&nbsp;");
		$("#MainInfo i").html("&nbsp;");
	
		event.preventDefault(); 
						
		var v_surname = document.getElementById("surname").value;
		var v_name = document.getElementById("name").value;
		var v_fathername = document.getElementById("fathername").value;
		var v_birthdate = document.getElementById("birthdate").value;
	
		if(v_birthdate != '')
		{
			parts = v_birthdate.split('.');
			v_birthdate = parts[2] + "-" + parts[1] + "-" + parts[0];
		}
	
		var formData = {
						surname:v_surname,												
						name:v_name,
						fathername:v_fathername,
						birthdate:v_birthdate
					}; 
		
		$.ajaxSetup({cache: false}); 	

		$.ajax({
		  type: "POST",
		  url: "/stop/extended/person",
		  data: formData,			  
		  success: function(data){				
			PrintPersonResultTable(data);			
		  }
		});		
	}); 
	
	$("#vehicle form").submit(function(event){

		CleanResultCard();
		$("#form_vehicle_valid_message").hide();
		
		$("#MainInfo").removeClass();		
		$("#MainInfo").addClass("alert alert-light");
		$("#MainInfo span").html("&nbsp;");
		$("#MainInfo i").html("&nbsp;");
	
		event.preventDefault(); 		
				
		var v_vin = document.getElementById("vin").value.trim();
		var v_gos = document.getElementById("gos").value.trim();		
				
		var formData = {
						vin:v_vin,												
						gos:v_gos
					}; 
		
		$.ajaxSetup({cache: false}); 	

		$.ajax({
		  type: "POST",
		  url: "/stop/extended/vehicle",
		  data: formData,			  
		  success: function(data){				
			PrintVehicleResultTable(data);
		  }
		});		
	}); 
	
	$("#company form").submit(function(event){

		CleanResultCard();
		$("#form_company_valid_message").hide();
		
		$("#MainInfo").removeClass();		
		$("#MainInfo").addClass("alert alert-light");
		$("#MainInfo span").html("&nbsp;");
		$("#MainInfo i").html("&nbsp;");
	
		event.preventDefault(); 
						
		var v_inn = document.getElementById("inn").value;
					
		var formData = {
						inn:v_inn
					}; 
		
		$.ajaxSetup({cache: false}); 	

		$.ajax({
		  type: "POST",
		  url: "/stop/extended/company",
		  data: formData,			  
		  success: function(data){				
			PrintCompanyResultTable(data);
		  }
		});		
	}); 

	
});

function CleanResultCard() {
	
	$("#res_fio").text("-");
	$("#res_dr").text("-");
	$("#res_add_date").text("-");
	$("#res_reason").text("-");
	$("#res_category").text("-");
	$("#res_source").text("-");
	$("#res_address").text("-");
	
	
}


function capitalizeFirstLetter(string) {
	string = string.toLowerCase();
    return string.charAt(0).toUpperCase() + string.slice(1);
}

function PrintPersonResultTable (data){
		
	$("#ResultTable").text("");	
	//var Results = $.parseJSON(data);
	Results = $.parseJSON(data);
	
	//alert("here");
	var render = "";
	var fin_column = -1;
	var id_column = -1;
	
	render += "<table id=\"DetailsTable\" style=\"width:100%\" class=\"table table-bordered compact\" cellspacing=\"0\">";
	render += "	<thead>";
	render += "		<tr>";
	render += "		   <td>№</td>";
	render += "		   <td>ФИО</td>";
	render += "		   <td>Дата рождения</td>";
	render += "		</tr>";
	render += "	</thead>";
	render += "	<tbody>";
	
	var index = 1
	
	var newdate
	var dateparts	
	var init_element
	
	for (var part in Results) {
		//alert(Results[part]);
		render += "		<tr>";
		render += "		   <td >" + index + "</td>";
		render += "		   <td><span id=\"slperson_" + Results[part].pid + "\" onClick=\"PrintDetails(this,'no');\">" + 	capitalizeFirstLetter(Results[part]["sname"].trim()) + " " + 
									capitalizeFirstLetter(Results[part].name.trim()) + " " + 
									capitalizeFirstLetter(Results[part].fname.trim()) + "</span></td>";
		
		dateparts = Results[part]["bd"].split('T');
		dateparts = dateparts[0].split('-');		
		
		if (index == 1)
			PrintDetails(null,"slperson_" + Results[part].pid);
		
		render += "		   <td> " + dateparts[2] + "." + dateparts[1] + "." + dateparts[0] + " </td>";
		render += "		</tr>";
		
		index++;
	}
	
	render += "	</tbody>";
	render += "</table>";
	
	
	$("#ResultTable").append(render);
	
	
	$('#DetailsTable').DataTable({
		"ordering": false,
		"language": {
			"search": "Поиск:",
            "lengthMenu": "По _MENU_ записей на страницу",			
            "zeroRecords": "Ничего не найдено",
            "info": "Страница _PAGE_ из _PAGES_",
            "infoEmpty": "Нет записей",			
			"paginate": {
				"first":      "Первая",
				"previous":   "Предыдущая",
				"next":       "Следующая",
				"last":       "Последняя"
			},            
        }
		
	});
	
	
		
	
}


function PrintVehicleResultTable (data){
		
	//alert(data);
	//return;
		
	$("#ResultTable").text("");	
	
	Results = $.parseJSON(data);
	
	//alert("here");
	var render = "";
	var fin_column = -1;
	var id_column = -1;
	
	render += "<table id=\"DetailsTable\" style=\"width:100%\" class=\"table table-bordered compact\" cellspacing=\"0\">";
	render += "	<thead>";
	render += "		<tr>";
	render += "		   <td>№</td>";
	render += "		   <td>VIN</td>";
	render += "		   <td>Гос номер</td>";
	render += "		</tr>";
	render += "	</thead>";
	render += "	<tbody>";
	
	var index = 1
	
	var newdate
	var dateparts	
	
	for (var part in Results) {
		//alert(Results[part]);
		render += "		<tr>";
		render += "		   <td >" + index + "</td>";
		render += "		   <td><span id=\"slvehicle_" + Results[part].vid + "\" onClick=\"PrintDetails(this,'no');\">" + 	
									Results[part].vin + 
								"</span></td>";
		
		if (index == 1)
			PrintDetails(null,"slvehicle_" + Results[part].vid);
		
		render += "		   <td><span id=\"slvehicle_" + Results[part].vid + "\" onClick=\"PrintDetails(this,'no');\">" + Results[part].gos + " </td>";
		render += "		</tr>";
		
		index++;
	}
	
	render += "	</tbody>";
	render += "</table>";
	
	
	$("#ResultTable").append(render);
	
	
	$('#DetailsTable').DataTable({
		"ordering": false,
		"language": {
			"search": "Поиск:",
            "lengthMenu": "По _MENU_ записей на страницу",			
            "zeroRecords": "Ничего не найдено",
            "info": "Страница _PAGE_ из _PAGES_",
            "infoEmpty": "Нет записей",			
			"paginate": {
				"first":      "Первая",
				"previous":   "Предыдущая",
				"next":       "Следующая",
				"last":       "Последняя"
			},            
        }
		
	});
		
	
}


function PrintCompanyResultTable (data){
		
	//alert(data);
	//return;
		
	$("#ResultTable").text("");	
	
	Results = $.parseJSON(data);
	
	//alert("here");
	var render = "";
	var fin_column = -1;
	var id_column = -1;
	
	render += "<table id=\"DetailsTable\" style=\"width:100%\" class=\"table table-bordered compact\" cellspacing=\"0\">";
	render += "	<thead>";
	render += "		<tr>";
	render += "		   <td>№</td>";
	render += "		   <td>ИНН</td>";
	render += "		   <td>Название</td>";
	render += "		</tr>";
	render += "	</thead>";
	render += "	<tbody>";
	
	var index = 1
	
	var newdate
	var dateparts	
	
	var cname
	
	for (var part in Results) {
		//alert(Results[part]);
		render += "		<tr>";
		render += "		   <td >" + index + "</td>";
		render += "		   <td><span id=\"slcompany_" + Results[part].cid + "\" onClick=\"PrintDetails(this,'no');\">" + 	
									Results[part].inn + 
								"</span></td>";
		
		if (index == 1)
			PrintDetails(null,"slcompany_" + Results[part].cid);
		
		cname = Results[part].name.replace(/\[#Q\]/g,"\"") 		
		
		render += "		   <td><span id=\"slcompany_" + Results[part].cid + "\" onClick=\"PrintDetails(this,'no');\">" + cname + " </td>";
		render += "		</tr>";
		
		index++;
	}
	
	render += "	</tbody>";
	render += "</table>";
	
	
	$("#ResultTable").append(render);
	
	
	$('#DetailsTable').DataTable({
		"ordering": false,
		"language": {
			"search": "Поиск:",
            "lengthMenu": "По _MENU_ записей на страницу",			
            "zeroRecords": "Ничего не найдено",
            "info": "Страница _PAGE_ из _PAGES_",
            "infoEmpty": "Нет записей",			
			"paginate": {
				"first":      "Первая",
				"previous":   "Предыдущая",
				"next":       "Следующая",
				"last":       "Последняя"
			},            
        }
		
	});
		
	
}


function PrintDetails(element, init_element){
		
	var id
	var type
	
	if(init_element == "no"){
	
		id = element.id.split("_")[1];
		type = element.id.split("_")[0];
	} else {
		id = init_element.split("_")[1];
		type = init_element.split("_")[0];
	}
	
	
	if(type == "slperson"){	
	
		for (var part in Results) {
			if(Results[part].pid == id){
				
				$("#res_fio").text(
					capitalizeFirstLetter(Results[part]["sname"].trim()) + " " + 
					capitalizeFirstLetter(Results[part].name.trim()) + " " + 
					capitalizeFirstLetter(Results[part].fname.trim()));
				
				dateparts = Results[part]["bd"].split('T');
				dateparts = dateparts[0].split('-');		
			
			
				$("#res_dr").text(dateparts[2] + "." + dateparts[1] + "." + dateparts[0]);
				
			
				var reason = Results[part].reason.replace(/\[#NL\]/g,"</br>") 
				reason = reason.replace(/\[#Q\]/g,"\"") 
				
				$("#res_add_date").text(Results[part].add);
				$("#res_reason").html(reason);
				$("#res_category").text(Results[part].cat);
				$("#res_source").text(Results[part].source);
				$("#res_address").text(Results[part].address);
				
				break;
			}
		}
	}
	
	if(type == "slvehicle"){	
	
		for (var part in Results) {
			if(Results[part].vid == id){
				
				$("#res_fio").text(Results[part]["vin"]);				
				$("#res_dr").text(Results[part]["gos"]);
				
			
				var reason = Results[part].reason.replace(/\[#NL\]/g,"</br>") 
				reason = reason.replace(/\[#Q\]/g,"\"") 

				dateparts = Results[part]["add"].split('T');
				dateparts = dateparts[0].split('-');		

				
				$("#res_add_date").text(dateparts[2] + "." + dateparts[1] + "." + dateparts[0]);
				$("#res_reason").html(reason);
				$("#res_category").text(Results[part].cat);
				$("#res_source").text(Results[part].source);
				$("#res_descr").text(Results[part].desc);
				
				break;
			}
		}
	}
	
	if(type == "slcompany"){	
	
		for (var part in Results) {
			if(Results[part].cid == id){
				
				$("#res_fio").text(Results[part]["inn"]);				
			
				var reason = Results[part].reason.replace(/\[#NL\]/g,"</br>") 
				reason = reason.replace(/\[#Q\]/g,"\"") 

				dateparts = Results[part]["add"].split('T');
				dateparts = dateparts[0].split('-');		

				
				$("#res_add_date").text(dateparts[2] + "." + dateparts[1] + "." + dateparts[0]);
				$("#res_reason").html(reason);
				$("#res_category").text(Results[part].cat);
				$("#res_source").text(Results[part].source);
				$("#res_descr").text("ОГРН - " + Results[part].ogrn);
				
				break;
			}
		}
	}
	
}


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
		
	
	//var birthdate = document.getElementById("birthdate").value;
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
	else
	{
	
		if(birthdate == "" && name == "" && fathername == "")
		{
			message = message + error_index + ". Одной фамилии для поиска недостаточно</br>";
			error_index++;
			result = false;
		}
		
		if(birthdate == "" && name != "" && fathername == "") 
		{
			message = message + error_index + ". Фамилии и имени недостаточно для поиска</br>";
			error_index++;
			result = false;
		}
		
		if(birthdate == "" && name == "" && fathername != "") 
		{
			message = message + error_index + ". Фамилии и отчества недостаточно для поиска</br>";
			error_index++;
			result = false;
		}
	}
	
		
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

	if (!/^\d\d\.\d\d\.\d{4}$/.test (str)) {/*alert ('Error, unformat');*/ return false;}
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




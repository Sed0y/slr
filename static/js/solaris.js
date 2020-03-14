

var SolResults




jQuery(document).ready(function() 
{
	setInterval(CheckAuth, 5*60*1000); 
	
	solaris_HideAllRequestParameter();	
	
	
	$("button").prop( "disabled", false ); 
	$("#WaitTable").hide();
			
		
	$('#myTab a:first').tab('show');
	$("#solaris_form_person_valid_message").hide();
	
	$(".ClearForm").click(function(){
		
		$("#solaris_surname").val("");
		$("#solaris_name").val("");
		$("#solaris_fathername").val("");
		
		$("#solaris_birthdate_day").val("");
		$("#solaris_birthdate_month").val("");
		$("#solaris_birthdate_year").val("");
		
		$("#solaris_inn").val("");
		
	});
	
	$("#solaris_person form").submit(function(event){

		CleanResultCard();	

		
		
		$("#solaris_form_person_valid_message").hide();
		
		var v_surname = ""
		var v_name = ""
		var v_fathername = ""
		var v_birthdate_day = ""
		var v_birthdate_month = ""
		var v_birthdate_year = ""
		var v_inn = ""
	
		event.preventDefault(); 
						
		el = $("#req_person_surname");
		if(el.css("display") != "none")
			v_surname = $("#req_person_surname .req_value").html();
		
		el = $("#req_person_name");
		if(el.css("display") != "none")
			v_name = $("#req_person_name .req_value").html();	
		
		el = $("#req_person_fathername");
		if(el.css("display") != "none")
			v_fathername = $("#req_person_fathername .req_value").html();
		
		el = $("#req_person_birthdate");
		if(el.css("display") != "none"){
			
			var full_bd = $("#req_person_birthdate .req_value").html();
			full_bd = full_bd.replace(" г.р.", "");
			var full_bd_array = full_bd.split('.');
			
			v_birthdate_day = full_bd_array[0].replace(/_/g,"");
			v_birthdate_month = full_bd_array[1].replace(/_/g,"");
			v_birthdate_year = full_bd_array[2].replace(/_/g,"");
			
		}

		el = $("#req_person_inn");
		if(el.css("display") != "none")
			v_inn = $("#req_person_inn .req_value").html();
		

		var formData = {
						surname:v_surname,												
						name:v_name,
						fathername:v_fathername,
						bd_day: v_birthdate_day,
						bd_month: v_birthdate_month,
						bd_year: v_birthdate_year,
						inn: v_inn
					}; 
		
		$.ajaxSetup({cache: false}); 	

		$("#SolarisResultTable").text("");
		$.ajax({
		  type: "POST",
		  url: "/solaris/people/full_request",
		  data: formData,			  
		  success: function(data){	
			if(data == "redirect"){
				window.location.replace("/login");
				return;
			}
			
			solaris_PrintPersonResultTable(data);	
			$("button").prop( "disabled", false ); 
			$("#WaitTable").hide();
		  }
		});	

		$("button").prop( "disabled", true ); 
		$("#WaitTable").show();
	}); 
	
	//$( "selector" ).on( "event", function(){return false} )
	
	
	
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

function CheckAuth() {
	
	
	$.ajaxSetup({cache: false}); 	
	
	var formData = {
						nothing:"empty"
					}; 
					
	$.ajax({
	  type: "POST",
	  url: "/auth/check",
	  data: formData,			  
	  success: function(data){
		  
		if(data == "redirect"){
			window.location.replace("/login");
			return;
		}
		
	  }
	});	
			
		
		
}

function capitalizeFirstLetter(string) {
	string = string.toLowerCase();
    return string.charAt(0).toUpperCase() + string.slice(1);
}


function solaris_PrintPersonResultTable (data){
	
			
	$("#SolarisResultTable").text("");			
	
	var iferr = data.split("###");
		
	if(iferr.length == 1) { 	
	
		if (data == "nothing"){
			var empty = ""
			
			empty += "<div class=\"card\" style=\"margin-top:-10px; padding-top: 10px;\"><div class=\"card-body\"><p class=\"card-part-header\"><em><i class=\"fa fa-exclamation-triangle\" style=\"color:#a20b0b;\"></i>  По запросу ничего не найдено...</em></p></div></div>"

			$("#SolarisResultTable").append(empty);
		} else {
			$("#SolarisResultTable").append(data);
		}
	} else {
		
		empty += "<div class=\"card\" style=\"margin-top:-10px; padding-top: 10px;\"><div class=\"card-body\"><p class=\"card-part-header\"><em><i class=\"fa fa-exclamation-triangle\" style=\"color:#a20b0b;\"></i>  Ошибка выполнения запроса: </br> " + iferr[1] + "</em></p></div></div>"
	}
	
	
		
	
}


function solaris_ValidPerson(data){
	
	$("#solaris_form_person_valid_message").hide();

	var el
	
	var surname = ""
	var name = ""
	var fathername = ""
	var birthdate_day = ""
	var birthdate_month = ""
	var birthdate_year = ""
	var inn = ""
	
	
	el = $("#req_person_surname");
	if(el.css("display") != "none")
		surname = $("#req_person_surname .req_value").html();
	
	el = $("#req_person_name");
	if(el.css("display") != "none")
		name = $("#req_person_name .req_value").html();	
	
	el = $("#req_person_fathername");
	if(el.css("display") != "none")
		fathername = $("#req_person_fathername .req_value").html();
	
	el = $("#req_person_birthdate");
	if(el.css("display") != "none"){
		
		var full_bd = $("#req_person_birthdate .req_value").html();
		full_bd = full_bd.replace(" г.р.", "");
		var full_bd_array = full_bd.split('.');
		
		birthdate_day = full_bd_array[0].replace(/_/g,"");
		birthdate_month = full_bd_array[1].replace(/_/g,"");
		birthdate_year = full_bd_array[2].replace(/_/g,"");
		
	}

	el = $("#req_person_inn");
	if(el.css("display") != "none")
		inn = $("#req_person_inn .req_value").html();
	

	
	if( inn != "" )		
		return true;
	
	if( surname != "" && name != "" )
		return true;
	
	if( surname != "" && birthdate_year != "" )
		return true;
	
	if( name != "" && 
		fathername != "" &&
		birthdate_day != "" &&
		birthdate_month != "" &&
		birthdate_year != "" )
			return true;

			
	
	$("#solaris_form_person_valid_message").html("Недостаточно данных для поиска");
	$("#solaris_form_person_valid_message").show();
	
	
	return false;
		
}


function solaris_HideRequestParameter(row_id) {

	$("#"+row_id).hide();		
	solaris_CheckRequestParametrField();
}


function solaris_CheckRequestParametrField() {
	
	
	var nothing = true;
	var dsp = "none";
	$("#RequestParameters table tbody tr").each(function (index, element) {	
		dsp = $(element).css("display");		
		if (dsp != "none"){
			nothing = false;
		}
	});
	
	
	if(nothing) {		
		$("#PanelTitle").show();
	} else {
		$("#PanelTitle").hide();
	}

}


function solaris_HideAllRequestParameter() {
	$("#RequestParameters tr").hide();
	solaris_CheckRequestParametrField();
}

function solaris_AddRequestParameter(row_id) {

	
	solaris_CheckRequestParametrField();
	
	
	var error_index = 1;
	var message = "";
	var result = true;
	
	$("#solaris_form_person_valid_message").html("");
	$("#solaris_form_person_valid_message").hide();	
	
	if (row_id == "req_person_surname") {
		
		var v_surname = document.getElementById("solaris_surname").value;		
		v_surname = v_surname.trim();
				
		if(v_surname == "")
		{
			message = message + error_index + ". Пустое значение фамилии</br>";
			error_index++;
			result = false;
		}
		
		if(v_surname != "" && !(/^[а-яёА-ЯЁ\-]+$/.test(v_surname)) )
		{
			message = message + error_index + ". Недопустимые символы в фамилии</br>";
			error_index = error_index + 1;
			result = false;
		}		
		
		if(result == false)	{
			$("#solaris_form_person_valid_message").html(message);
			$("#solaris_form_person_valid_message").show();		
		} else {
			
			$("#solaris_form_person_valid_message").html("");
			$("#solaris_form_person_valid_message").hide();	
			
			
			$("#" + row_id +  " td.req_value").html(v_surname);
			$("#" + row_id).show();
		}

	}
	
	
	if (row_id == "req_person_name") {
		
		var v_name = document.getElementById("solaris_name").value;		
		v_name = v_name.trim();
				
		if(v_name == "")
		{
			message = message + error_index + ". Пустое значение имени</br>";
			error_index++;
			result = false;
		}
		
		if(v_name != "" && !(/^[а-яёА-ЯЁ]+$/.test(v_name)) )
		{
			message = message + error_index + ". Недопустимые символы в имени</br>";
			error_index = error_index + 1;
			result = false;
		}		
		
		if(result == false)	{
			$("#solaris_form_person_valid_message").html(message);
			$("#solaris_form_person_valid_message").show();		
		} else {
			
			$("#solaris_form_person_valid_message").html("");
			$("#solaris_form_person_valid_message").hide();	
			
			
			$("#" + row_id +  " td.req_value").html(v_name);
			$("#" + row_id).show();
		}

	}
	
		
	if (row_id == "req_person_fathername") {
		
		var v_fathername = document.getElementById("solaris_fathername").value;		
		v_fathername = v_fathername.trim();
				
		if(v_fathername == "")
		{
			message = message + error_index + ". Пустое значение отчества</br>";
			error_index++;
			result = false;
		}
		
		if(v_fathername != "" && !(/^[а-яёА-ЯЁ]+$/.test(v_fathername)) )
		{
			message = message + error_index + ". Недопустимые символы в отчестве</br>";
			error_index = error_index + 1;
			result = false;
		}		
		
		if(result == false)	{
			$("#solaris_form_person_valid_message").html(message);
			$("#solaris_form_person_valid_message").show();		
		} else {
			
			$("#solaris_form_person_valid_message").html("");
			$("#solaris_form_person_valid_message").hide();	
			
			
			$("#" + row_id +  " td.req_value").html(v_fathername);
			$("#" + row_id).show();
		}

	}
	
	
	if (row_id == "req_person_inn") {
		
		var v_inn = document.getElementById("solaris_inn").value;		
		v_inn = v_inn.trim();
				
		if(v_inn == "")
		{
			message = message + error_index + ". Пустое значение ИНН ФЛ</br>";
			error_index++;
			result = false;
		}
		
		if(v_inn != "" && !(/\d{12}/.test(v_inn)) )
		{
			message = message + error_index + ". Некорректный номер ИНН ФЛ</br>";
			error_index = error_index + 1;
			result = false;
		}		
		
		if(result == false)	{
			$("#solaris_form_person_valid_message").html(message);
			$("#solaris_form_person_valid_message").show();		
		} else {
			
			$("#solaris_form_person_valid_message").html("");
			$("#solaris_form_person_valid_message").hide();	
			
			
			$("#" + row_id +  " td.req_value").html(v_inn);
			$("#" + row_id).show();
		}

	}
	
	
	if (row_id == "req_person_birthdate") {
		
		var v_day = document.getElementById("solaris_birthdate_day").value;		
		var v_month = document.getElementById("solaris_birthdate_month").value;		
		var v_year = document.getElementById("solaris_birthdate_year").value;		
		
		v_day = v_day.trim();
		v_month = v_month.trim();
		v_year = v_year.trim();
				
		if(v_day == "" && v_month == "" && v_year == "")
		{
			message = message + error_index + ". Пустое значение даты рождения</br>";
			error_index++;
			result = false;
		}
		
		if(v_day != "" && (v_day <= 0 || v_day > 31 || !(/^[0-9]+$/.test(v_day)) ) )
		{
			message = message + error_index + ". Некорректный день рождения</br>";
			error_index = error_index + 1;
			result = false;
		}		
		
		if(v_month != "" && (v_month <= 0 || v_month > 12 || !(/^[0-9]+$/.test(v_month)) ))
		{
			message = message + error_index + ". Некорректный месяц рождения</br>";
			error_index = error_index + 1;
			result = false;
		}
		
		
		
		if(v_year != "" && (v_year <= 1900 || v_year > 2019 || !(/^[0-9]+$/.test(v_year)) ))
		{
			message = message + error_index + ". Некорректный год рождения</br>";
			error_index = error_index + 1;
			result = false;
		}
		
		
		if(v_day != "" && v_month != "" && v_year != "")
		{
			var test_date = "";
					
			if(v_day.length == 1) {
				test_date += "0" + v_day + ".";
			} else {
				test_date += v_day + ".";
			}
			
			if(v_month.length == 1) {
				test_date += "0" + v_month + ".";
			} else {
				test_date += v_month + ".";
			}
			
			test_date += v_year;
			
			
			if(!isValidDate(test_date)){
				message = message + error_index + ". Такой даты не существует</br>";
				error_index++;
				result = false;
			}
		}
		
		
		if(result == false)	{
			$("#solaris_form_person_valid_message").html(message);
			$("#solaris_form_person_valid_message").show();		
		} else {
			
			$("#solaris_form_person_valid_message").html("");
			$("#solaris_form_person_valid_message").hide();	
			
			
			var req_bd = "";
			
			if( v_day != "" ){
				
				if(v_day.length == 1) {
					req_bd += "0" + v_day + ".";
				} else {
					req_bd += v_day + ".";
				}
			}
			else
				req_bd += "__" + ".";
			
			
			
			
			if( v_month != "" ){
				if(v_month.length == 1) {
					req_bd += "0" + v_month + ".";
				} else {
					req_bd += v_month + ".";
				}
					
			}
			else
				req_bd += "__" + ".";
			
			if( v_year != "" )
				req_bd += v_year + " г.р.";
			else
				req_bd += "____" + " г.р.";
			
			
			$("#" + row_id +  " td.req_value").html(req_bd);
			$("#" + row_id).show();
		}

	}
	
	if (row_id == "all") {
		
		var dateerr = true;
		
		var v_surname = document.getElementById("solaris_surname").value;
		var v_name = document.getElementById("solaris_name").value;		
		var v_fathername = document.getElementById("solaris_fathername").value;		
				
		var v_day = document.getElementById("solaris_birthdate_day").value;		
		var v_month = document.getElementById("solaris_birthdate_month").value;		
		var v_year = document.getElementById("solaris_birthdate_year").value;	

		var v_inn = document.getElementById("solaris_inn").value;	
		
		v_surname = v_surname.trim();
		v_name = v_name.trim();
		v_fathername = v_fathername.trim();
		
		v_day = v_day.trim();
		v_month = v_month.trim();
		v_year = v_year.trim();
		
		v_inn = v_inn.trim();
		
		//alert(v_surname + ' - ' + v_name + ' - ' + v_fathername + ' - ' + v_day + ' - ' + v_month + ' - ' + v_year);		
		// проверка Фамилии
				
		if(v_surname != "") {

			if(!(/^[а-яёА-ЯЁ\-]+$/.test(v_surname)))
			{				
				message = message + error_index + ". Недопустимые символы в фамилии</br>";
				error_index = error_index + 1;
				result = false;
			} else {			
				$("#req_person_surname td.req_value").html(v_surname);
				$("#req_person_surname").show();
			}			
		}
		
		// проверка Имени
		
		if(v_name != "") {
			
			if(!(/^[а-яёА-ЯЁ]+$/.test(v_name)) ){
				message = message + error_index + ". Недопустимые символы в имени</br>";
				error_index = error_index + 1;
				result = false;
			} else {
				
				$("#req_person_name td.req_value").html(v_name);
				$("#req_person_name").show();
			}	
		}
		
		// проверка Отчества
		
		if(v_fathername != "") {

			if (!(/^[а-яёА-ЯЁ]+$/.test(v_fathername)) )
			{
				message = message + error_index + ". Недопустимые символы в отчестве</br>";
				error_index = error_index + 1;
				result = false;
			} else {
				
				$("#req_person_fathername td.req_value").html(v_fathername);
				$("#req_person_fathername").show();
			}		
		}
				
		
		// проверка Даты рождения
				
		var req_date = ""

		if(v_day != "") {
			if ( (v_day <= 0 || v_day > 31 || !(/^[0-9]+$/.test(v_day)) ) )
			{
				message = message + error_index + ". Некорректный день рождения</br>";
				error_index = error_index + 1;
				req_date = "__.";
				dateerr = false;
				result = false;
				
			} else {
				
				if(v_day.length == 1) {
					req_date += "0" + v_day + ".";
				} else {
					req_date += v_day + ".";
				}			
			}
		} else {
			req_date = "__.";
		}
		
		if(v_month != "") {
			
			if ((v_month <= 0 || v_month > 12 || !(/^[0-9]+$/.test(v_month)) ))
			{
				message = message + error_index + ". Некорректный месяц рождения</br>";
				error_index = error_index + 1;
				req_date += "__.";
				dateerr = false;
				result = false;
			} else {
				
				if(v_month.length == 1) {
					req_date += "0" + v_month + ".";
				} else {
					req_date += v_month + ".";
				}			
			}
		} else {
			req_date += "__.";
		}
				
		if(v_year != "") {
			
			if  (v_year <= 1900 || v_year > 2019 || !(/^[0-9]+$/.test(v_year)) )
			{
				message = message + error_index + ". Некорректный год рождения</br>";
				error_index = error_index + 1;
				req_date += "____ г.р.";
				dateerr = false;
				result = false;
			} else {
				req_date += v_year + " г.р.";			
			}
		} else {
			req_date += "____ г.р.";
		}
		
		
		if(v_day != "" || v_month != "" || v_year != "" ){
			if (req_date != "__.__.____ г.р.") {
				$("#req_person_birthdate td.req_value").html(req_date);
				$("#req_person_birthdate").show();
			}
		}
		
			
		// ИНН

		if(v_inn != "") {

			if(!(/\d{12}/.test(v_inn)))
			{				
				message = message + error_index + ". Некорректный ИНН ФЛ</br>";
				error_index = error_index + 1;
				result = false;
			} else {			
				$("#req_person_inn td.req_value").html(v_inn);
				$("#req_person_inn").show();
			}			
		}
			
		if(result == false)	{
			$("#solaris_form_person_valid_message").html(message);
			$("#solaris_form_person_valid_message").show();		
		} 
	}
	
		
	solaris_CheckRequestParametrField();
	
}


function CheckForFullString(){
	
	var str = "";	
	
	

    setTimeout(function(e) {
        
		str = document.getElementById("solaris_surname").value;	
		str = str.trim();
		var words_count = 0	
		
		words_count = str.split(" ");		
		
		if(words_count.length == 4){
			$("#solaris_surname").val(words_count[0]);
			$("#solaris_name").val(words_count[1]);		
			$("#solaris_fathername").val(words_count[2]);
			
			if(isValidDate(words_count[3]))
			{
				dt = words_count[3].split(".");
				$("#solaris_birthdate_day").val(dt[0]);
				$("#solaris_birthdate_month").val(dt[1]);
				$("#solaris_birthdate_year").val(dt[2]);
				
			}
		}
		
		if(words_count.length == 3){
			$("#solaris_surname").val(words_count[0]);
			$("#solaris_name").val(words_count[1]);
			$("#solaris_fathername").val(words_count[2]);
		}
		
		if(words_count.length == 2){
			$("#solaris_surname").val(words_count[0]);
			$("#solaris_name").val(words_count[1]);		
			
		}
		
	  
	
    }, 0);

		  
		  
	
}







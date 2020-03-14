

var KadrResults


jQuery(document).ready(function() 
{
	
	
	
	
	$('#myTab a:first').tab('show')

	$("#form_person_valid_message").hide();
	$("#form_vehicle_valid_message").hide();
	$("#form_company_valid_message").hide();
	
	
	
	$("#kadr_person form").submit(function(event){

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
		  url: "/kadr/person",
		  data: formData,			  
		  success: function(data){							
			PrintKadrPersonResultTable(data);			
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


function PrintKadrPersonResultTable (data){
		
	$("#ResultTable").text("");	
	//var Results = $.parseJSON(data);
	
	KadrResults = $.parseJSON(data);
	
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
	render += "		   <td>Дата проверки</td>";
	render += "		</tr>";
	render += "	</thead>";
	render += "	<tbody>";
	
	var index = 1
	
	var newdate
	var dateparts	
	var init_element
	
	for (var part in KadrResults) {
		//alert(Results[part]);
		if (KadrResults[part]["result"] == 0){
			render += "		<tr class=\"otkaz\">";
		} else {
			render += "		<tr>";
		}
		render += "		   <td >" + index + "</td>";
		render += "		   <td><span id=\"kadrperson_" + index + "\" onClick=\"PrintKadrDetails(this,'no');\">" + 	
									capitalizeFirstLetter(KadrResults[part]["sname"].trim()) + " " + 
									capitalizeFirstLetter(KadrResults[part]["name"].trim()) + " " + 
									capitalizeFirstLetter(KadrResults[part]["fname"].trim()) + "</span></td>";
		
		if (index == 1)
			PrintKadrDetails(null,"kadrperson_" + index);
		
		render += "		   <td> " + KadrResults[part]["birthdate"] + " </td>";
		render += "		   <td> " + KadrResults[part]["checked_date"] + " </td>";
		
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


function PrintKadrDetails(element, init_element){
		
	
	
	
	var id
	var type
	
	if(init_element == "no"){
	
		id = element.id.split("_")[1];
		type = element.id.split("_")[0];
	} else {
		id = init_element.split("_")[1];
		type = init_element.split("_")[0];
	}
	
	//alert(type + "-" + id);
	if(type == "kadrperson"){	
		//alert(KadrResults);
		for (var part in KadrResults) {
			//alert(part + "-" + "r" + id);
			if(part == "r" + id){
				
				$("#res_fio").text(
					capitalizeFirstLetter(KadrResults[part]["sname"].trim()) + " " + 
					capitalizeFirstLetter(KadrResults[part]["name"].trim()) + " " + 
					capitalizeFirstLetter(KadrResults[part]["fname"].trim()));
				
		
				$("#res_dr").text(KadrResults[part]["birthdate"]);				
			/*
				var reason = Results[part].reason.replace(/\[#NL\]/g,"</br>") 
				reason = reason.replace(/\[#Q\]/g,"\"") 
*/				
				$("#res_source").text("Не указан");
				if( KadrResults[part].source  == "ACCESS"){
					$("#res_source").text("Журнал ЦПМ");
				}
				if( KadrResults[part].source  == "RD17"){
					$("#res_source").text("RD17");
				}
				
				$("#res_result").text("Не указано");				
				if( KadrResults[part].result  == 0){
					$("#res_result").text("Отказ");
				}
				if( KadrResults[part].result  == 1){
					$("#res_result").text("Согласован");
				}
				
				$("#res_guarantee").text("Не указано");
				if( KadrResults[part].guarantee  == 0){
					$("#res_guarantee").text("Нет");
				}
				if( KadrResults[part].guarantee  == 1){
					$("#res_guarantee").text("Да");
				}				
				
				
				$("#res_comment").text(KadrResults[part].comment);
				$("#res_checked_by").text(KadrResults[part].checked_by);
				$("#res_checked_data").text(KadrResults[part].checked_date);
			
				break;
			}
		}
	}
	
	
	
}


function ValidKadrPerson(data){
	
	
	$("#form_person_valid_message").hide();
	var result = true;	
	
	var error_index = 1;
	var message  = "";
	
	var surname = document.getElementById("surname").value;
	var name = document.getElementById("name").value;
	var fathername = document.getElementById("fathername").value;
	var birthdate = document.getElementById("birthdate").value;
		
	
	
	if(surname == "" && (name == "" || fathername == "" || birthdate == ""))
	{
		message = message + error_index + ". Фамилия обязательна для ввода или заполнить полностью все остальные поля</br>";
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


 



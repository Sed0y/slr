

jQuery(document).ready(function() 
{
	//'yyyy-mm-dd'
	$('#begin-data-input').dcalendarpicker({theme: 'green', format: 'dd-mm-yyyy' });
	$('#end-data-input').dcalendarpicker({theme: 'green', format: 'dd-mm-yyyy' });
		
	$('#uvk_money-report #wrap_body table tr td.click').on('click', function(){		
		
		
		var debtor_type = $(this).attr('id');
		
		
		var	p_filial_id = $('#filials').val();

		var formData = {
						id:p_filial_id,												
						begin:FormatDateTo_ymd($('#begin-data-input').val()),
						end:FormatDateTo_ymd($('#end-data-input').val()),
						debtor_type:debtor_type,
						curator: $('#curator').prop('checked'),
					}; 
						
		$.ajaxSetup({cache: false}); 
		
		$.ajax({
		  type: "POST",
		  url: "/data/uvk/money/details/filial",
		  data: formData,			  
		  success: function(data){
			//alert(data);
			PrintUvkDetailData(data);
		  }
		});		
	});
	
	
	$('#calculate_uvk_money').on('click', function() {
			
		//alert("click");		return;
		
		$("#Uvk_Money_Details").text("");
		
		var formData = {						
					id: $('#filials').val(),
					curator: $('#curator').prop('checked'),
					begin: FormatDateTo_ymd($('#begin-data-input').val()),
					end: FormatDateTo_ymd($('#end-data-input').val())
				}; 
		
		$.ajaxSetup({cache: false}); 
		$.ajax({
		  type: "POST",
		  url: "/data/uvk/money/filial",
		  data: formData,			  
		  success: function(data){	
			//alert(data);
			PrintUvkReportData(data);
		  }
		});
			
	});
	
	
});


function PrintUvkReportData(data){


	var Results = $.parseJSON(data);		

	var all_money = 0
	for (var InsType in Results) {
		
		for (var Parameter in Results[InsType]) {
		
			if( Parameter == "Financial"){					
				$('#' + InsType).text(ToMoney(Results[InsType][Parameter])+"");
				all_money = all_money + 1*Results[InsType][Parameter];
			}
		}		
	}	

	$('#UVK-All').text(ToMoney(all_money)+"");
	

}


function PrintUvkDetailData(data){
		
	$("#Uvk_Money_Details").text("");
	
	//alert("parseJSON");
	//$("#Uvk_Money_Details").text(data);
	var Results = $.parseJSON(data);
	
	
	var render = "";
	var fin_column = -1;
	var id_column = -1;
	
	render += "<table id=\"DetailsTable\" style=\"width:100%\" class=\"table table-bordered compact\" cellspacing=\"0\">";
	render += "	<thead>";
	render += "		<tr>";
	
	for (var part in Results["header"]) {
		if (Results["header"][part] == "finanse" || Results["header"][part] == "money"){
			render += "			<td type=\"money\">" + Translate(Results["header"][part]) + "</td>";
			fin_column = part*1;
		}
		else if (Results["header"][part] == "id"){		
			id_column = part*1;
		}else{		
			render += "			<td>" + Translate(Results["header"][part]) + "</td>";
		}
	}
	
	render += "  </tr>";
	render += "</thead>";	
	render += "<tbody>";
	
	//alert("print_body");
	
	var index
	for (var part in Results["data"]) {			
		render += "<tr>";
		
		index = 0
		for (var val in Results["data"][part]) {
			
			if(index == id_column){
				index++;
				continue;
			}
			
			if (Results["data"][part][val].indexOf("T00:00:00Z") !== -1) {
				res = Results["data"][part][val].replace("T00:00:00Z","");
				
				render += "<td>" + IhDateToNashDate(res) + "</td>";		
				
			} else if(Results["data"][part][val].indexOf("Зона ф-ла") !== -1) {
				res = Results["data"][part][val].replace("Зона ф-ла","");
				render += "<td>" + res + "</td>";
			} else {
				render += "<td>" + Results["data"][part][val] + "</td>";
			}				
			index++;
		}	
		render += "</tr>";
		
	}
	render += "</tbody>";	
	render += "</table>";
	
	if (id_column != -1){
		fin_column = fin_column -1;
	}
	
	$("#Uvk_Money_Details").append(render);
	
	list= document.getElementById("DetailsTable").getElementsByTagName("tr"); 

	for (var i=1; i< list.length; ++i) 
	{ 
		list[i].getElementsByTagName("td")[fin_column].innerHTML= ToMoney2(list[i].getElementsByTagName("td")[fin_column ].innerHTML);  
	}; 

	
	$('#DetailsTable').DataTable({
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
	
	//alert("here");
	
}	









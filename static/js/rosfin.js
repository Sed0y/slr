
var ClickedObject


jQuery(document).ready(function() 
{
	$('#rosfin-actual-data').dcalendarpicker({theme: 'green', format: 'yyyy-mm-dd' });
		
		
	$( "#rosfin_progressbar" ).progressbar({
      value: 1,
	  change: function() {		
        $( "#rosfin_progressbar .progress-label" ).text( $( "#rosfin_progressbar" ).progressbar( "value" ) + "%");
      },
	  complete: function() {
        $( "#rosfin_progressbar .progress-label" ).text( "Готово" );
		ClickedObject.css('color','#444');
		setTimeout( PB_hide, 1800 );
		
		var attributes = ClickedObject.attr('id').split('-');
				
		$( "#LoadedDate" + attributes[2] ).text(GetNowDate());
		
		var link = document.createElement('a');			
		var elem = document.getElementById("download-link");
		
		elem.appendChild(link);
		link.href = "/static/files/rosfin/results/" + $("#download-link").text() + ".zip";
		setTimeout( function () { link.click(); }, 2300 );
		
		
					
		
      }
    });
	
	$("#rosfin_progressbar").hide();
	$("#rosfin_progress_console").hide();	
		
	
	$('#RosfinFiles').DataTable({
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
			}            
        }        		
	});
	
  	
	
	$('#RosfinFiles .Update').on('click', function(){

		Progress_Count = -1
		ClickedObject = $(this)
		ClickedObject.css('color','red');
		
		
		var attributes = ClickedObject.attr('id').split('-');
		
		var formData = 	{
							id:attributes[2]																			
						}; 
		
		$.ajaxSetup({cache: false}); 
		
		$.ajax({
		  type: "POST",
		  url: "/admin/rosfin/parse",
		  data: formData,			  
		  success: function(data){				
			
			answ = data.split('#');			
			
			if(answ[0] == "success"){
				$( "#rosfin_progressbar .progress-label" ).text( "Обработка файла..." );
				$( "#rosfin_progressbar" ).fadeIn( "slow" );	
				$( "#rosfin_progress_console" ).fadeIn( "slow" );
				
				$("#rosfin_progressbar" ).progressbar( "value", 3 );
				$("#download-link").text(answ[1]);				
				
				setTimeout(function() {object.func(attributes[2])}, 500);
				
				
			} else {
				alert(data);
			}
		
		  }
		});
		
		return
		
		
	});
		
		
	
	
	$('#RosfinFiles .Delete').on('click', function(){

		
		ClickedObject = $(this)
		ClickedObject.css('color','red');
		
		
		var attributes = ClickedObject.attr('id').split('-');
		
		var formData = 	{
							id:attributes[2]																			
						}; 
		
		$.ajaxSetup({cache: false}); 
		
		$.ajax({
		  type: "POST",
		  url: "/admin/rosfin/delete",
		  data: formData,			  
		  success: function(data){				
			
			location.reload();
		
		  }
		});
		
		return
		
		
	});
		
				
	
});


object = {  
   func:  function CheckProgress(id_value){
							
			var formData = 	{
								id:id_value																			
							}; 
								
			$.ajax({
				  type: "POST",
				  url: "/admin/rosfin/parse/progress",
				  data: formData,			  
				  success: function(data){				
					
					p = data.split('#');
					
					if (p[0]*1 == 100) {
						$("#rosfin_progressbar" ).progressbar( "value", 100 );						
					} else {							
						PB_set(p[0]*1, p[1]);
						
						
						setTimeout (function() {object.func(id_value)},500);							
					}
					
				
				  }
				});
		
		}
}
	
	

function PB_hide() {
	$("#rosfin_progressbar" ).fadeOut( "slow" );
	
}

function PB_set(progress_value, progress_count) {
	
	var val = $("#rosfin_progressbar" ).progressbar( "value" ) || 0;
	if ( progress_value < 99 ) {	
		$("#rosfin_progressbar" ).progressbar( "value", progress_value );
	}
	
	$("#progress_count").text(progress_count);

	
}	
	
function GetNowDate() {
		
	var now = new Date();
	
	var year = now.getFullYear();
	var month = now.getMonth() + 1;
	var day = now.getDate();
	
	if ( month < 10){
		month = "0" + month;
	}
			
	return day + "." + month + "." + year;
		
}

function validate_upload_form() {
	
	var date
	var description
	
	date = document.upload_form.actual.value
	description = document.upload_form.description.value
	
	if(date == "") {
		alert ("Поле \"Актуальность\" обязательно для заполнения");
		return false;
	}

	d = date.split('-');
	
	if (d.length != 3) {
		// Тут по-хоршему надо делать полную проверку, но будем считать, что заполняться будет только через плагин
		// главное чтоб не пустое
		alert ("Неверный формат даты");
		return false;
	}
	
	
	return true;
}









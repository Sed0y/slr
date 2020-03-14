

jQuery(document).ready(function() 
{
	
	$('#begin-data-input').dcalendarpicker({theme: 'green', format: 'dd-mm-yyyy' });
	$('#end-data-input').dcalendarpicker({theme: 'green', format: 'dd-mm-yyyy' });
	
	
	$('#calculate_filial').on('click', function() {
				
		var formData = {						
					id: $('#filials').val(),
					curator: $('#curator').prop('checked'),
					begin: FormatDateTo_ymd($('#begin-data-input').val()),
					end: FormatDateTo_ymd($('#end-data-input').val())
				}; 
		
		$.ajaxSetup({cache: false}); 
		$.ajax({
		  type: "POST",
		  url: "/data/upma/reguliration/filial",
		  data: formData,			  
		  success: function(data){					
			PrintUpmaReportData(data);
			
		  }
		});
			
		
		var formData = {						
					id: $('#filials').val(),
					curator: $('#curator').prop('checked'),
					begin: FormatDateTo_ymd($('#begin-data-input').val()),
					end: FormatDateTo_ymd($('#end-data-input').val())
				}; 
		
		$.ajaxSetup({cache: false}); 
		$.ajax({
		  type: "POST",
		  url: "/data/uiivs/reguliration/filial",
		  data: formData,			  
		  success: function(data){	
			//alert(data);
			PrintUiivsReportData(data);
		  }
		});
		
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
	
	$('#download_excel').on('click', function() {
		
		var formData = {						
					id: $('#filials').val(),
					curator: $('#curator').prop('checked'),
					begin: FormatDateTo_ymd($('#begin-data-input').val()),
					end: FormatDateTo_ymd($('#end-data-input').val())
				}; 
			
		
		$.ajaxSetup({cache: false}); 
		$.ajax({
		  type: "POST",
		  url: "/report/filial/excel",
		  data: formData,			  
		  success: function(data){	
			
			answ = data.split('#');			
			
			if(answ[0] == "success"){
				
				var link = document.createElement('a');			
				var elem = document.getElementById("download-link");
		
				elem.appendChild(link);
				link.href = answ[1];
				link.click(); 
			}
		  }
		});

		
		
		
	});

	
});











  <!-- Navigation-->
  <nav class="navbar navbar-expand-lg navbar-dark bg-dark fixed-top" id="mainNav">
    <a class="navbar-brand" href="#">Solaris</a>
    <button class="navbar-toggler navbar-toggler-right" type="button" data-toggle="collapse" data-target="#navbarResponsive" aria-controls="navbarResponsive" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarResponsive">
      <ul class="navbar-nav navbar-sidenav" id="exampleAccordion">
		
		{{if .Root -}}
		
		{{- else}}
				
		{{- end}}
        
		{{if or .Managment .Root -}}
        
		{{- else}}
				
		{{- end}}
       
		
		<li class="nav-item" data-toggle="tooltip" data-placement="right" title="Checker">
          <a class="nav-link nav-link-collapse collapsed" data-toggle="collapse" href="#collapseChecker" data-parent="#reportAccordion">
            <i class="fa fa-fw fa-search"></i>
            <span class="nav-link-text">Проверка</span>
          </a>
          <ul class="sidenav-second-level collapse" id="collapseChecker">		  
			        		
			<li>
              <a href="/check/solaris">Солярис</a>
            </li>
			
          </ul>
        </li>   
		
		
<!--     
        <li class="nav-item" data-toggle="tooltip" data-placement="right" title="Link">
          <a class="nav-link" href="/links">
            <i class="fa fa-fw fa-link"></i>
            <span class="nav-link-text">Ссылки</span>
          </a>
        </li>
-->
      </ul>
      <ul class="navbar-nav sidenav-toggler">
        <li class="nav-item">
          <a class="nav-link text-center" id="sidenavToggler">
            <i class="fa fa-fw fa-angle-left"></i>
          </a>
        </li>
      </ul>
      <ul class="navbar-nav ml-auto">        
        <li class="nav-item">
          <a class="nav-link" data-toggle="modal" data-target="#exampleModal">
            <i class="fa fa-fw fa-sign-out"></i>Выход</a>
        </li>
      </ul>
    </div>
  </nav>
  
  
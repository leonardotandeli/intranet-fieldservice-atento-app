//var nome = "{{.Nome}}"

//document.getElementById("navbarDarkDropdownMenuLink").innerHTML = nome.substring(0,8);
//document.getElementById("nome_pagina").innerHTML = "{{ .Pagina }}";

jQuery(function($) {
    var path = window.location.href; 
    // because the 'href' property of the DOM element is the absolute path
    $('ul a').each(function() {
      if (this.href === path) {
        $(this).addClass('active');
      }
    });
  });
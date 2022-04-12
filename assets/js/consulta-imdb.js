$('#form-imdb').on('submit', consultaIMDB)


function consultaIMDB(evento) {
    evento.preventDefault();


    var re = $("#re").val();
    
    //var login = $("#login").val();
           $("#diframeHolder").removeClass("esconde");
        $("#iframeHolder").attr("src","http://reportcorp.atento.com.br/ReportServer/Pages/ReportViewer.aspx?%2fIMDB%2frptUserGrupoProxy&COD_RE="+re+"&rc:Area=Toolbar&rc:LinkTarget=frmMain&rc:JavaScript=True&rs:ClearSession=true&rc:Parameters=false");
     
    }




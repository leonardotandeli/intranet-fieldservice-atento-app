$('#form-catraca').on('submit', consultaCatraca)


function consultaCatraca(evento) {
    evento.preventDefault();


    var re = $("#re").val();
    var data = $("#data").val();
           $("#diframeHolder").removeClass("esconde");
        $("#iframeHolder").attr("src","http://reportcorp.atento.com.br/ReportServer/Pages/ReportViewer.aspx?%2fCATRACA%2fLogAcessoCatracaFuncionario&MATRICULAS="+re+"&DTMARCACAO="+data+"&rc:Area=Toolbar&rc:LinkTarget=frmMain&rc:JavaScript=True&rs:ClearSession=true&rc:Parameters=false");
     
    }




$('#form-ativo').on('submit', consultaAtivo)


function consultaAtivo(evento) {
    evento.preventDefault();
    var serie = $("#serie").val();

    var locador = $("#locador").val();
    console.log(locador)
    
           $("#diframeHolder").removeClass("esconde");
        $("#iframeHolder").attr("src","http://reportcorp.atento.com.br/ReportServer/Pages/ReportViewer.aspx?/GestaoSuprimentos/rpt_equipamento&NOM_EQUIPAMENTO="+locador+"&NOM_EQUIPAMENTO_ANTIGO=&NUM_ATIVO_SAP=&NOM_NUMERO_SERIE="+serie+"&NOM_MARCA=&NOM_MODELO=&COD_GRUPO=&COD_TIPO_EQUIPAMENTO=&COD_SITUACAO_EQUIPAMENTO=&NOM_FORNECEDOR=&NOM_PEDIDO_COMPRA=&NOM_DOC_AQUISICAO=&NUM_CHAMADO_CAD=&NUM_CHAMADO_REG=&NUM_RE_FUNCIONARIO=&NUM_CONTRATO_LEASING=&COD_ESTABELECIMENTO=&COD_CATEGORIA=&COD_CLASSIF_CONTABIL=&DES_EQUIPAMENTO=&DAT_INVENTARIO_INI=&DAT_INVENTARIO_FIM=&DAT_CADASTRO_INI=&DAT_CADASTRO_FIM=&DAT_AQUISICAO_INI=&DAT_AQUISICAO_FIM=&DAT_TSP_ATUALIZACAO_INI=&DAT_TSP_ATUALIZACAO_FIM=&rc:Area=Toolbar&rc:LinkTarget=frmMain&rc:JavaScript=True&rs:ClearSession=true&rc:Parameters=false");
     
    }




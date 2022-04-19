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


$('#login').on('submit', fazerLogin);

function fazerLogin(evento) {

    //previne o comportamento padrão do formulario
    evento.preventDefault();

    //
    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            login_nt: $('#login_nt').val(),
            senha: $('#senha').val()

        }
    }).done(function() {
     window.location = "/home"
    }).fail(function(){

        $('#alertd').fadeIn(1000);
        $('#alertd').text("Algo deu errado! O login ou a senha que você inseriu não estão corretos.");

        })
}


$('#formulario-impressora').on('submit', criarImpressora);
$('#atualizar-impressora').on('click', atualizarImpressora);
$('#deletar-impressora').on('click', deletarImpressora);

function criarImpressora(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/impressoras/adicionar",
        method: "POST",
        data: {
            local: $('#local').val(),
            servidor: $('#servidor').val(),
            ip: $('#ip').val(),
            ip_switch: $('#ip_switch').val(),
            porta_switch: $('#porta_switch').val(),
            fila: $('#fila').val(),
            modelo: $('#modelo').val(),
            serialnumber: $('#serialnumber').val(),
            status: $('#status').val()
        }
    }).done(function() {
        Swal.fire({
            type: 'success',
            title: 'Solicitação criada com sucesso!',
            showConfirmButton: false,
            timer: 1500
          }).then(function() {window.location = '/impressoras';})
     
    }).fail(function(erro){
        console.log(erro)
        Swal.fire({
            type: 'error',
            title: 'Algo deu errado! Não foi possível enviar a solicitação.',
            showConfirmButton: false,
            timer: 1500
          })
    })
}



function atualizarImpressora(evento) {
    evento.preventDefault();
    const impressoraId = $(this).data('impressora-id');
    $.ajax({
        url: '/impressoras/'+impressoraId,
        method: "PUT",
        data: {
            local: $('#local').val(),
            servidor: $('#servidor').val(),
            ip: $('#ip').val(),
            ip_switch: $('#ip_switch').val(),
            porta_switch: $('#porta_switch').val(),
            fila: $('#fila').val(),
            modelo: $('#modelo').val(),
            serialnumber: $('#serialnumber').val(),
            status: $('#status').val()
        }
    }).done(function() {

        Swal.fire({
            type: 'success',
            title: 'Publicação atualizada com sucesso!!',
            showConfirmButton: false,
            timer: 1500
          }).then(function() {window.location = '/impressoras/'+publicacaoId;})
     
    }).fail(function(erro){
        console.log(erro)
        Swal.fire({
            type: 'error',
            title: 'Algo deu errado! Não foi possível atualizar a publicação.',
            showConfirmButton: false,
            timer: 1500
          })
    })
}

function deletarImpressora(evento) {
    evento.preventDefault();
    const impressoraId = $(this).data('impressora-id');
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir esse usuário? Essa ação é irreversível!",
        cancelButtonText: "Cancelar",
        showCancelButton: true,
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;

    
        $.ajax({
            url: '/impressoras/'+impressoraId,
            method: "DELETE"
        }).done(function() {
            window.location = '/impressoras';
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao excluir o usuário!", "error");
        });
    })
}

$('#formulario').on('submit', criarChamado);
$('#formulario-mapa').on('submit', criarMapa);
$('#deletar-mapa').on('click', DeletarMapa);
$('#atualizar-chamado').on('click', atualizarChamado);
$('#atualizar-mapa').on('click', atualizarMapa);
$('#atualizar-categoria').on('click', atualizarCategoria);
$('#criar-usuario-usuario').on('click', criarUsuarioChamado);
$('#criar-categoria').on('submit', criarCategoria);
function criarChamado(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/formulario",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            chamado: $('#chamado').val(),
            ativocpu: $('#ativocpu').val(),
            ativomonitor: $('#ativomonitor').val(),
            endereco: $('#endereco').val(),
            numero: $('#numero').val(),
            cep: $('#cep').val(),
            senha: $('#senha').val(),
            transporte: $('#transporte').val(),
            acionamento: $('#acionamento').val(),
            status: $('#status').val(),
            bairro: $('#bairro').val(),
            obs: $('#obs').val(),
            office: $('#office').val(),      
            ramal: $('#ramal').val(),
            logindac: $('#logindac').val(),
            re: $('#re').val(),
            ativoretornomonitor: $('#ativoretornomonitor').val(),
            ativoretornocpu: $('#ativoretornocpu').val(),
            periferico_mouse: $('#periferico_mouse').val(),
            periferico_teclado: $('#periferico_teclado').val(),
            periferico_head: $('#periferico_head').val(),
            periferico_rede: $('#periferico_rede').val(),
            analistafield: $('#analistafield').val(),
            gerenteoperador: $('#gerenteoperador').val(),
            asite: $('#asite').val()
        }
    }).done(function() {
        Swal.fire({
            type: 'success',
            title: 'Solicitação criada com sucesso!',
            showConfirmButton: false,
            timer: 1500
          }).then(function() {window.location = '/chamados';})
     
    }).fail(function(erro){
        console.log(erro)
        Swal.fire({
            type: 'error',
            title: 'Algo deu errado! Não foi possível enviar a solicitação.',
            showConfirmButton: false,
            timer: 1500
          })
    })
}



function DeletarMapa(evento) {
    evento.preventDefault();
    const mapaId = $(this).data('mapa-id');
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa operação? Essa ação é irreversível!",
        cancelButtonText: "Cancelar",
        showCancelButton: true,
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;

    
        $.ajax({
            url: '/mapa/operacoes/'+mapaId,
            method: "DELETE"
        }).done(function() {
            window.location = '/mapa/operacoes';
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao excluir a operação!", "error");
        });
    })
}


function criarMapa(evento) {
    evento.preventDefault();
  
  var textareaText = $('#template').val();
  textareaText = textareaText.replace(/\r?\n/g, '<br />');



    $.ajax({
        url: "/formulario/mapa",
        method: "POST",
        data: {
            operacao: $('#operacao').val(),
            vlan_dados: $('#vlan_dados').val(),
            vlan_voz: $('#vlan_voz').val(),
            config_contratual: $('#config_contratual').val(),
            versao_windows: $('#versao_windows').val(),
            imagem: $('#imagem').val(),
            template: textareaText,
            grupo_imdb: $('#grupo_imdb').val(),
            gravador: $('#gravador').val(),
            observacoes: $('#observacoes').val(),
            id_site: $('#id_site').val(),
            id_cliente: $('#id_cliente').val(),
            id_dac: $('#id_dac').val(),
            id_dominio: $('#id_dominio').val()
        }
    }).done(function() {
        Swal.fire({
            type: 'success',
            title: 'Operação adicionada com sucesso!',
            showConfirmButton: false,
            timer: 1500
          }).then(function() {window.location = '/mapa/operacoes';})
     
    }).fail(function(erro){
        console.log(erro)
        Swal.fire({
            type: 'error',
            title: 'Algo deu errado! Verifique se todos os campos foram preenchidos.',
            showConfirmButton: false,
            timer: 1500
          })
    })
}






function atualizarMapa() {
    $(this).prop('disabled', true);

    const mapaId = $(this).data('mapa-id');
    console.log(mapaId)
    var textareaText = $('#template').val();
    textareaText = textareaText.replace(/\r?\n/g, '<br />');
    
    $.ajax({
        url: '/mapa/operacoes/'+mapaId,
        method: "PUT",
        data: {
            operacao: $('#operacao').val(),
            vlan_dados: $('#vlan_dados').val(),
            vlan_voz: $('#vlan_voz').val(),
            config_contratual: $('#config_contratual').val(),
            versao_windows: $('#versao_windows').val(),
            imagem: $('#imagem').val(),
            template: textareaText,
            grupo_imdb: $('#grupo_imdb').val(),
            gravador: $('#gravador').val(),
            observacoes: $('#observacoes').val(),
            id_site: $('#id_site').val(),
            id_cliente: $('#id_cliente').val(),
            id_dac: $('#id_dac').val(),
            id_dominio: $('#id_dominio').val()
        }
    }).done(function() {
        Swal.fire({
            type: 'success',
            title: 'Informações atualizadas com sucesso!',
            showConfirmButton: false,
            timer: 2000
          }).then(function() {$('#atualizar-chamado').prop('disabled', false); window.location = '/mapa/operacoes';})
    }).fail(function() {
        Swal.fire({
            type: 'error',
            title: 'Algo deu errado! Não foi possível enviar a solicitação.',
            showConfirmButton: false,
            timer: 1500
          })
    })
}





function atualizarCategoria() {
    $(this).prop('disabled', true);

    const catId = $(this).data('cat-id');
    console.log(catId)
  
    
    $.ajax({
        url: '/base/editar/categoria/'+catId,
        method: "PUT",
        data: {
            nome: $('#nome').val(),
        }
    }).done(function() {
        Swal.fire({
            type: 'success',
            title: 'Informações atualizadas com sucesso!',
            showConfirmButton: false,
            timer: 2000
          }).then(function() {window.location = '/base';})
    }).fail(function() {
        Swal.fire({
            type: 'error',
            title: 'Algo deu errado! Não foi possível enviar a solicitação.',
            showConfirmButton: false,
            timer: 1500
          })
    }).always(function() {
        $('#atualizar-categoria').prop('disabled', false);
    })
}





function criarUsuarioChamado(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/formulario/usuario",
        method: "POST",
        data: {
            NOME: $('#nome').val(),
            ENDERECO: $('#chamado').val(),
            NUMERO: $('#ativocpu').val(),
            CEP: $('#ativomonitor').val(),
            BAIRRO: $('#endereco').val(),
            RE: $('#numero').val(),
            ID_GERENTE: $('#id_gerente').val()
        }
    }).done(function() {
        Swal.fire({
            type: 'success',
            title: 'Solicitação criada com sucesso!',
            showConfirmButton: false,
            timer: 1500
          }).then(function() {window.location = '/chamados';})
     
    }).fail(function(erro){
        console.log(erro)
        Swal.fire({
            type: 'error',
            title: 'Algo deu errado! Não foi possível enviar a solicitação.',
            showConfirmButton: false,
            timer: 1500
          })
    })
}

function criarCategoria(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/base/editar/categoria",
        method: "POST",
        data: {
            nome: $('#nome').val()
        }
    }).done(function() {
        Swal.fire({
            type: 'success',
            title: 'Categoria criada com sucesso!',
            showConfirmButton: false,
            timer: 1500
          }).then(function() {window.location = '/base';})
     
    }).fail(function(erro){
        console.log(erro)
        Swal.fire({
            type: 'error',
            title: 'Algo deu errado! Não foi possível enviar a solicitação.',
            showConfirmButton: false,
            timer: 1500
          })
    })
}



function atualizarChamado() {
    $(this).prop('disabled', true);

    const chamadoId = $(this).data('chamado-id');
    
    $.ajax({
        url: '/chamados/'+chamadoId,
        method: "PUT",
        data: {
            nome: $('#nome').val(),
            chamado: $('#chamado').val(),
            ativocpu: $('#ativocpu').val(),
            ativomonitor: $('#ativomonitor').val(),
            endereco: $('#endereco').val(),
            numero: $('#numero').val(),
            cep: $('#cep').val(),
            senha: $('#senha').val(),
            transporte: $('#transporte').val(),
            acionamento: $('#acionamento').val(),
            status: $('#status').val(),
            bairro: $('#bairro').val(),
            obs: $('#obs').val(),
            office: $('#office').val(),      
            ramal: $('#ramal').val(),
            logindac: $('#logindac').val(),
            re: $('#re').val(),
            ativoretornomonitor: $('#ativoretornomonitor').val(),
            ativoretornocpu: $('#ativoretornocpu').val(),
            perifericomouse: $('#perifericomouse').val(),
            perifericoteclado: $('#perifericoteclado').val(),
            perifericohead: $('#perifericohead').val(),
            perifericorede: $('#perifericorede').val(),
            analistafield: $('#analistafield').val(),
            gerenteoperador: $('#gerenteoperador').val(),
            asite: $('#asite').val()
        }
    }).done(function() {
        Swal.fire({
            type: 'success',
            title: 'Informações atualizadas com sucesso!',
            showConfirmButton: false,
            timer: 2000
          }).then(function() {window.location = '/chamados/'+chamadoId+'/editar';})
    }).fail(function() {
        Swal.fire({
            type: 'error',
            title: 'Algo deu errado! Não foi possível enviar a solicitação.',
            showConfirmButton: false,
            timer: 1500
          })
    }).always(function() {
        $('#atualizar-chamado').prop('disabled', false);
    })
}
$('#nova-publicacao').on('submit', criarPublicacao);
$('#atualizar-publicacao').on('click', atualizarPublicacao);
$('#deletar-publicacao').on('click', deletarPublicacao);

function criarPublicacao(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/base",
        method: "POST",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
            id_categoria: $('#id_categoria').val(),
            id_usuario: $('#id_usuario').val(),
            id_site: $('#id_site').val(),
            id_cliente: $('#id_cliente').val()
        }
    }).done(function() {

        Swal.fire({
            type: 'success',
            title: 'Publicação criada com sucesso!',
            showConfirmButton: false,
            timer: 1500
          }).then(function() {window.location = '/base';})
     
    }).fail(function(erro){
        console.log(erro)
        Swal.fire({
            type: 'error',
            title: 'Algo deu errado! Não foi possível criar a publicação.',
            showConfirmButton: false,
            timer: 1500
          })
    })
}
function GetURLParameter(sParam)
{
    var sPageURL = window.location.search.substring(1);
    var sURLVariables = sPageURL.split('&');
    for (var i = 0; i < sURLVariables.length; i++) 
    {
        var sParameterName = sURLVariables[i].split('=');
        if (sParameterName[0] == sParam) 
        {
            return sParameterName[1];
        }
    }
};  
var site = GetURLParameter('site');
var cliente = GetURLParameter('cliente');

$('#site').val(site);  
$('#cliente').val(cliente);   



function atualizarPublicacao(evento) {
    evento.preventDefault();
    const publicacaoId = $(this).data('publicacao-id');
    $.ajax({
        url: '/base/'+publicacaoId,
        method: "PUT",
        data: {
            titulo: $('#titulo').val(),
            conteudo: $('#conteudo').val(),
            id_categoria: $('#id_categoria').val(),
            id_usuario: $('#id_usuario').val(),
            id_site: $('#id_site').val(),
            id_cliente: $('#id_cliente').val()
        }
    }).done(function() {

        Swal.fire({
            type: 'success',
            title: 'Publicação atualizada com sucesso!!',
            showConfirmButton: false,
            timer: 1500
          }).then(function() {window.location = '/base/'+publicacaoId;})
     
    }).fail(function(erro){
        console.log(erro)
        Swal.fire({
            type: 'error',
            title: 'Algo deu errado! Não foi possível atualizar a publicação.',
            showConfirmButton: false,
            timer: 1500
          })
    })
}



function deletarPublicacao(evento) {
    evento.preventDefault();
    const publicacaoId = $(this).data('publicacao-id');
    Swal.fire({
        title: "Atenção!",
        text: "Tem certeza que deseja excluir essa publicação? Essa ação é irreversível!",
        showCancelButton: true,
        cancelButtonText: "Cancelar",
        icon: "warning"
    }).then(function(confirmacao) {
        if (!confirmacao.value) return;

    
        $.ajax({
            url: '/base/'+publicacaoId,
            method: "DELETE"
        }).done(function() {
            window.location = '/base';
        }).fail(function() {
            Swal.fire("Ops...", "Erro ao excluir a publicação!", "error");
        });
    })
}

$('#form-ativo').on('submit', consultaAtivo)


function consultaAtivo(evento) {
    evento.preventDefault();
    var serie = $("#serie").val();

    var locador = $("#locador").val();
    console.log(locador)
    
           $("#diframeHolder").removeClass("esconde");
        $("#iframeHolder").attr("src","http://reportcorp.atento.com.br/ReportServer/Pages/ReportViewer.aspx?/GestaoSuprimentos/rpt_equipamento&NOM_EQUIPAMENTO="+locador+"&NOM_EQUIPAMENTO_ANTIGO=&NUM_ATIVO_SAP=&NOM_NUMERO_SERIE="+serie+"&NOM_MARCA=&NOM_MODELO=&COD_GRUPO=&COD_TIPO_EQUIPAMENTO=&COD_SITUACAO_EQUIPAMENTO=&NOM_FORNECEDOR=&NOM_PEDIDO_COMPRA=&NOM_DOC_AQUISICAO=&NUM_CHAMADO_CAD=&NUM_CHAMADO_REG=&NUM_RE_FUNCIONARIO=&NUM_CONTRATO_LEASING=&COD_ESTABELECIMENTO=&COD_CATEGORIA=&COD_CLASSIF_CONTABIL=&DES_EQUIPAMENTO=&DAT_INVENTARIO_INI=&DAT_INVENTARIO_FIM=&DAT_CADASTRO_INI=&DAT_CADASTRO_FIM=&DAT_AQUISICAO_INI=&DAT_AQUISICAO_FIM=&DAT_TSP_ATUALIZACAO_INI=&DAT_TSP_ATUALIZACAO_FIM=&rc:Area=Toolbar&rc:LinkTarget=frmMain&rc:JavaScript=True&rs:ClearSession=true&rc:Parameters=false");
     
    }







$('#form-imdb').on('submit', consultaIMDB)


function consultaIMDB(evento) {
    evento.preventDefault();


    var re = $("#re").val();
    
    //var login = $("#login").val();
           $("#diframeHolder").removeClass("esconde");
        $("#iframeHolder").attr("src","http://reportcorp.atento.com.br/ReportServer/Pages/ReportViewer.aspx?%2fIMDB%2frptUserGrupoProxy&COD_RE="+re+"&rc:Area=Toolbar&rc:LinkTarget=frmMain&rc:JavaScript=True&rs:ClearSession=true&rc:Parameters=false");
     
    }




    $('#form-catraca').on('submit', consultaCatraca)


    function consultaCatraca(evento) {
        evento.preventDefault();
    
    
        var re = $("#re").val();
        var data = $("#data").val();
               $("#diframeHolder").removeClass("esconde");
            $("#iframeHolder").attr("src","http://reportcorp.atento.com.br/ReportServer/Pages/ReportViewer.aspx?%2fCATRACA%2fLogAcessoCatracaFuncionario&MATRICULAS="+re+"&DTMARCACAO="+data+"&rc:Area=Toolbar&rc:LinkTarget=frmMain&rc:JavaScript=True&rs:ClearSession=true&rc:Parameters=false");
         
        }
    
    
    
    
        $('#formulario-cadastro').on('submit', criarUsuario)
        $('#atualizar-cadastro').on('click', atualizarCadastro);
        $('#atualizar-senha').on('click', atualizarSenha);
        $('#deletar-cadastro').on('click', deletarCadastro);
        
        function criarUsuario(evento) {
            evento.preventDefault();
        
            if ($('#senha').val() != $('#confirmar-senha').val()) {
                Swal.fire({
                    type: 'error',
                    title: 'As senhas não coincidem!',
                    showConfirmButton: false,
                    timer: 1500
                  })
                return
            }
        
            $.ajax({
                url: "/usuarios",
                method: "POST",
                data: {
                    nome: $('#nome').val(),
                    login_nt: $('#login_nt').val(),
                    re: $('#re').val(),
                    cargo: $('#cargo').val(),
                    email: $('#email').val(),
                    v_usuarios: $('#v_usuarios').val(),
                    v_bdc_posts: $('#v_bdc_posts').val(),
                    v_bdc_adm: $('#v_bdc_adm').val(),
                    v_imdb: $('#v_imdb').val(),
                    v_gsa: $('#v_gsa').val(),
                    v_mapa_operacional: $('#v_mapa_operacional').val(),
                    id_site: $('#id_site').val(),
                    senha: $('#senha').val(),
                }
            }).done(function() {
                Swal.fire({
                    type: 'success',
                    title: 'Usuário criado com sucesso!!',
                    showConfirmButton: false,
                    timer: 1500
                  })
            }).fail(function(erro){
                console.log(erro)
                Swal.fire({
                    type: 'error',
                    title: 'Algo deu errado! Não foi possível cadastrar o usuário.',
                    showConfirmButton: false,
                    timer: 1500
                  })
            })
        }
        
        
        
        function atualizarCadastro(evento) {
         evento.preventDefault();
        
        
            const usuarioId = $(this).data('usuario-id');
            $.ajax({
                url: '/usuarios/'+usuarioId,
                method: "PUT",
                data: {
                    nome: $('#nome').val(),
                    login_nt: $('#login_nt').val(),
                    re: $('#re').val(),
                    cargo: $('#cargo').val(),
                    email: $('#email').val(),
                    v_usuarios: $('#v_usuarios').val(),
                    v_bdc_posts: $('#v_bdc_posts').val(),
                    v_bdc_adm: $('#v_bdc_adm').val(),
                    v_imdb: $('#v_imdb').val(),
                    v_gsa: $('#v_gsa').val(),
                    v_mapa_operacional: $('#v_mapa_operacional').val(),
                    id_site: $('#id_site').val(),
                }
            }).done(function() {
        
                Swal.fire({
                    type: 'success',
                    title: 'Usuário atualizado com sucesso!!',
                    showConfirmButton: false,
                    timer: 1500
                  }).then(function() {window.location = '/usuarios'})
             
            }).fail(function(erro){
                console.log(erro)
                Swal.fire({
                    type: 'error',
                    title: 'Algo deu errado! Não foi possível editar o usuário.',
                    showConfirmButton: false,
                    timer: 1500
                  })
            })
        }
        
        function atualizarSenha(evento) {
            evento.preventDefault();
        
            if ($('#nova').val() != $('#confirmar-senha').val()) {
                Swal.fire({
                    type: 'error',
                    title: 'As senhas não coincidem!',
                    showConfirmButton: false,
                    timer: 1500
                  })
                return
            }
            const usuarioId = $(this).data('usuario-id');
            $.ajax({
                url: '/usuarios/'+usuarioId+'/atualizar-senha',
                method: "POST",
                data: {
                    nova: $('#nova').val(),
                }
            }).done(function() {
        
                Swal.fire({
                    type: 'success',
                    title: 'Senha atualizada com sucesso!!',
                    showConfirmButton: false,
                    timer: 1500
                  }).then(function() {window.location = '/usuarios'})
             
            }).fail(function(erro){
                console.log(erro)
                Swal.fire({
                    type: 'error',
                    title: 'Algo deu errado! Não foi possível editar o usuário.',
                    showConfirmButton: false,
                    timer: 1500
                  })
            })
        }
        
        
        function deletarCadastro(evento) {
            evento.preventDefault();
            const usuariooId = $(this).data('usuario-id');
            Swal.fire({
                title: "Atenção!",
                text: "Tem certeza que deseja excluir esse usuário? Essa ação é irreversível!",
                showCancelButton: true,
                cancelButtonText: "Cancelar",
                icon: "warning"
            }).then(function(confirmacao) {
                if (!confirmacao.value) return;
        
            
                $.ajax({
                    url: '/usuarios/'+usuarioId,
                    method: "DELETE"
                }).done(function() {
                    window.location = '/usuarios';
                }).fail(function() {
                    Swal.fire("Ops...", "Erro ao excluir o usuário!", "error");
                });
            })
        }


        $('#meuModal').on('shown.bs.modal', function () {
          $('#meuInput').trigger('focus')
        })
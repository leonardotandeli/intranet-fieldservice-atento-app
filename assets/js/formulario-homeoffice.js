$('#formulario').on('submit', criarChamado)
$('#formulario-mapa').on('submit', criarMapa)
$('#atualizar-chamado').on('click', atualizarChamado);
$('#atualizar-mapa').on('click', atualizarMapa);
$('#criar-usuario-usuario').on('click', criarUsuarioChamado);
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

function criarMapa(evento) {
    evento.preventDefault();

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
            template: $('#template').val(),
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
            title: 'Algo deu errado! Não foi possível enviar a solicitação.',
            showConfirmButton: false,
            timer: 1500
          })
    })
}


function atualizarMapa() {
    $(this).prop('disabled', true);

    const mapaId = $(this).data('mapa-id');
    console.log(mapaId)

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
            template: $('#template').val(),
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

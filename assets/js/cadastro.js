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
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
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
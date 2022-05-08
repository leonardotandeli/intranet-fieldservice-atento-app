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
        Swal.fire({
            type: 'error',
            title: 'Algo deu errado! Login ou senha incorretos!.',
            showConfirmButton: false,
            timer: 2500
          })

        })
}
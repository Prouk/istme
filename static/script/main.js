$(document).ready(function (){
    $('.modalButton').on('click', function (){
        $('.modal').fadeToggle(250)
    })
    $('.apiContainer').on('click', '#buttonApi1', function (){
        let data = {
            text: $('input[name=test]').val()
        }
        $.ajax({
            method: 'POST',
            url: 'apiTest',
            data: data
        })
            .done(function( resp ) {
                $('.apiResp1').text(resp.Msg)
            });
    })
})
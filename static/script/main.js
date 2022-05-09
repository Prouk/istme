$(document).ready(function (){
    $('.modalButton').on('click', function (){
        $('.modal').fadeToggle(250)
    })

    $('.apiTest').on('click', 'button', function (){
        $('.apiTest button').addClass('rotating')
        $('.apiTest button').prop("disabled",true);
        let data = {
            text: $('.apiTest input[name=test]').val()
        }
        $.ajax({
            method: 'POST',
            url: 'apiTest',
            data: data
        })
            .done(function(resp) {
                $('.apiTest .response').text(resp.Msg)
                $('.apiTest button').removeClass('rotating')
                $('.apiTest button').prop("disabled",false);
            });
    })

    $('.apiFfxiv').on('click', 'button', function (){
        $('.apiFfxiv button').addClass('rotating')
        $('.apiFfxiv button').prop("disabled",true);
        let data = {
            name: $('.apiFfxiv input[name=name]').val().replace(' ', '+'),
            world: $('.apiFfxiv select[name=world]').val()
        }
        $('.apiFfxiv .response').text("request : \n"+JSON.stringify(data, undefined, 2))
        $.ajax({
            method: 'GET',
            url: 'apiFfxiv',
            data: data
        })
            .done(function(resp) {
                $('.apiFfxiv .response').html(ffxivCharConstructor(resp))
                $('.apiFfxiv button').removeClass('rotating')
                $('.apiFfxiv button').prop("disabled",false);
            });
    })
})

function ffxivCharConstructor(data) {
    return `${data.Title}<br><a href="${data.CharUrl}" target="_blank"><img src="${data.ImgUrl}" alt="CharImg" style="height: 170px;border-radius: 7px;"></a><br><img src="${data.JobImg}" alt="jobImg"><br>${data.Level}<br>${data.GrandCompany}`
}
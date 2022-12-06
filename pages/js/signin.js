// sigin business and event handle
(() => {
    'use strict'

    $("#signin-submit").on("click", () => {
      
        const $form = $("#signin-form")
        const body = $form.serializeObject()
        console.log(body)
        $.ajax({
            type: 'post',
            url: '../../api/signin',
            contentType: 'application/json',
            dataType: 'json',
            data: JSON.stringify(body),
            success: function(data, textStatus, jqXHR) {
                const token = jqXHR.getResponseHeader('Authorization')
                localStorage.setItem("authorization", token)
                window.location.href = '../index.html'
            },
            error: function(xhr, textStatus, errorThrown) {
                console.log(xhr)
                alert(xhr.responseJSON.msg)
            },
        });
    })

})()


  

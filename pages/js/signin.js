// sigin business and event handle
(() => {
    'use strict';

    $("#signin-submit").on("click", () => {
      
        const $form = $("#signin-form");
        const body = $form.serializeObject();

        $.ajax({
            type: 'post',
            url: '../../api/signin',
            contentType: 'application/json',
            dataType: 'json',
            data: JSON.stringify(body),
            success: function(data, textStatus, jqXHR) {
                const token = jqXHR.getResponseHeader('Authorization');
                localStorage.setItem("authorization", token);
                localStorage.setItem("role", data.role);
                if (data.role == "admin") {
                    window.location.href = 'admin.html';
                } else {
                    window.location.href = 'post-list.html';
                }
                
            },
            error: function(xhr, textStatus, errorThrown) {
                alert(xhr.responseJSON.msg);
            },
        });
    });

})();


  

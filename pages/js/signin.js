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
                let claim = $.getJwtClaim(token);
                
                if (claim.role == "admin") {
                    window.location.href = 'admin.html';
                } else {
                    window.location.href = 'post-list.html';
                }
                
            }
        });
    });

})();


  

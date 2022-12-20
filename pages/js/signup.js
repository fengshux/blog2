// sigup business and event handle
(() => {
    'use strict';
    window.auth = $.getAuth(localStorage.getItem("authorization"));
    
    // customer password validate
    var password = document.getElementById("password"), 
        confirm_password = document.getElementById("confirm-password"); 
    function validatePassword() { 
        if (password.value != confirm_password.value) 
            confirm_password.setCustomValidity("Passwords Don't Match"); 
        else 
            confirm_password.setCustomValidity('');
    } 
    password.onchange = validatePassword; 
    confirm_password.onkeyup = validatePassword;

    document.getElementById("singup-submit").addEventListener("click", event => {        
        const form = $("#sinup-form")[0];
        if (!form.checkValidity()) {
            event.preventDefault();
            event.stopPropagation();
        }
        form.classList.add('was-validated');
        
        if (form.checkValidity()) {
            
            const body = $(form).serializeObject();
            $.ajax({
                type: 'post',
                url: '../../api/user',
                contentType: 'application/json',
                dataType: 'json',
                data: JSON.stringify(body),
                success: function(data, textStatus, jqXHR) {
                    
                    if (window.auth.role == "admin") {
                            // admin 创建用户，跳转到admin页            
                            window.location.href = "./admin.html";
                        } else {
                            // 如是开放注册，没登录注册，跳转到登录页
                            alert("注册成功，请登录");
                            window.location.href = "./signin.html";
                        }
                    
                }
            });            
        }

    });

})();




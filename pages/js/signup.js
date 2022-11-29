// sigup business and event handle
(() => {
    'use strict'
    $.fn.serializeObject = function() {
      var o = {};
      var a = this.serializeArray();
      $.each(a, function() {
          if (o[this.name]) {
              if (!o[this.name].push) {
                  o[this.name] = [o[this.name]];
              }
              o[this.name].push(this.value || '');
          } else {
              o[this.name] = this.value || '';
          }
      });
      return o;
  };
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
      
        const form = $("#sinup-form")[0]
        if (!form.checkValidity()) {
          event.preventDefault()
          event.stopPropagation()
        }
        form.classList.add('was-validated')

        if (form.checkValidity()) {
          
          const body = $(form).serializeObject()
          console.log(body)
          fetch("../../api/user", {
            method: 'POST', // *GET, POST, PUT, DELETE, etc.
            mode: 'cors', // no-cors, *cors, same-origin
            cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
            credentials: 'same-origin', // include, *same-origin, omit
            headers: {
              'Content-Type': 'application/json'
              // 'Content-Type': 'application/x-www-form-urlencoded',
            },
            redirect: 'follow', // manual, *follow, error
            referrerPolicy: 'no-referrer', // no-referrer, *no-referrer-when-downgrade, origin, origin-when-cross-origin, same-origin, strict-origin, strict-origin-when-cross-origin, unsafe-url
            body: JSON.stringify(body),
          }).then(resp => resp.json())
          .then(data => {
            if (data.msg){ 
              alert(data.msg)
            } else {
              alert("注册成功，请登录")
              window.location.href = "./signin.html"
            }
          })
          .catch(err => {
            console.log("error=", err)
          });
      }

  }) 

})()


  

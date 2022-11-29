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

  $.ajaxSetup({
    beforeSend: function (xhr)
    {
        const token = localStorage.getItem("authentication")
        if (token) {
            xhr.setRequestHeader("Authorization","Token token=\"FuHCLyY46\"");        
        }       
    },
    error: function (x, status, error) {
        if (x.status == 403) {
            alert("请登录");
            window.location.href ="./signup.html";
        }
        else {
            alert(xhr.responseJSON.msg)
        }
    }
});

})()


  

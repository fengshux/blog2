// sigup business and event handle
(() => {
    'use strict';
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
    
    $.getQueryVars = function ()
    {
        return new Proxy(new URLSearchParams(window.location.search), {
            get: (searchParams, prop) => {
                let values = searchParams.getAll(prop);
                if (values.length == 0) {
                    return undefined;
                } else if (values.length == 1) {
                    return values[0];
                }
                return values;
            },
        });
    };
    
    $.getJwtClaim = function (token) {
        var base64Url = token.split('.')[1];
        var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
        var jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
            return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
        }).join(''));

        return JSON.parse(jsonPayload);
    };

    $.getAuth = function(token) {
        if (!token) {
            return null;
        }
        let claim = $.getJwtClaim(token);
        if (claim.exp * 1000 <= Date.now()) {
            return null;
        }
        return claim;
    };
    
    // global setting
    $.ajaxSetup({
        beforeSend: function (xhr)
        {
            const token = localStorage.getItem("authorization");
            if (token) {
                xhr.setRequestHeader("Authorization", token);        
            }       
        },
        error: function (x, status, error) {
            if (x.status == 401) {
                alert("请登录");
                window.location.href ="./signin.html";
            }
            else {
                if (x.responseJSON && x.responseJSON.msg) {
                    alert(x.responseJSON.msg);
                } else {
                    alert(error);
                }
                
            }
        }
    });
})();




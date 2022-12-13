(()=>{
    "use strict";


    $.ajax({
        type: 'get',
        url: '../../api/setting',
        contentType: 'application/json',
        dataType: 'json',
        // data: JSON.stringify(body),
        success: function(data, textStatus, jqXHR) {
            console.log(data);
            renderSettings(data);
            
        },
        error: function(xhr, textStatus, errorThrown) {
            console.log(xhr);
            alert(xhr.responseJSON.msg);
        },
    });





    $.ajax({
        type: 'get',
        url: '../../api/user',
        contentType: 'application/json',
        dataType: 'json',
        // data: JSON.stringify(body),
        success: function(data, textStatus, jqXHR) {
            console.log(data);
            renderUserList(data);
            
        },
        error: function(xhr, textStatus, errorThrown) {
            console.log(xhr);
            alert(xhr.responseJSON.msg);
        },
    });
    
    function renderSettings(data) {
        for (var v of data) {
            switch( v["key"] ) {
            case 'about' :
                $('#about').text(v.data.content);
               
            }
        }
    }
    
    function renderUserList(data) {
        
        if (data.list && data.list.length > 0) {
            var html = '';
            for (var d of data.list) {
                html += `<tr><td>${d.username}</td><td>${d.nickname}</td>`
                    +`<td>${d.email}</td><td>${d.gender}</td><td>${d.role}</td>`
                    +`<td><button type="button" class="btn btn-sm btn-primary">修改密码</button>`
                    +`<button type="button" class="btn btn-sm btn-danger">删除</button></td></tr>`;
            }
            $("#user-list-table").append(html);
        }
    }

    
})();

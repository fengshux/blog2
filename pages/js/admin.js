(()=>{
    "use strict";

    // 请求配置信息
    $.ajax({
        type: 'get',
        url: '../../api/setting',
        contentType: 'application/json',
        dataType: 'json',
        // data: JSON.stringify(body),
        success: function(data, textStatus, jqXHR) {
            console.log(data);
            renderSettings(data);
            
        }
    });

    // 请求用户列表
    $.ajax({
        type: 'get',
        url: '../../api/user',
        contentType: 'application/json',
        dataType: 'json',
        // data: JSON.stringify(body),
        success: function(data, textStatus, jqXHR) {
            console.log(data);
            renderUserList(data);
            
        }
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
                    +`<td><button type="button" class="btn btn-sm btn-primary op-btn" value="${d.id}">修改密码</button>`
                    +`<button type="button" class="btn btn-sm btn-danger op-btn" value="${d.id}">删除</button></td></tr>`;
            }
            $("#user-list-table").append(html);
        }


        // add click 
        $("button.op-btn").on("click", function() {

            if ($(this).text() == "修改密码")  {
                // TODO 修改密码
                
            } else if ($(this).text() == "删除") {
                if (confirm("删除后不可恢复，确定删除吗？")) {
                    // TODO 删除用户
                }
            }
            
            
        });

    }

    
})();

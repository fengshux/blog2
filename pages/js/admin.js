// 管理页面主逻辑
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
        url: '../../api/user?page=1&size=1000',
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


        // 用户列表事件绑定
        $("button.op-btn").on("click", function(e) {
            let id = $(this).val();
           
            if ($(this).text() == "修改密码")  {
                // TODO 修改密码
                $("#modalChangePassword").modal('show');
                $("#modalChangePassword").data("id", id);
                                
            } else if ($(this).text() == "删除") {
                if (confirm("删除后不可恢复，确定删除吗？")) {
                    // TODO 删除用户
                    $.ajax({
                        type: 'delete',
                        url: `../../api/user/${id}`,
                        contentType: 'application/json',
                        dataType: 'json',
                        success: function(data, textStatus, jqXHR) {
                            window.location.reload();
                        }
                    });
                }
            }            
        });

        // 设置事件绑定
        $("#about-submit").on("click", function(e) {
            
            var content = $('#about').val();
            var body = {"key": "about", "data":{"content": content}};
            
            $.ajax({
                type: 'put',
                url: '../../api/setting',
                contentType: 'application/json',
                dataType: 'json',
                data: JSON.stringify(body),
                success: function(data, textStatus, jqXHR) {
                    alert("设置成功");
                }
            });
        });

    }
    
})();

// 修改密码模态框逻辑
(()=>{
    "use strict";
    // 修改密码模态框事件绑定
    $("#modalChangePassword").on("hide.bs.modal", function(e) {
        $("#modalChangePassword").data("id", "");
    });

    
    $("#change-password-form button").on("click", function(e) {
        let id = $("#modalChangePassword").data("id");
        console.log(id);
        let body = $("#change-password-form").serializeObject();
        $.ajax({
            type: 'patch',
            url: `../../api/user/${id}/password`,
            contentType: 'application/json',
            dataType: 'json',
            data: JSON.stringify(body),
            success: function(data, textStatus, jqXHR) {
                $("#modalChangePassword").modal('hide');
            }
        });        
    });
    
})();

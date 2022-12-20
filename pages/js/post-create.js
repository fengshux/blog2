(function(){
    'use strict';

    let auth = $.getAuth(localStorage.getItem("authorization"));
    if (!auth) {
        alert("请先登录");
        window.location.href =  "signin.html";
    }
    
    const search = $.getQueryVars();
    $(function() {
        // render summernote
        $('#summernote').summernote({
            tabsize: 2,
            height: 300,
        });
      

        // 如果是编辑页面，则赋值
        if (search.id) {
            // 如果是编辑页面，可以删除文章
            $("#button-group").append(`<button type="button" class="btn btn-danger" id="post-delete">删除</button>`);
            
            $.ajax({
                type: 'get',
                url: `../../api/post/${search.id}`,
                contentType: 'application/json', 
                dataType: 'json',
                success: function(data, textStatus, jqXHR) {
                    $("#title").val(data.title);
                    $('#summernote').summernote('code', data.body);
                }
            });


            
        }


        // 提交按钮事件邦定
        $("#button-group button").on("click",  function(e) {

            const btn = $(this).text();
            
            if( btn == "删除") {
                deletePost(search.id);
            } else if (btn == "提交") {
                const $form = $("#post-create");
                const body = $form.serializeObject();
                body.body = $('#summernote').summernote('code');
                body.status = 'published';
                if ( search.id   ) {
                    editPost(search.id, body);
                } else {
                    createPost(body);
                }
                
            }
            
        });
        
    });


    function editPost(id, post) {
        // 更新            
            $.ajax({
                type: 'put',
                url: `../../api/post/${search.id}`,
                contentType: 'application/json',
                dataType: 'json',
                data: JSON.stringify(post),
                success: function(data, textStatus, jqXHR) {
                    window.location.href = `./post-detail.html?id=${search.id}`;
                }
            });
    }

    function createPost(post) {
            // 创建                    
            $.ajax({
                type: 'post',
                url: '../../api/post',
                contentType: 'application/json',
                dataType: 'json',
                data: JSON.stringify(post),
                success: function(data, textStatus, jqXHR) {
                    window.location.href = "./post-list.html";
                }
            });
    }

    function deletePost(id) {
            // 创建                    
            $.ajax({
                type: 'delete',
                url: `../../api/post/${id}`,
                contentType: 'application/json',
                dataType: 'json',
                success: function(data, textStatus, jqXHR) {
                    window.location.href = "./post-list.html";
                }
            });
    }
    



    
})();

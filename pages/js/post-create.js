(function(){
    'use strict';
    $(document).ready(function() {
        $('#summernote').summernote({
            tabsize: 2,
            height: 200,
        });
    });


    // 提交按钮事件邦定
    $("#post-submit").on("click", () => {      
        const $form = $("#post-create");
        const body = $form.serializeObject();
        body.body = $('#summernote').summernote('code');
        body.status = 'published';
        $.ajax({
            type: 'post',
            url: '../../api/post',
            contentType: 'application/json',
            dataType: 'json',
            data: JSON.stringify(body),
            success: function(data, textStatus, jqXHR) {
                window.location.href = "./post-list.html";
            }
        });
    });


    
})();

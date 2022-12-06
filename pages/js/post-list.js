(function(){

    $.ajax({
        type: 'get',
        url: '../../api/post',
        contentType: 'application/json',
        dataType: 'json',
        // data: JSON.stringify(body),
        success: function(data, textStatus, jqXHR) {
            console.log(data)
            renderPost(data)
            
        },
        error: function(xhr, textStatus, errorThrown) {
            console.log(xhr)
            alert(xhr.responseJSON.msg)
        },
    });

    function renderPost(data) {
        
        if (data.list && data.list.length > 0) {

            // render first post
            const first = data.list[0];
            if (first.body.length > 30) {
                first.body.length = 30;
            }
            const html = `<h1 class="display-4 fst-italic">${first.title}</h1>
                     <p class="lead my-3">${first.body}</p>
                     <p class="lead mb-0"><a href="post-detail.html?post_id=${first.id}" class="text-white fw-bold">Continue reading...</a></p>`;

            $("#first-blog").html(html)            
        }
    }
    
})();

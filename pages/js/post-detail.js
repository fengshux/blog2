(function(){

    const params = new Proxy(new URLSearchParams(window.location.search), {
        get: (searchParams, prop) => searchParams.get(prop),
    });
    let id = params.id;
    $.ajax({
        type: 'get',
        url: `../../api/post/${id}`,
        contentType: 'application/json', 
        dataType: 'json',
        success: function(data, textStatus, jqXHR) {
            renderPost(data);            
        }
    });

    function renderPost(p) {

        const html = `<h2 class="blog-post-title mb-1">${p.title}</h2>
                        <p class="blog-post-meta">
                          ${$.format.date(p.create_time, "yyyy-MM-dd HH:mm")} by ${p.user? p.user.nickname : "注消用户"}
                          <a href="post-create.html?id=${p.id}">编辑</a>
                        </p>
                        <p>${p.body}</p>`;
        
        $("#post-detail").html(html);
        
    }
    
})();

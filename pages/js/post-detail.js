(function(){

    const params = new Proxy(new URLSearchParams(window.location.search), {
        get: (searchParams, prop) => searchParams.get(prop),
    });
    let id = params.post_id;
    $.ajax({
        type: 'get',
        url: `../../api/post/${id}`,
        contentType: 'application/json', 
        dataType: 'json',
        success: function(data, textStatus, jqXHR) {
            renderPost(data);            
        }
    });

    function renderPost(data) {

        const html = `<h2 class="blog-post-title mb-1">${data.title}</h2>
                        <p class="blog-post-meta">January 1, 2021 by <a href="#">Mark</a></p>
                        <p>${data.body}</p>`;
        
        $("#post-detail").html(html);
        
    }
    
})();

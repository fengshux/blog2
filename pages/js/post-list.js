(function(){
    // 设置每每多少数据
    const PAGE_SIZE = 10;
    
    function getPostData(page, size) {
        $.ajax({
            type: 'get',
            url: `../../api/post?page=${page}&size=${size}`,
            contentType: 'application/json',
            dataType: 'json',
            // data: JSON.stringify(body),
            success: function(data, textStatus, jqXHR) {         
                renderPost(data, page, size);            
            }
        });
    }   

    function renderPost(data, page, size) {
        
        
        if (data.list && data.list.length > 0) {
            let list = data.list;
            if (page == 1) {
                // render first post
                let first = list.shift();
                if (first.body.length > 60) {
                    first.body = first.body.substr(0,60);
                }
                const html = `<h1 class="display-4 fst-italic">${first.title}</h1>
                     <p class="lead my-3">${first.body}</p>
                     <p class="lead mb-0"><a href="post-detail.html?id=${first.id}" class="text-white fw-bold">Continue reading...</a></p>`;

                $("#first-blog").html(html);

            }           

            // render post list
            $("#post-list article").remove();
            $("#post-list nav.blog-pagination").remove();
            for (var p of list) {
                // 防止文章过长
                if (p.body.length > 200) {
                    p.body = p.body.substr(0,200);
                }
                
                let article = `<article class="blog-post">
                        <h2 class="blog-post-title mb-1">${p.title}</h2>
                        <p class="blog-post-meta">${$.format.date(p.create_time, "yyyy-MM-dd HH:mm")} by ${p.user? p.user.nickname : "注消用户"}</p>
                        <p>${p.body}</p>
                        <p class="lead mb-0"><a href="post-detail.html?id=${p.id}" class="text-black">Continue reading...</a></p>
                    </article>`;
                $("#post-list").append(article);
            }
            
            // render page button
            $("#post-list").append(`
                    <nav class="blog-pagination" aria-label="Pagination">                        
                        <a class="btn rounded-pill ${page == 1?'btn-outline-secondary disabled':'btn-outline-primary'} " data-page="${page-1}">Newer</a>
                        <a class="btn rounded-pill ${page*size >= data.total?'btn-outline-secondary disabled':'btn-outline-primary'}" data-page="${page+1}">Older</a>
                    </nav>
            `);


            // 分页事件绑定
            $(".blog-pagination a:not(.disabled)").on('click', function(e) {
                getPostData($(this).data("page"), PAGE_SIZE);
                
            });
        }
    }
   
    // 请求数据
    getPostData(1, PAGE_SIZE);

})();

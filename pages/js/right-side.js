(()=>{
    "use strict";

    let html = `     <div class="position-sticky" style="top: 2rem;">
                        <div class="p-4 mb-3 bg-light rounded" id="about-panel">
                            <h4 class="fst-italic">关于</h4>
                            <p class="mb-0"></p>
                        </div>

                        <div class="p-4">
                            <h4 class="fst-italic">Archives</h4>
                            <ol class="list-unstyled mb-0">
                                <li><a href="#">March 2021</a></li>
                                <li><a href="#">February 2021</a></li>
                                <li><a href="#">January 2021</a></li>
                                <li><a href="#">December 2020</a></li>
                                <li><a href="#">November 2020</a></li>
                                <li><a href="#">October 2020</a></li>
                                <li><a href="#">September 2020</a></li>
                                <li><a href="#">August 2020</a></li>
                                <li><a href="#">July 2020</a></li>
                                <li><a href="#">June 2020</a></li>
                                <li><a href="#">May 2020</a></li>
                                <li><a href="#">April 2020</a></li>
                            </ol>
                        </div>

                        <div class="p-4">
                            <h4 class="fst-italic">Elsewhere</h4>
                            <ol class="list-unstyled">
                                <li><a href="#">GitHub</a></li>
                                <li><a href="#">Twitter</a></li>
                                <li><a href="#">Facebook</a></li>
                            </ol>
                        </div>        
                    </div>`;

    console.log($("#right-side"));
    $("#right-side").html(html);


    function renderAbout( data ){
        $("#about-panel p").text(data.data.content);
    }
    

    $(function (){
        $.ajax({
            type: 'get',
            url: '../../api/setting/about',
            contentType: 'application/json',
            dataType: 'json',
            success: function(data, textStatus, jqXHR) {         
                renderAbout(data);
                
            }
        });
    });

    
})();

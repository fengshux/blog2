// business for header
(() => {
    'use strict';

    $(function() {
        
        window.auth = $.getAuth(localStorage.getItem("authorization"));

        // render header
        $("#header").html(
            `<div class="container">
                    <div id="header-compose" class="d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start">
                        <a href="/" class="d-flex align-items-center mb-2 mb-lg-0 text-dark text-decoration-none">
                            <img class="bi me-2" width="40" height="32" role="img" aria-label="Bootstrap" src="../assets/logo-40x40.png"></img>
                        </a>

                        <ul class="nav col-12 col-lg-auto me-lg-auto mb-2 justify-content-center mb-md-0">
                            <li><a href="/" class="nav-link px-2 link-secondary">Home</a></li>
                            <!--
                            <li><a href="#" class="nav-link px-2 link-dark">Inventory</a></li>
                            <li><a href="#" class="nav-link px-2 link-dark">Customers</a></li>
                            <li><a href="#" class="nav-link px-2 link-dark">Products</a></li>
                            -->
                        </ul>
                        <form class="col-12 col-lg-auto mb-3 mb-lg-0 me-lg-3" role="search">
                            <input type="search" class="form-control" placeholder="Search..." aria-label="Search">
                        </form>                        
                    </div>
            </div>`
        );        

        var loginHtml;
        if (window.auth) {
            var ulHtml = "";
            if (window.auth.role != "admin") {
                ulHtml = `<li><a class="dropdown-item" href="post-create.html">写文章</a></li>
                <li><hr class="dropdown-divider"></li>`;
            }
            
            
            loginHtml = `<div class="dropdown text-end">
                <a href="#" class="d-block link-dark text-decoration-none dropdown-toggle" data-bs-toggle="dropdown" aria-expanded="false">
                <img src="../assets/avatar.jpeg" alt="mdo" width="32" height="32" class="rounded-circle">
                </a>
                <ul class="dropdown-menu text-small">
                ${ulHtml}
                <li><a class="dropdown-item" href="#" id="sign-out-btn">Sign out</a></li>
                </ul>
                </div>`;
        } else {            

            loginHtml = `<div class="text-end">
                <a type="button" class="btn btn-outline-primary me-2" href="signin.html">Sign in</a>                            
                </div>`; 
        }

        $("#header-compose").append(loginHtml);


        // sign out 事件
        $("a#sign-out-btn").on('click', ()=> {
            localStorage.removeItem("authorization");
            localStorage.removeItem("role");
            window.location.href = "post-list.html";
        });        
    });
    


})();

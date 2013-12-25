{{define  "header"}}
 {{$config:=.SiteConfig}}
<header id="masthead" role="banner">
    <div class="row-fluid top-banner">
        <div class="container">
            <div class="banner-overlay"></div>
            <a class="brand brand-image" href="{{$config.SiteURL}}" title="北飘漂—.Net,C#,C/C++,Golang,TypeScript,NoSql,Linux等开发技术" rel="home">
                <img src="{{$config.StaticURL}}/static/{{$config.ThemeName}}/img/logo.png" alt="北京飘漂—.Net,C#,C/C++,Golang,TypeScript,NoSql,Linux等开发技术"><h1><small>一个乐于分享| .Net | C/C++ | Golang | TypeScript | NoSql | linux |等技术的IT博客</small></h1>
            </a>
            <div class="top-banner-social pull-right"></div>
        </div>
    </div>
    <div class="row-fluid bluth-navigation">
        <div class="container">
            <div class="mini-logo">
                <a class="mini mini-image" href="{{$config.SiteURL}}" title="北飘漂—.Net,C#,C/C++,Golang,TypeScript,NoSql,Linux等开发技术" rel="home">
                    <img src="{{$config.StaticURL}}/static/{{$config.ThemeName}}/img/logo_normal.png" alt="北飘漂—.Net,C#,C/C++,Golang,TypeScript,NoSql,Linux等开发技术"></a>
            </div>
            <div class="navbar navbar-inverse">
                <div class="navbar-inner">
                    <div class="visible-tablet visible-phone bl_search">
                        <form action="{{$config.SiteURL}}" method="get" class="searchform" role="search">
                            <fieldset>
                                <a href="{{$config.SiteURL}}#"><i class="icon-search-1"></i></a>
                                <input type="text" name="s" value="" placeholder="搜索...">
                            </fieldset>
                        </form>
                    </div>
                    <!-- Responsive Navbar Part 1: Button for triggering responsive navbar (not covered in tutorial). Include responsive CSS to utilize. -->
                    <button data-target=".nav-collapse" data-toggle="collapse" class="btn btn-navbar" type="button"><i class="icon-menu-1"></i></button>
                    <div class="nav-collapse collapse">
                        <ul id="menu-top" class="nav">
                            <li id="menu-item-home" class="menu-item menu-item-type-custom menu-item-object-custom menu-item-home"><a href="{{$config.SiteURL}}">首页</a></li>
                            <li id="menu-item-14" class="menu-item menu-item-type-custom menu-item-object-custom"><a href="#">博客介绍</a></li>
                            {{range $index,$value:=.CategoryList}}
                               {{$nodes:=$value.Nodes}}
                                {{if $nodes}}
                                   <li id="menu-item-{{$value.Name}}" class="menu-item menu-item-type-taxonomy menu-item-object-category dropdown">
                                        <a href="{{$config.SiteURL}}/category/{{$value.Name}}" class="dropdown-toggle" data-toggle="dropdown">{{$value.Title}}<i class="icon-angle-down"></i></a>
                                        <ul class="dropdown-menu">
                                            {{range $inx,$val:=$nodes}}
                                            <li id="menu-item-{{$val.Name}}" class="menu-item menu-item-type-taxonomy menu-item-object-category"><a href="{{$config.SiteURL}}/node/{{$val.Name}}">{{$val.Title}}</a></li>
                                            <!--.dropdown-->
                                            {{end}}
                                        </ul>
                                    </li> 
                                {{else}}
								 <li id="menu-item-{{$value.Name}}" class="menu-item menu-item-type-taxonomy menu-item-object-category">
                                        <a href="{{$config.SiteURL}}/category/{{$value.Name}}">{{$value.Title}}</a>
                                    </li> 
                                {{end}}
                            {{end}}
                        </ul>
                    </div>
                </div>
                <!-- /.navbar-inner -->
                <div class="nav-line" style="left: 0px; width: 60px;"></div>
            </div>
            <div class="bl_search visible-desktop nav-collapse collapse">
                <form action="{{$config.SiteURL}}" method="get" class="searchform" role="search">
                    <fieldset>
                        <a href="{{$config.SiteURL}}#"><i class="icon-search-1"></i></a>
                        <input type="text" name="s" value="" placeholder="搜索...">
                    </fieldset>
                </form>
            </div>
        </div>
    </div>
</header>
{{end}}
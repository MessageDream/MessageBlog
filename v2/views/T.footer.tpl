{{define "footer"}}
 {{$config:=.SiteConfig}}
  <footer id="colophon" class="site-footer" role="contentinfo">
            <div class="container">
                <div class="row-fluid" id="footer-body">
                    <div id="text-6" class="span3 widget_text">
                        <h3 class="widget-head">博客介绍</h3>
                        <div class="textwidget">
                            <div style="color: #45B0EE;">北飘漂是一个乐于分享| .Net | C/C++ | Golang | TypeScript | NoSql | linux |等技术的IT博客</div>
                            <br>
                             <!--<div style="color: #45B0EE;">
                                如果您觉得博客的内容对您的学习有所帮助：<a href="http://me.alipay.com/tystudio">
                                    <img src="{{$config.StaticURL}}/static/{{$config.ThemeName}}/img/pay_encourage.png" title="小小赞助大大帮助">
                                </a>
                            </div>-->
                            <br>
                            <a href="{{$config.SiteURL}}/root">登录</a>
                            <br>
                            <!-- <a href="http://www.tystudio.net/index.php/profile/">联系MessageDream</a>-->
                        </div>
                    </div>
                   <!-- <div id="bl_html-3" class="cleanwidget  nopadding span3 bl_html">
                        <h3 class="widget-head">免费提供MT主机博客空间</h3>
                        <div class="widget-body">
                            <p>
                                如果你想建立独立的博客，天屹可以为你提供免费的mediatemple主机空间，<a href="https://www.mediatemple.net/webhosting/shared/" traget="_blank">点此了解mt</a>。域名需要自己提供。<br>
                                主机支持：<br>
                                PHP 5<br>
                                Perl v5.8.4<br>
                                Python v2.4.4<br>
                                Apache 2.0.54<br>
                                MySQL 5.1.26
                            </p>
                        </div>
                    </div>-->
                    <div id="text-9" class="span3 widget_text">
                        <h3 class="widget-head">友情链接</h3>
                        <div class="textwidget">
                            <a href="http://blog.bossma.cn/" target="_blank">波斯马的博客</a>
                            <br>
                        </div>
                    </div>
                    <div id="bl_newsletter-3" class="span3 bl_newsletter">
                        <h3 class="widget-head">订阅到邮箱</h3>
                        <div class="widget-body">
                            <p>输入邮箱订阅本博客的精彩内容</p>
                            <div class="input-append">
                                <input class="bl_newsletter_email" type="text" placeholder="邮件地址">
                                <button class="btn bluth green btn-primary" type="button">订阅</button>
                            </div>
                        </div>
                        <script type="text/javascript">
                            window.locale = {
                                "no_email_provided": '未提供Email',
                                "thank_you_for_subscribing": '感谢您的订阅，一封确认订阅邮件已发送到您的邮箱',
                            };
                        </script>
                    </div>
                </div>
            </div>
            <!-- .site-info -->
            <div class="row-fluid" id="footer-bottom">
                <span>Copyright&nbsp;&nbsp;©&nbsp;&nbsp;2014&nbsp;&nbsp;</span><a href="{{$config.SiteURL}}">{{$config.SiteTitle}}</a> <span>&nbsp;&nbsp;京ICP备14000585号-1&nbsp;|&nbsp;Powered By&nbsp;&nbsp;<a href="http://golang.org/" target="_blank">Golang</a>&nbsp;&nbsp;and&nbsp;&nbsp;<a href="http://www.mongodb.org/" target="_blank">MongoDB</a>&nbsp;&nbsp;静态文件和图片存储由<a href="https://portal.qiniu.com" target="_blank">七牛云存储</a>提供
                    <div style="float: right; margin-right: 30px">
                        <div id="clicki_widget_5773" style="display: inline;">
                            <a href="http://www.clicki.cn/" target="_blank" style="display: inline; text-decoration: none; border: none; padding: 0; margin: 0;">
                                <img src="{{$config.StaticURL}}/static/{{$config.ThemeName}}/img/16_1.gif" alt="Clicki: We Love Data" title="Clicki: We Love Data" style="display: inline; border: none; padding: 0; margin: 0;"></a>
                        </div>
                        <script>
                            (function (i, s, o, g, r, a, m) {
                                i['GoogleAnalyticsObject'] = r; i[r] = i[r] || function () {
                                    (i[r].q = i[r].q || []).push(arguments)
                                }, i[r].l = 1 * new Date(); a = s.createElement(o),
                                m = s.getElementsByTagName(o)[0]; a.async = 1; a.src = g; m.parentNode.insertBefore(a, m)
                            })(window, document, 'script', '//www.google-analytics.com/analytics.js', 'ga');

                            ga('create', 'UA-39732799-1', 'blog.buxuxiao.com');
                            ga('send', 'pageview');
                        </script>
                        <script type="text/javascript">
                            (function () {
                                var c = document.createElement('script');
                                c.type = 'text/javascript';
                                c.async = true;
                                c.src = ('https:' == document.location.protocol ? 'https://' : 'http://') + 'www.clicki.cn/boot/49009';
                                var h = document.getElementsByTagName('script')[0];
                                h.parentNode.insertBefore(c, h);
                            })();
                        </script>
                    </div>
                </span>
            </div>
        </footer>
 {{end}}
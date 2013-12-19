<!-- #content 右侧-->
{{define "right_side"}}
 {{$config:=.SiteConfig}}
     <aside id="side-bar" class="span4">
                    <div id="bl_author-2" class="box bl_author">
                        <h3 class="widget-head">Me</h3>
                        <img src="{{$config.StaticURL}}/static/{{$config.ThemeName}}/img/lowpolyscenery-870x400.jpg"><div class="widget-body">
                            <div class="bl_author_img">
                                <img src="{{$config.StaticURL}}/static/{{$config.ThemeName}}/img/messagedream.jpg">
                            </div>
                            <div class="bl_author_bio">
                                <h3>MessageDream</h3>
                                <p class="muted">敢于尝试新东西，勇于面对新挑战。</p>
                                <div style="height: 40px; width: 175px; margin: 0 auto;">
                                    <ul id="social-list">
                                        <li><a href="http://weibo.com/u/2017203357" target="_blank"><span class="sprite-round-icon-sina">新浪微博</span></a></li>
                                        <li><a href="http://t.qq.com/messagedream" target="_blank"><span class="sprite-round-icon-qq">腾讯微博</span></a></li>
                                        <li><a href="#"><span class="sprite-round-icon-weixin">微信</span></a>
                                        </li>
                                    </ul>
                                </div>
                                <p></p>
                            </div>
                        </div>
                    </div>
                    <div id="bl_newsletter-2" class="box bl_newsletter">
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
                    <div id="views-2" class="box widget_views">
                        <h3 class="widget-head">热门文章</h3>
                        <ul>
                            {{range $index,$value:=.HotList}}
                   <li>
                       <a href="{{$config.SiteURL}}/article/{{$value.Name}}" title="{{$value.Title}}">{{$value.Title}}</a>
                   </li>
                            {{end}}
                        </ul>
                    </div>
                    <div id="tag_cloud-2" class="box widget_tag_cloud">
                        <h3 class="widget-head">标签</h3>
                        <div class="tagcloud">
                            {{range $index,$value:=.TagList}}
                    <a href="{{$config.SiteURL}}/tag/{{$value.Name}}" class="tag-link-{{$index}}"  title="{{$value.Count}}个话题">{{$value.Title}}</a>
                            {{end}}
                        </div>
                    </div>
                    <div id="recent-posts-4" class="box widget_recent_entries">
                        <h3 class="widget-head">最新文章</h3>
                        <ul>
                            {{ range $index,$value:=.RecentList}}
                    <li>
                        <a href="{{$config.SiteURL}}/article/{{$value.Name}}" title="{{$value.Title}}">{{$value.Title}}</a>
                    </li>
                            {{end}}
                        </ul>
                    </div>
                   <!--  <div id="recent-comments-2" class="box widget_recent_comments">
                        <h3 class="widget-head">最新评论</h3>
                        <ul id="recentcomments">
                            <li class="recentcomments"><a href="{{$config.SiteURL}}" rel="external nofollow" class="url">天屹</a>发表在《<a href="http://www.tystudio.net/2013/04/02/asp-net-mvc-structuremap/#comment-2325">使用StructureMap扩展ASP.NET MVC三层架构1-Entity Framework Model层的实现</a>》</li>
                            <li class="recentcomments"><a href="http://www.cnblogs.com/duguqing40" rel="external nofollow" class="url">独孤青</a>发表在《<a href="http://www.tystudio.net/2013/04/02/asp-net-mvc-structuremap/#comment-2324">使用StructureMap扩展ASP.NET MVC三层架构1-Entity Framework Model层的实现</a>》</li>
                            <li class="recentcomments"><a href="{{$config.SiteURL}}" rel="external nofollow" class="url">天屹</a>发表在《<a href="http://www.tystudio.net/2013/03/20/mvc4-simplemembership-permission-system/#comment-2323">MVC4 Simplemembership后台权限管理系统（附源码下载）</a>》</li>
                            <li class="recentcomments"><a href="http://450024895@qq.com/" rel="external nofollow" class="url">伍仟蚊</a>发表在《<a href="http://www.tystudio.net/2013/03/20/mvc4-simplemembership-permission-system/#comment-2322">MVC4 Simplemembership后台权限管理系统（附源码下载）</a>》</li>
                            <li class="recentcomments"><a href="{{$config.SiteURL}}" rel="external nofollow" class="url">天屹</a>发表在《<a href="http://www.tystudio.net/2013/04/02/asp-net-mvc-structuremap/#comment-2321">使用StructureMap扩展ASP.NET MVC三层架构1-Entity Framework Model层的实现</a>》</li>
                        </ul>
                    </div> -->
                </aside>
{{end}}
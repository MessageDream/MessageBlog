(function() {  

    var fullurl;
    tinymce.create('tinymce.plugins.dropcap', {  
        init : function(ed, url) {  
            fullurl = url;
            ed.addButton('dropcap', {  
                title : '效果和图标一样，字母全大写分带背景和不带背景两种',
                image : url+'/dropcap.png',  
                onclick : function() {  
                     ed.selection.setContent('[dropcap background="yes|no"]Letter[/dropcap]');  
                }  
            });  
        },  
        createControl : function(n, cm) {  
            return null;  
        },  
    });  
    tinymce.PluginManager.add('dropcap', tinymce.plugins.dropcap); 

    tinymce.create('tinymce.plugins.pullquote', {  
        createControl: function(n, cm) {
            switch (n) {
                case 'pullquote':
                    var c = cm.createSplitButton('pullquotesplitbutton', {
                        title : '抽取引用',
                        image : fullurl+'/pullquote.png',
                        onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[pullquote align="left" background="on"]' + tinyMCE.activeEditor.selection.getContent() + '[/pullquote]');  
                        }
                    });

                    c.onRenderMenu.add(function(c, m) {
                        m.add({title : '样式', 'class' : 'mceMenuItemTitle'}).setDisabled(1);

                        m.add({title : '左边显示', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[pullquote align="left"]' + tinyMCE.activeEditor.selection.getContent() + '[/pullquote]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '左边显示带背景', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[pullquote align="left" background="on"]' + tinyMCE.activeEditor.selection.getContent() + '[/pullquote]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '右边显示', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[pullquote align="right"]' + tinyMCE.activeEditor.selection.getContent() + '[/pullquote]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '右边显示带背景', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[pullquote align="right" background="on"]' + tinyMCE.activeEditor.selection.getContent() + '[/pullquote]');  
                            tinyMCE.activeEditor.focus();
                        }});
                    });

                  // Return the new splitbutton instance
                  return c;
            }
            return null;
        }
    });  
    tinymce.PluginManager.add('pullquote', tinymce.plugins.pullquote); 

    tinymce.create('tinymce.plugins.alert', {  
        createControl: function(n, cm) {
            switch (n) {
                case 'alert':
                    var c = cm.createSplitButton('alertsplitbutton', {
                        title : '警告框',
                        image : fullurl+'/alert.png',
                        onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[alert style="blue"]' + tinyMCE.activeEditor.selection.getContent() + '[/alert]');  
                        }
                    });

                    c.onRenderMenu.add(function(c, m) {
                        m.add({title : '警告框样式', 'class' : 'mceMenuItemTitle'}).setDisabled(1);

                        m.add({title : '蓝', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[alert style="blue"]' + tinyMCE.activeEditor.selection.getContent() + '[/alert]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '红', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[alert style="red"]' + tinyMCE.activeEditor.selection.getContent() + '[/alert]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '绿', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[alert style="green"]' + tinyMCE.activeEditor.selection.getContent() + '[/alert]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '黄', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[alert style="yellow"]' + tinyMCE.activeEditor.selection.getContent() + '[/alert]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '紫', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[alert style="purple"]' + tinyMCE.activeEditor.selection.getContent() + '[/alert]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '深红', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[alert style="darkred"]' + tinyMCE.activeEditor.selection.getContent() + '[/alert]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '棕色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[alert style="brown"]' + tinyMCE.activeEditor.selection.getContent() + '[/alert]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '灰', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[alert style="grey"]' + tinyMCE.activeEditor.selection.getContent() + '[/alert]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '暗色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[alert style="dark"]' + tinyMCE.activeEditor.selection.getContent() + '[/alert]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '草色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[alert style="grass"]' + tinyMCE.activeEditor.selection.getContent() + '[/alert]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '粉红', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[alert style="pink"]' + tinyMCE.activeEditor.selection.getContent() + '[/alert]');  
                            tinyMCE.activeEditor.focus();
                        }});
                    });

                  // Return the new splitbutton instance
                  return c;
            }
            return null;
        }
    });   
    tinymce.PluginManager.add('alert', tinymce.plugins.alert); 


    tinymce.create('tinymce.plugins.label', {  
        createControl: function(n, cm) {
            switch (n) {
                case 'label':
                    var c = cm.createSplitButton('labelsplitbutton', {
                        title : 'Label标签',
                        image : fullurl+'/label.png',
                        onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[label style="blue"]' + tinyMCE.activeEditor.selection.getContent() + '[/label]');  
                        }
                    });

                    c.onRenderMenu.add(function(c, m) {
                        m.add({title : '标签样式', 'class' : 'mceMenuItemTitle'}).setDisabled(1);

                        m.add({title : '蓝', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[label style="blue"]' + tinyMCE.activeEditor.selection.getContent() + '[/label]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '红', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[label style="red"]' + tinyMCE.activeEditor.selection.getContent() + '[/label]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '绿', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[label style="green"]' + tinyMCE.activeEditor.selection.getContent() + '[/label]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '黄', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[label style="yellow"]' + tinyMCE.activeEditor.selection.getContent() + '[/label]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '紫', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[label style="purple"]' + tinyMCE.activeEditor.selection.getContent() + '[/label]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '深红', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[label style="darkred"]' + tinyMCE.activeEditor.selection.getContent() + '[/label]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '棕色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[label style="brown"]' + tinyMCE.activeEditor.selection.getContent() + '[/label]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '灰', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[label style="grey"]' + tinyMCE.activeEditor.selection.getContent() + '[/label]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '深色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[label style="dark"]' + tinyMCE.activeEditor.selection.getContent() + '[/label]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '草色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[label style="grass"]' + tinyMCE.activeEditor.selection.getContent() + '[/label]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '粉红', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[label style="pink"]' + tinyMCE.activeEditor.selection.getContent() + '[/label]');  
                            tinyMCE.activeEditor.focus();
                        }});
                    });

                  // Return the new splitbutton instance
                  return c;
            }
            return null;
        }
    });  
    tinymce.PluginManager.add('label', tinymce.plugins.label); 


    tinymce.create('tinymce.plugins.badge', {  
        createControl: function(n, cm) {
            switch (n) {
                case 'badge':
                    var c = cm.createSplitButton('badgesplitbutton', {
                        title : '徽章',
                        image : fullurl+'/badge.png',
                        onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[badge style="blue"]' + tinyMCE.activeEditor.selection.getContent() + '[/badge]');  
                        }
                    });

                    c.onRenderMenu.add(function(c, m) {
                        m.add({title : '徽章样式', 'class' : 'mceMenuItemTitle'}).setDisabled(1);

                        m.add({title : '蓝', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[badge style="blue"]' + tinyMCE.activeEditor.selection.getContent() + '[/badge]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '红', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[badge style="red"]' + tinyMCE.activeEditor.selection.getContent() + '[/badge]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '绿', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[badge style="green"]' + tinyMCE.activeEditor.selection.getContent() + '[/badge]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '黄', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[badge style="yellow"]' + tinyMCE.activeEditor.selection.getContent() + '[/badge]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '紫', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[badge style="purple"]' + tinyMCE.activeEditor.selection.getContent() + '[/badge]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '深红', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[badge style="darkred"]' + tinyMCE.activeEditor.selection.getContent() + '[/badge]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '棕色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[badge style="brown"]' + tinyMCE.activeEditor.selection.getContent() + '[/badge]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '灰', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[badge style="grey"]' + tinyMCE.activeEditor.selection.getContent() + '[/badge]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '暗色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[badge style="dark"]' + tinyMCE.activeEditor.selection.getContent() + '[/badge]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '草色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[badge style="grass"]' + tinyMCE.activeEditor.selection.getContent() + '[/badge]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '粉红', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[badge style="pink"]' + tinyMCE.activeEditor.selection.getContent() + '[/badge]');  
                            tinyMCE.activeEditor.focus();
                        }});
                    });

                  // Return the new splitbutton instance
                  return c;
            }
            return null;
        }
    });  
    tinymce.PluginManager.add('badge', tinymce.plugins.badge); 


    tinymce.create('tinymce.plugins.well', {  
        init : function(ed, url) {  
            ed.addButton('well', {  
                title : 'Add a Well（不知道怎么翻译，使用后文字包含在灰色背景的盒子里）',
                image : url+'/well.png',  
                onclick : function() {  
                     ed.selection.setContent('[well]Well content goes here[/well]');  
                }  
            });  
        },  
        createControl : function(n, cm) {  
            return null;  
        },  
    });  
    tinymce.PluginManager.add('well', tinymce.plugins.well); 


    tinymce.create('tinymce.plugins.button', {  
        createControl: function(n, cm) {
            switch (n) {
                case 'button':
                    var c = cm.createSplitButton('buttonsplitbutton', {
                        title : '按钮',
                        image : fullurl+'/button.png',
                        onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="#" style="blue" size="default" block="false" target="_blank" icon="check"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                        }
                    });

                    c.onRenderMenu.add(function(c, m) {
                        m.add({title : '按钮样式', 'class' : 'mceMenuItemTitle'}).setDisabled(1);

                        m.add({title : '蓝', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" style="blue"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '红', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" style="red"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '绿', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" style="green"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '黄', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" style="yellow"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '紫', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" style="purple"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '深红', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" style="darkred"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '棕色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" style="brown"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '灰', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" style="grey"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '暗色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" style="dark"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '草色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" style="grass"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '粉红', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" style="pink"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '大按钮', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" size="large"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '小按钮', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" size="small"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '迷你按钮', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" size="mini"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '块状按钮', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" block="true"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '带图标的按钮', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" icon="check-1"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                        m.add({title : '点击打开新窗口', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[button url="http://" target="_blank"]' + tinyMCE.activeEditor.selection.getContent() + '[/button]');  
                            tinyMCE.activeEditor.focus();
                        }});
                    });

                  // Return the new splitbutton instance
                  return c;
            }
            return null;
        }
    });  
    tinymce.PluginManager.add('button', tinymce.plugins.button);


    tinymce.create('tinymce.plugins.blockquote', {  
        init : function(ed, url) {  
            ed.addButton('blockquote', {  
                title : '添加块状引用',
                image : url+'/blockquote.png',  
                onclick : function() {  
                     ed.selection.setContent('[blockquote source="Name of the source"]引用文字写在这[/blockquote]');
                }  
            });  
        },  
        createControl : function(n, cm) {  
            return null;  
        },  
    });  
    tinymce.PluginManager.add('blockquote', tinymce.plugins.blockquote);


    tinymce.create('tinymce.plugins.syntax', {  
        init : function(ed, url) {  
            ed.addButton('syntax', {  
                title : '添加代码高亮',
                image : url+'/syntax.png',  
                onclick : function() {  
                     ed.selection.setContent('[syntax type="html|php|js|css"]代码写在这[/syntax]');
                }  
            });  
        },  
        createControl : function(n, cm) {  
            return null;  
        },  
    });  
    tinymce.PluginManager.add('syntax', tinymce.plugins.syntax);


    tinymce.create('tinymce.plugins.tooltip', {  
        init : function(ed, url) {  
            ed.addButton('tooltip', {  
                title : '添加提示工具',
                image : url+'/tooltip.png',  
                onclick : function() {  
                     ed.selection.setContent('[tooltip text="提示的文字写在这里" trigger="hover|click"]文字写在这[/tooltip]');
                }  
            });  
        },  
        createControl : function(n, cm) {  
            return null;  
        },  
    });  
    tinymce.PluginManager.add('tooltip', tinymce.plugins.tooltip); 

    tinymce.create('tinymce.plugins.progress', {  
        init : function(ed, url) {  
            ed.addButton('progress', {  
                title : '添加进度条',
                image : url+'/progress.png',  
                onclick : function() {  
                     ed.selection.setContent('[progress length="50" color="#3bd2f8"]文字[/progress]');
                }  
            });  
        },  
        createControl : function(n, cm) {  
            return null;  
        },  
    });  
    tinymce.PluginManager.add('progress', tinymce.plugins.progress);  


    tinymce.create('tinymce.plugins.popover', {  
        init : function(ed, url) {  
            ed.addButton('popover', {  
                title : '添加Popover（带标题和内容的提示框）',
                image : url+'/popover.png',  
                onclick : function() {  
                     ed.selection.setContent('[popover title="提示标题写在这里" trigger="hover|click" placement="top|bottom|left|right" text="提示内容写在这里"]文字写在这里[/popover]');
                }  
            });  
        },  
        createControl : function(n, cm) {  
            return null;  
        },  
    });  
    tinymce.PluginManager.add('popover', tinymce.plugins.popover);  

    tinymce.create('tinymce.plugins.accordion', {  
        init : function(ed, url) {  
            fullurl = url;
            ed.addButton('accordion', {  
                title : '添加折叠手风琴文本',
                image : url+'/accordion.png',  
                onclick : function() {  
                     ed.selection.setContent('[accordion]<br />[accordion-group title="标题写在这里"]文字内容写在这里[/accordion-group]<br />[accordion-group title="标题写在这里"]文字内容写在这里[/accordion-group]<br />[accordion-group title="标题写在这里"]文字内容写在这里[/accordion-group]<br />[/accordion]');
                  }  
            });  
        },   
        createControl: function(n, cm) {
            switch (n) {
                case 'accordion':
                    var c = cm.createSplitButton('accordionsplitbutton', {
                        title : '添加折叠手风琴文本',
                        image : fullurl+'/accordion.png',
                        onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[accordion]<br />[accordion-group title="标题写在这里"]文字内容写在这里[/accordion-group]<br />[accordion-group title="标题写在这里"]文字内容写在这里[/accordion-group]<br />[accordion-group title="标题写在这里"]文字内容写在这里[/accordion-group]<br />[/accordion]');
                        }
                    });

                    c.onRenderMenu.add(function(c, m) {
                        m.add({title : '折叠手风琴样式', 'class' : 'mceMenuItemTitle'}).setDisabled(1);

                        m.add({title : '暗色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[accordion]<br />[accordion-group title="Title goes here"]Text goes here[/accordion-group]<br />[accordion-group title="Title goes here"]Text goes here[/accordion-group]<br />[accordion-group title="Title goes here"]Text goes here[/accordion-group]<br />[/accordion]');  
                        }});
                        m.add({title : '亮色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[accordion]<br />[accordion-group style="light-grey" title="Title goes here"]Text goes here[/accordion-group]<br />[accordion-group style="light-grey" title="Title goes here"]Text goes here[/accordion-group]<br />[accordion-group style="light-grey" title="Title goes here"]Text goes here[/accordion-group]<br />[/accordion]');  
                        }});
                    });

                  // Return the new splitbutton instance
                  return c;
            }

            return null;
        }
    });  
    tinymce.PluginManager.add('accordion', tinymce.plugins.accordion);  
    
    tinymce.create('tinymce.plugins.divider', {  
        init : function(ed, url) {  
            ed.addButton('divider', {  
                title : '添加分割线',
                image : url+'/divider.png',  
                onclick : function() {  
                     ed.selection.setContent('[divider type="white|thin|thick|short|dotted|dashed" spacing="10"]');  
                }  
            });  
        },  
        createControl: function(n, cm) {
            switch (n) {
                case 'divider':
                    var c = cm.createSplitButton('dividersplitbutton', {
                        title : '分割线',
                        image : fullurl+'/divider.png',
                        onclick : function() {}
                    });

                    c.onRenderMenu.add(function(c, m) {
                        m.add({title : '分割线样式', 'class' : 'mceMenuItemTitle'}).setDisabled(1);

                        m.add({title : '白色', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[divider type="white"]');  
                        }});
                        m.add({title : '细线', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[divider type="thin"]');  
                        }});
                        m.add({title : '粗线', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[divider type="thick"]');  
                        }});
                        m.add({title : '短线', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[divider type="short"]');  
                        }});
                        m.add({title : '密集点线', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[divider type="dotted"]');  
                        }});
                        m.add({title : '稀疏点线', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[divider type="dashed"]');  
                        }});
                        m.add({title : '带空间的密集点线', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[divider type="dashed" spacing="25"]');  
                        }});
                    });

                  // Return the new splitbutton instance
                  return c;
            }

            return null;
        }
    });  
    tinymce.PluginManager.add('divider', tinymce.plugins.divider); 

    tinymce.create('tinymce.plugins.columns', {
        createControl: function(n, cm) {
            switch (n) {
                case 'columns':
                    var c = cm.createSplitButton('mysplitbutton', {
                        title : '多列',
                        image : fullurl+'/columns.png',
                        onclick : function() {}
                    });

                    c.onRenderMenu.add(function(c, m) {
                        m.add({title : '多列', 'class' : 'mceMenuItemTitle'}).setDisabled(1);

                        m.add({title : '均分两列', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[two_first]<br /><br />[/two_first][two_second]<br /><br />[/two_second]');  
                        }});

                        m.add({title : '2/1（第一个是第二的两倍）', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[two_one_first]<br /><br />[/two_one_first][two_one_second]<br /><br />[/two_one_second]');  
                        }});

                        m.add({title : '1/2（第一个是第二个的一半）', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[one_two_first]<br /><br />[/one_two_first][one_two_second]<br /><br />[/one_two_second]');  
                        }});

                        m.add({title : '均分三列', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[three_first]<br /><br />[/three_first][three_second]<br /><br />[/three_second][three_third]<br /><br />[/three_third]');  
                        }});

                        m.add({title : '均分四列', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[four_first]<br /><br />[/four_first][four_second]<br /><br />[/four_second][four_third]<br /><br />[/four_third][four_fourth]<br /><br />[/four_fourth]');  
                        }});

                        m.add({title : '1/1/2', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[one_one_two_first]<br /><br />[/one_one_two_first][one_one_two_second]<br /><br />[/one_one_two_second][one_one_two_third]<br /><br />[/one_one_two_third]');  
                        }});

                        m.add({title : '2/1/1', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[two_one_one_first]<br /><br />[/two_one_one_first][two_one_one_second]<br /><br />[/two_one_one_second][two_one_one_third]<br /><br />[/two_one_one_third]');  
                        }});

                        m.add({title : '1/2/1', onclick : function() {
                            tinyMCE.activeEditor.selection.setContent('[one_two_one_first]<br /><br />[/one_two_one_first][one_two_one_second]<br /><br />[/one_two_one_second][one_two_one_third]<br /><br />[/one_two_one_third]');  
                        }});
                    });

                  // Return the new splitbutton instance
                  return c;
            }

            return null;
        }
    });
    tinymce.PluginManager.add('columns', tinymce.plugins.columns);

})();
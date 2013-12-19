var editor= CKEDITOR.replace('content1', {
                filebrowserBrowseUrl:'/root/node',
                filebrowserImageUploadUrl: '/root/upload'
             });
    // $(function(){
    //     $("#isThumbnail").change(function(){
    //      if($(this).attr("checked")){
    //         $("#featuredPic").css("display","none")
    //     }else{
    //         $("#featuredPic").css("display","block")
    //     }
    // });
    //})

var Articles;
(function (Articles) {
    var Article = (function () {
        function Article(btn) {
            this.btn = btn;
        }
        Article.prototype.gettoupdate = function () {
            var _this = this;
            for (var i = 0; i < this.btn.length; i++) {
                this.btn[i].onclick = function (e) {
                    var _thisbtn = e.currentTarget;
                    _this.id_ = _thisbtn.parentElement.parentElement.children[0].textContent;
                    $.get("/root/article", { id: _this.id_ }, function (data) {
                       // $("#selectnode").val(data.NName)
                        var opt= $("#selectnode option");
                       for (var i = 0; i < opt.length; i++) {
                          if($(opt[i]).val()==data.NName){
                            var txt=$(opt[i]).text();
                            var li=$(".active-result")
                            for (var i = 0; i < li.length; i++) {
                               $(li[i]).removeClass("result-selected");
                               if($(li[i]).text()==txt){
                                $(li[i]).addClass("result-selected");
                                $("a.chzn-single > span").text(txt)
                                $(opt[i]).attr("selected","selected");
                                break;
                               }
                            };
                            break;
                          } 
                       };
                      // $("selectnode").trigger("liszt:updated");


                        $("#id").val(data.Id_)
                        $("#title").val(data.Title);
                        $("#name").val(data.Name);
                         editor.setData(data.Text);
                        if (data.IsThumbnail) {
                            $("#uniform-isThumbnail > span").addClass("checked")
                           // $("#isThumbnail").attr("checked", "checked");
                        }else{
                            $("#uniform-isThumbnail > span").removeClass("checked")
                          //   $("#isThumbnail").attr("checked", "");
                        }
                        $("#featuredPicURL").val(data.FeaturedPicURL);

                        $("#tags option").each(function(){
                           if($.inArray($(this).val(),data.Tags)!=-1){
                            $(this).attr("selected","selected")
                          }
                        });
                        updateselect($("#tags"))

                        if($("#hidecontent").is(":hidden")) {
                            $("#expand").click();
                        }
                       
                    });
                };
            }
        };
        Article.prototype.deletebyid = function () {
            var _this = this;
            for (var i = 0; i < this.btn.length; i++) {
                this.btn[i].onclick = function (e) {
                    if(confirm("真的要删除此文章吗？")){
                    var _thisbtn = e.currentTarget;
                    _this.id_ = _thisbtn.parentElement.parentElement.children[0].textContent;
                    $.post("/root/article", { id: _this.id_ },function(data){
                        if(data){

                             self.location="/root/article";
                        }
                    });
                    }
                };
            }
        };
        return Article;
    })();

    function updateselect(obj){
        var selects = obj;
        var selected = [];
        selects.find("option").each(function() {
            if (this.selected) {
                selected[this.value] = this;
            }
        }).each(function() {
            this.disabled = selected[this.value] && selected[this.value] !== this;
        });
        selects.trigger("liszt:updated");
    }
    window.onload = function () {
        new Article(document.getElementsByClassName("btn btn-info")).gettoupdate();
        new Article(document.getElementsByClassName("btn btn-danger")).deletebyid();
        $("#cancelAdd").click(function(e){
           e.preventDefault();
          $("#expand").click();
        });
        $('#addTag').click(function(e){
                    e.preventDefault();
                    $('#myModal').modal('show');
        });
        $('#addtagok').click(function(e){
           e.preventDefault();
            var postdata=$('#newTag').val();
             $.post("/root/tag", { title: postdata },function(data){
                        if(data==1){
                             $("#tags").append("<option selected value='"+postdata+"'>"+postdata+"</option>");
                              updateselect($("#tags"));
                              $("#closetag").click();
                        }
                    });
        });
    };
})(Articles || (Articles = {}));

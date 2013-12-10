var Nodes;
(function (Nodes) {
    var Node = (function () {
        function Node(btn) {
            this.btn = btn;
        }
        Node.prototype.gettoupdate = function () {
            var _this = this;
            for (var i = 0; i < this.btn.length; i++) {
                this.btn[i].onclick = function (e) {
                    var _thisbtn = e.currentTarget;
                    _this.name = _thisbtn.parentElement.parentElement.children[0].textContent;
                    $.get("/root/node", { name: _this.name }, function (data) {
                         var opt= $("#selectcategory option")
                       for (var i = 0; i < opt.length; i++) {
                          if($(opt[i]).val()==data.CategoryId){
                            var txt=data.CategoryTitle;
                            var li=$(".active-result")
                            for (var i = 0; i < li.length; i++) {
                               if($(li[i]).text()==txt){
                                $(li[i]).addClass("result-selected");
                                $("a.chzn-single > span").text(txt)
                                break;
                               }
                            };
                            break;
                          } 
                       };
                       // $("#selectcategory").val(data.CategoryId)
                        $("#title").val(data.Title);
                        $("#name").val(data.Name);
                        $("#nodecontent").val(data.Content);
                        if($("#hidecontent").is(":hidden")) {
                            $("#expand").click()
                        }
                       
                    });
                };
            }
        };
        Node.prototype.deletebyid = function () {
            var _this = this;
            for (var i = 0; i < this.btn.length; i++) {
                this.btn[i].onclick = function (e) {
                    if(confirm("真的要删除此节点吗？")){
                    var _thisbtn = e.currentTarget;
                    _this.id_ = _thisbtn.children[1].textContent;
                    _this.nname = _thisbtn.parentElement.parentElement.children[0].textContent;
                    $.post("/root/node", { id: _this.id_ ,nname:_this.nname},function(data){
                        if(data){
                             self.location="/root/node";
                        }
                    });
                    }
                };
            }
        };
        return Node;
    })();
    window.onload = function () {
        new Node(document.getElementsByClassName("btn btn-info")).gettoupdate();
        new Node(document.getElementsByClassName("btn btn-danger")).deletebyid();
        showmsg();
    };
})(Nodes || (Nodes = {}));

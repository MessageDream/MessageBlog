
var Categories;
(function (Categories) {
    var Category = (function () {
        function Category(btn) {
            this.btn = btn;
        }
        Category.prototype.gettoupdate = function () {
            var _this = this;
            for (var i = 0; i < this.btn.length; i++) {
                this.btn[i].onclick = function (e) {
                    var _thisbtn = e.currentTarget;
                    _this.id_ = _thisbtn.parentElement.parentElement.children[0].textContent;
                    $.get("/root/category", { id: _this.id_ }, function (data) {
                        $("#id").val(data.Id_)
                        $("#title").val(data.Title);
                        $("#name").val(data.Name);
                        $("#contentid").val(data.Content);
                        if($("#hidecontent").is(":hidden")) {
                            $("#expand").click()
                        }
                       
                    });
                };
            }
        };
        Category.prototype.deletebyid = function () {
            var _this = this;
            for (var i = 0; i < this.btn.length; i++) {
                this.btn[i].onclick = function (e) {
                    if(confirm("真的要删除此分类吗？")){
                    var _thisbtn = e.currentTarget;
                    _this.id_ = _thisbtn.parentElement.parentElement.children[0].textContent;
                    $.post("/root/category", { id: _this.id_ },function(data){
                        if(data){
                             self.location="/root/category";
                        }
                    });
                    }
                };
            }
        };
        return Category;
    })();
    window.onload = function () {
        new Category(document.getElementsByClassName("btn btn-info")).gettoupdate();
        new Category(document.getElementsByClassName("btn btn-danger")).deletebyid();
        showmsg();
    };
})(Categories || (Categories = {}));
